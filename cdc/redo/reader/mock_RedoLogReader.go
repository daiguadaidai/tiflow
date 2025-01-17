// Code generated by mockery v2.14.1. DO NOT EDIT.

package reader

import (
	context "context"

	model "github.com/pingcap/tiflow/cdc/model"
	mock "github.com/stretchr/testify/mock"
)

// MockRedoLogReader is an autogenerated mock type for the RedoLogReader type
type MockRedoLogReader struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockRedoLogReader) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadMeta provides a mock function with given fields: ctx
func (_m *MockRedoLogReader) ReadMeta(ctx context.Context) (uint64, uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context) uint64); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReadNextDDL provides a mock function with given fields: ctx, maxNumberOfEvents
func (_m *MockRedoLogReader) ReadNextDDL(ctx context.Context, maxNumberOfEvents uint64) ([]*model.RedoDDLEvent, error) {
	ret := _m.Called(ctx, maxNumberOfEvents)

	var r0 []*model.RedoDDLEvent
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []*model.RedoDDLEvent); ok {
		r0 = rf(ctx, maxNumberOfEvents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RedoDDLEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, maxNumberOfEvents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadNextLog provides a mock function with given fields: ctx, maxNumberOfEvents
func (_m *MockRedoLogReader) ReadNextLog(ctx context.Context, maxNumberOfEvents uint64) ([]*model.RedoRowChangedEvent, error) {
	ret := _m.Called(ctx, maxNumberOfEvents)

	var r0 []*model.RedoRowChangedEvent
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []*model.RedoRowChangedEvent); ok {
		r0 = rf(ctx, maxNumberOfEvents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RedoRowChangedEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, maxNumberOfEvents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetReader provides a mock function with given fields: ctx, startTs, endTs
func (_m *MockRedoLogReader) ResetReader(ctx context.Context, startTs uint64, endTs uint64) error {
	ret := _m.Called(ctx, startTs, endTs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, startTs, endTs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockRedoLogReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRedoLogReader creates a new instance of MockRedoLogReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRedoLogReader(t mockConstructorTestingTNewMockRedoLogReader) *MockRedoLogReader {
	mock := &MockRedoLogReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
