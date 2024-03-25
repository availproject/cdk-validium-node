// SPDX-License-Identifier: Apache-2.0
package types

import (
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type BatchDAData struct {
	BlockNumber uint
	Proof       []string `json:"proof"`
	Width       uint     `json:"number_of_leaves"`
	LeafIndex   uint     `json:"leaf_index"`
}

// write a function that encode batchDAData struct into ABI-encoded bytes
func (b BatchDAData) Encode() ([]byte, error) {
	abi, err := abi.JSON(strings.NewReader(`[{"type":"uint","name":"blockNumber"},{"type":"string[]","name":"proof"},{"type":"uint","name":"width"},{"type":"uint","name":"leafIndex"}]`))
	if err != nil {
		panic(err)
	}
	encoded, err := abi.Pack("BatchDAData", b.BlockNumber, b.Proof, b.Width, b.LeafIndex)
	if err != nil {
		return nil, err
	}
	
	return encoded[4:], nil
}

func (b BatchDAData) IsEmpty() bool {
	return reflect.DeepEqual(b, BatchDAData{})
}
