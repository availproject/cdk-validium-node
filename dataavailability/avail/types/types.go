// SPDX-License-Identifier: Apache-2.0
package types

import (
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type BatchDAData struct {
	BlockNumber uint
	LeafIndex   uint
}

// write a function that encode batchDAData struct into ABI-encoded bytes
func (b BatchDAData) Encode() ([]byte, error) {
	abi, err := abi.JSON(strings.NewReader(`[{"type":"uint256","name":"blockNumber"},{"type":"uint256","name":"leafIndex"}]`))
	if err != nil {
		return nil, err
	}
	return abi.Pack("BatchDAData", b.BlockNumber, b.LeafIndex)
}

func (b BatchDAData) IsEmpty() bool {
	return reflect.DeepEqual(b, BatchDAData{})
}
