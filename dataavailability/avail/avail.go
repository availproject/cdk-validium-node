// SPDX-License-Identifier: Apache-2.0
package avail

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	availConfig "github.com/0xPolygonHermez/zkevm-node/dataavailability/avail/config"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/polygondatacommittee"
	"github.com/0xPolygonHermez/zkevm-node/log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	availTypes "github.com/0xPolygonHermez/zkevm-node/dataavailability/avail/types"
)

type AccountNextIndexRPCResponse struct {
	Result uint `json:"result"`
}

type DataProofRPCResponse struct {
	Result DataProof `json:"result"`
}

type DataProof struct {
	Root           string   `json:"root"`
	Proof          []string `json:"proof"`
	NumberOfLeaves uint     `json:"numberOfLeaves"`
	LeafIndex      uint     `json:"leafIndex"`
	Leaf           string   `json:"leaf"`
}

type AvailBackend struct {
	config              availConfig.Config
	attestationContract *polygondatacommittee.Polygondatacommittee
	api                 *gsrpc.SubstrateAPI
	httpApi             string
	meta                *types.Metadata
	appId               int
	genesisHash         types.Hash
	rv                  *types.RuntimeVersion
	keyringPair         signature.KeyringPair
}

func New(l1RPCURL string, dataCommitteeAddr common.Address) (*AvailBackend, error) {
	var config availConfig.Config
	err := config.GetConfig("/app/avail-config.json")
	if err != nil {
		log.Fatalf("cannot get config:%w", err)
	}

	ethClient, err := ethclient.Dial(l1RPCURL)
	if err != nil {
		log.Errorf("error connecting to %s: %+v", l1RPCURL, err)
		return nil, err
	}

	attestationContract, err := polygondatacommittee.NewPolygondatacommittee(dataCommitteeAddr, ethClient)
	if err != nil {
		return nil, err
	}

	api, err := gsrpc.NewSubstrateAPI(config.WsApiUrl)
	if err != nil {
		log.Fatalf("cannot get ws api:%w", err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("cannot get metadata:%w", err)
	}

	appId := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appId = config.AppID
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		log.Fatalf("cannot get block hash:%w", err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		log.Fatalf("cannot get runtime version:%w", err)
	}

	keyringPair, err := signature.KeyringPairFromSecret(config.Seed, 42)
	if err != nil {
		log.Fatalf("cannot create keypair:%w", err)
	}

	return &AvailBackend{
		config:              config,
		attestationContract: attestationContract,
		api:                 api,
		httpApi:             config.HttpApiUrl,
		meta:                meta,
		appId:               appId,
		genesisHash:         genesisHash,
		rv:                  rv,
		keyringPair:         keyringPair,
	}, nil
}

func (a *AvailBackend) Init() error {
	return nil
}

func (a *AvailBackend) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	byteArrayType, _ := abi.NewType("bytes[]", "", nil)
	args := abi.Arguments{
		{Type: byteArrayType, Name: "data"},
	}
	sequence, err := args.Pack(&batchesData)
	if err != nil {
		return nil, fmt.Errorf("cannot pack data:%w", err)
	}

	log.Infof("⚡️ Prepared data for Avail:%d bytes", len(sequence))

	newCall, err := types.NewCall(a.meta, "DataAvailability.submit_data", types.NewBytes(sequence))
	if err != nil {
		return nil, fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := a.GetAccountNextIndex()
	if err != nil {
		return nil, fmt.Errorf("cannot get account next index:%w", err)
	}

	options := types.SignatureOptions{
		BlockHash:          a.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        a.genesisHash,
		Nonce:              nonce,
		SpecVersion:        a.rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(1000),
		AppID:              types.NewUCompactFromUInt(uint64(a.appId)),
		TransactionVersion: a.rv.TransactionVersion,
	}

	fmt.Printf("options: %+v\n", options)

	err = ext.Sign(a.keyringPair, options)
	if err != nil {
		return nil, fmt.Errorf("cannot sign:%w", err)
	}

	// Send the extrinsic
	sub, err := a.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, fmt.Errorf("cannot submit extrinsic:%w", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(a.config.Timeout) * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Infof("📥 Submit data extrinsic included in block %v", status.AsInBlock.Hex())
			}
			if status.IsFinalized {
				blockHash = status.AsFinalized
				break out
			} else if status.IsDropped {
				return nil, fmt.Errorf("❌ Extrinsic dropped")
			} else if status.IsUsurped {
				return nil, fmt.Errorf("❌ Extrinsic usurped")
			} else if status.IsRetracted {
				return nil, fmt.Errorf("❌ Extrinsic retracted")
			} else if status.IsInvalid {
				return nil, fmt.Errorf("❌ Extrinsic invalid")
			}
		case <-timeout:
			return nil, fmt.Errorf("⌛️ Timeout of %d seconds reached without getting finalized status for extrinsic", a.config.Timeout)
		}
	}

	var dataProof DataProof
	batchHash := crypto.Keccak256Hash(sequence)

	block, err := a.api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block:%w", err)
	}

	for i := 1; i <= len(block.Block.Extrinsics); i++ {
		resp, err := http.Post(a.httpApi, "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", i, blockHash)))
		if err != nil {
			return nil, fmt.Errorf("cannot post query request:%v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("cannot read body:%v", err)
		}

		var dataProofResp DataProofRPCResponse
		json.Unmarshal(data, &dataProofResp)

		if dataProofResp.Result.Leaf == fmt.Sprintf("%#x", batchHash) {
			dataProof = dataProofResp.Result
			break
		}
	}

	log.Infof("💿 received data proof:%+v", dataProof)
	var batchDAData availTypes.BatchDAData
	batchDAData.LeafIndex = dataProof.LeafIndex

	header, err := a.api.RPC.Chain.GetHeader(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get header:%+v", err)
	}

	batchDAData.BlockNumber = uint(header.Number)
	a.GetData(uint64(header.Number), dataProof.LeafIndex)
	log.Infof("🟢 prepared DA data:%+v", batchDAData)

	// todo: use bridge API data
	returnData, err := batchDAData.Encode()
	if err != nil {
		return nil, fmt.Errorf("cannot encode batch data:%w", err)
	}

	log.Infof("✅ Data submitted by sequencer:%d bytes against AppID %v sent with hash %#x", len(sequence), a.appId, blockHash)

	return returnData, nil
}

func (a *AvailBackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	// TODO: implement
	return nil, nil
}

func (a *AvailBackend) GetData(blockNumber uint64, index uint) ([]byte, error) {
	blockHash, err := a.api.RPC.Chain.GetBlockHash(uint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("cannot get block hash:%w", err)
	}

	block, err := a.api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block:%w", err)
	}

	var data [][]byte
	for _, ext := range block.Block.Extrinsics {
		if ext.Method.CallIndex.SectionIndex == 29 && ext.Method.CallIndex.MethodIndex == 1 {
			data = append(data, ext.Method.Args[2:])
		}
	}

	return data[index], nil
}

func (a *AvailBackend) GetAccountNextIndex() (types.UCompact, error) {
	resp, err := http.Post(a.httpApi, "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"system_accountNextIndex\",\"params\":[\"%v\"]}", a.keyringPair.Address)))
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot post query request:%v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot read body:%v", err)
	}

	var accountNextIndex AccountNextIndexRPCResponse
	json.Unmarshal(data, &accountNextIndex)

	return types.NewUCompactFromUInt(uint64(accountNextIndex.Result)), nil
}
