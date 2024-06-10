// SPDX-License-Identifier: Apache-2.0
package avail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	availConfig "github.com/0xPolygonHermez/zkevm-node/dataavailability/avail/config"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/availattestation"
	"github.com/0xPolygonHermez/zkevm-node/log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/umbracle/ethgo/abi"
)

type MerkleProofInput struct {
	DataRootProof [][32]byte `abi:"dataRootProof"`
	LeafProof     [][32]byte `abi:"leafProof"`
	RangeHash     [32]byte   `abi:"rangeHash"`
	DataRootIndex *big.Int   `abi:"dataRootIndex"`
	BlobRoot      [32]byte   `abi:"blobRoot"`
	BridgeRoot    [32]byte   `abi:"bridgeRoot"`
	Leaf          [32]byte   `abi:"leaf"`
	LeafIndex     *big.Int   `abi:"leafIndex"`
}

type BridgeAPIResponse struct {
	BlobRoot           common.Hash   `json:"blobRoot"`
	BlockHash          common.Hash   `json:"blockHash"`
	BridgeRoot         common.Hash   `json:"bridgeRoot"`
	DataRoot           common.Hash   `json:"dataRoot"`
	DataRootIndex      *big.Int      `json:"dataRootIndex"`
	DataRootCommitment common.Hash   `json:"dataRootCommitment"`
	DataRootProof      []common.Hash `json:"dataRootProof"`
	Leaf               common.Hash   `json:"leaf"`
	LeafIndex          *big.Int      `json:"leafIndex"`
	LeafProof          []common.Hash `json:"leafProof"`
	RangeHash          common.Hash   `json:"rangeHash"`
}

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
	attestationContract *availattestation.Availattestation
	api                 *gsrpc.SubstrateAPI
	httpApi             string
	bridgeApi           string
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
		log.Fatalf("cannot get config: %+v", err)
		return nil, err
	}

	ethClient, err := ethclient.Dial(l1RPCURL)
	if err != nil {
		log.Errorf("error connecting to %s: %+v", l1RPCURL, err)
		return nil, err
	}

	attestationContract, err := availattestation.NewAvailattestation(dataCommitteeAddr, ethClient)
	log.Infof("📜 Attestation contract address: %v", dataCommitteeAddr)
	if err != nil {
		return nil, err
	}

	api, err := gsrpc.NewSubstrateAPI(config.WsApiUrl)
	if err != nil {
		log.Fatalf("cannot get ws api: %+v", err)
		return nil, err
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("cannot get metadata: %+v", err)
		return nil, err
	}

	appId := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appId = config.AppID
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		log.Fatalf("cannot get block hash: %+v", err)
		return nil, err
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		log.Fatalf("cannot get runtime version: %+v", err)
		return nil, err
	}

	keyringPair, err := signature.KeyringPairFromSecret(config.Seed, 42)
	if err != nil {
		log.Fatalf("cannot create keypair: %+v", err)
		return nil, err
	}
	log.Infof("🔑 Using KeyringPair with address %v", keyringPair.Address)

	return &AvailBackend{
		config:              config,
		attestationContract: attestationContract,
		api:                 api,
		httpApi:             config.HttpApiUrl,
		bridgeApi:           config.BridgeApiUrl,
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
	fmt.Printf("keyringpair address: %v\n", a.keyringPair.Address)
	typ := abi.MustNewType("bytes[]")
	sequence, err := typ.Encode(batchesData)
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

	nonce, err := a.getAccountNextIndex()
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

	log.Infof("✅ Data submitted by sequencer:%d bytes against AppID %v sent with hash %#x", len(sequence), a.appId, blockHash)

	txIndex, err := a.getTxIndex(blockHash, sequence)
	if err != nil {
		return nil, fmt.Errorf("cannot get tx index:%+v", err)
	}

	var input BridgeAPIResponse
	waitTime := time.Duration(420) * time.Second
	for {
		log.Infof("Bridge API URL: %v", fmt.Sprintf("%s/eth/proof/%#x?index=%d", a.bridgeApi, blockHash, txIndex))
		resp, err := http.Get(fmt.Sprintf("%s/eth/proof/%#x?index=%d", a.bridgeApi, blockHash, txIndex))
		if err != nil {
			log.Infof("⏳ Attestation proof RPC errored, waiting...")
		} else {
			if resp.StatusCode == 200 {
				log.Infof("✅ Attestation proof received")
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("cannot read body:%v", err)
				}
				err = json.Unmarshal(data, &input)
				if err != nil {
					return nil, fmt.Errorf("cannot unmarshal data:%v", err)
				}
				break
			}
			defer resp.Body.Close()
		}
		time.Sleep(waitTime)
	}
	log.Infof("🔗 Attestation proof received: %+v", input)
	var dataRootProof [][32]byte
	for _, hash := range input.DataRootProof {
		dataRootProof = append(dataRootProof, hash)
	}
	var leafProof [][32]byte
	for _, hash := range input.LeafProof {
		leafProof = append(leafProof, hash)
	}
	merkleProofInput := &MerkleProofInput{
		DataRootProof: dataRootProof,
		LeafProof:     leafProof,
		RangeHash:     input.RangeHash,
		DataRootIndex: input.DataRootIndex,
		BlobRoot:      input.BlobRoot,
		BridgeRoot:    input.BridgeRoot,
		Leaf:          input.Leaf,
		LeafIndex:     input.LeafIndex,
	}
	log.Infof("🔗 Merkle proof input: %+v", merkleProofInput)
	typ = abi.MustNewType("tuple(bytes32[] dataRootProof,bytes32[] leafProof,bytes32 rangeHash,uint256 dataRootIndex,bytes32 blobRoot,bytes32 bridgeRoot,bytes32 leaf,uint256 leafIndex)")
	ret, err := typ.Encode(merkleProofInput)
	if err != nil {
		return nil, fmt.Errorf("cannot encode data:%v", err)
	}
	return ret, nil
}

