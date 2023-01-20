// Code generated by mockery v2.16.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// DbManagerMock is an autogenerated mock type for the dbManagerInterface type
type DbManagerMock struct {
	mock.Mock
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *DbManagerMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseBatch provides a mock function with given fields: ctx, params
func (_m *DbManagerMock) CloseBatch(ctx context.Context, params ClosingBatchParameters) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ClosingBatchParameters) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFirstBatch provides a mock function with given fields: ctx, sequencerAddress
func (_m *DbManagerMock) CreateFirstBatch(ctx context.Context, sequencerAddress common.Address) state.ProcessingContext {
	ret := _m.Called(ctx, sequencerAddress)

	var r0 state.ProcessingContext
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) state.ProcessingContext); ok {
		r0 = rf(ctx, sequencerAddress)
	} else {
		r0 = ret.Get(0).(state.ProcessingContext)
	}

	return r0
}

// DeleteTransactionFromPool provides a mock function with given fields: ctx, txHash
func (_m *DbManagerMock) DeleteTransactionFromPool(ctx context.Context, txHash common.Hash) error {
	ret := _m.Called(ctx, txHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) error); ok {
		r0 = rf(ctx, txHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalanceByStateRoot provides a mock function with given fields: ctx, address, root
func (_m *DbManagerMock) GetBalanceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatchesSince provides a mock function with given fields: ctx, forcedBatchNumber, dbTx
func (_m *DbManagerMock) GetForcedBatchesSince(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) ([]*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, dbTx)

	var r0 []*state.ForcedBatch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []*state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.ForcedBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastBatch(ctx context.Context) (*state.Batch, error) {
	ret := _m.Called(ctx)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context) *state.Batch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastBatchNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastClosedBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetLastClosedBatch(ctx context.Context) (*state.Batch, error) {
	ret := _m.Called(ctx)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context) *state.Batch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2BlockHeader provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLastL2BlockHeader(ctx context.Context, dbTx pgx.Tx) (*types.Header, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *types.Header); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNBatches provides a mock function with given fields: ctx, numBatches
func (_m *DbManagerMock) GetLastNBatches(ctx context.Context, numBatches uint) ([]*state.Batch, error) {
	ret := _m.Called(ctx, numBatches)

	var r0 []*state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*state.Batch); ok {
		r0 = rf(ctx, numBatches)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, numBatches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTrustedForcedBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *DbManagerMock) GetLastTrustedForcedBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestGer provides a mock function with given fields: ctx, gerFinalityNumberOfBlocks
func (_m *DbManagerMock) GetLatestGer(ctx context.Context, gerFinalityNumberOfBlocks uint64) (state.GlobalExitRoot, time.Time, error) {
	ret := _m.Called(ctx, gerFinalityNumberOfBlocks)

	var r0 state.GlobalExitRoot
	if rf, ok := ret.Get(0).(func(context.Context, uint64) state.GlobalExitRoot); ok {
		r0 = rf(ctx, gerFinalityNumberOfBlocks)
	} else {
		r0 = ret.Get(0).(state.GlobalExitRoot)
	}

	var r1 time.Time
	if rf, ok := ret.Get(1).(func(context.Context, uint64) time.Time); ok {
		r1 = rf(ctx, gerFinalityNumberOfBlocks)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, gerFinalityNumberOfBlocks)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWIPBatch provides a mock function with given fields: ctx
func (_m *DbManagerMock) GetWIPBatch(ctx context.Context) (*WipBatch, error) {
	ret := _m.Called(ctx)

	var r0 *WipBatch
	if rf, ok := ret.Get(0).(func(context.Context) *WipBatch); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*WipBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBatchClosed provides a mock function with given fields: ctx, batchNum
func (_m *DbManagerMock) IsBatchClosed(ctx context.Context, batchNum uint64) (bool, error) {
	ret := _m.Called(ctx, batchNum)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint64) bool); ok {
		r0 = rf(ctx, batchNum)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, batchNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkReorgedTxsAsPending provides a mock function with given fields: ctx
func (_m *DbManagerMock) MarkReorgedTxsAsPending(ctx context.Context) {
	_m.Called(ctx)
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *DbManagerMock) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessForcedBatch provides a mock function with given fields: forcedBatchNum, request
func (_m *DbManagerMock) ProcessForcedBatch(forcedBatchNum uint64, request state.ProcessRequest) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(forcedBatchNum, request)

	var r0 *state.ProcessBatchResponse
	if rf, ok := ret.Get(0).(func(uint64, state.ProcessRequest) *state.ProcessBatchResponse); ok {
		r0 = rf(forcedBatchNum, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, state.ProcessRequest) error); ok {
		r1 = rf(forcedBatchNum, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreProcessedTransaction provides a mock function with given fields: ctx, batchNumber, processedTx, coinbase, timestamp, dbTx
func (_m *DbManagerMock) StoreProcessedTransaction(ctx context.Context, batchNumber uint64, processedTx *state.ProcessTransactionResponse, coinbase common.Address, timestamp uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessTransactionResponse, common.Address, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDbManagerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewDbManagerMock creates a new instance of DbManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDbManagerMock(t mockConstructorTestingTNewDbManagerMock) *DbManagerMock {
	mock := &DbManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}