func (a *AvailBackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	typ := abi.MustNewType("tuple(bytes32[] dataRootProof,bytes32[] leafProof,bytes32 rangeHash,uint256 dataRootIndex,bytes32 blobRoot,bytes32 bridgeRoot,bytes32 leaf,uint256 leafIndex)")
	var inp *MerkleProofInput
	typ.DecodeStruct(dataAvailabilityMessage, &inp)
	attestationData, err := a.attestationContract.Attestations(nil, inp.Leaf)
	if err != nil {
		return nil, fmt.Errorf("cannot get attestation data from contract:%v", err)
	}
	blobData, err := a.getData(uint64(attestationData.BlockNumber), uint(attestationData.LeafIndex.Int64()))
	if err != nil {
		return nil, fmt.Errorf("cannot get data from block:%v", err)
	}
	typ = abi.MustNewType("bytes[]")
	parsed, err := typ.Decode(blobData)
	ret, ok := parsed.([][]byte)
	if !ok {
		return nil, fmt.Errorf("cannot parse data")
	}
	if err != nil {
		return nil, fmt.Errorf("cannot decode data:%v", err)
	}

	return ret, nil
}

func (a *AvailBackend) getTxIndex(blockHash types.Hash, blob []byte) (uint, error) {
	block, err := a.api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return 0, fmt.Errorf("❎ Cannot get block: %w", err)
	}

	var idx uint
	for i, ext := range block.Block.Extrinsics {
		if ext.Method.CallIndex.SectionIndex == 29 && ext.Method.CallIndex.MethodIndex == 1 {
			var availBlob []byte
			err = scale.NewDecoder(bytes.NewReader(ext.Method.Args)).Decode(&availBlob)
			if err != nil {
				return 0, fmt.Errorf("❎ Error while scale decoding blob: %w", err)
			}
			if bytes.Equal(availBlob, blob) {
				idx = uint(i)
				break
			}
		}
	}
	if idx == 0 {
		return 0, fmt.Errorf("❎ Cannot find tx index")
	}

	return idx, nil
}

func (a *AvailBackend) getData(blockNumber uint64, index uint) ([]byte, error) {
	blockHash, err := a.api.RPC.Chain.GetBlockHash(uint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("❎ Cannot get block hash:%w", err)
	}

	block, err := a.api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("❎ Cannot get block:%w", err)
	}

	var idx uint = 0
	for _, ext := range block.Block.Extrinsics {
		if ext.Method.CallIndex.SectionIndex == 29 && ext.Method.CallIndex.MethodIndex == 1 {
			var availBlob []byte
			err = scale.NewDecoder(bytes.NewReader(ext.Method.Args)).Decode(&availBlob)
			if err != nil {
				return nil, fmt.Errorf("❎ Error while scale decoding blob: %w", err)
			}
			if idx == index {
				return availBlob, nil
			}
			idx++
		}
	}
	return nil, fmt.Errorf("❎ Cannot find data")
}

func (a *AvailBackend) getAccountNextIndex() (types.UCompact, error) {
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
