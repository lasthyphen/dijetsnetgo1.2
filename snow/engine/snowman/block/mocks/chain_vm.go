// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	ids "github.com/lasthyphen/avalanchego-1.7.3/ids"
	common "github.com/lasthyphen/avalanchego-1.7.3/snow/engine/common"

	manager "github.com/lasthyphen/avalanchego-1.7.3/database/manager"

	mock "github.com/stretchr/testify/mock"

	snow "github.com/lasthyphen/avalanchego-1.7.3/snow"

	snowman "github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowman"
)

// ChainVM is an autogenerated mock type for the ChainVM type
type ChainVM struct {
	mock.Mock
}

// Bootstrapped provides a mock function with given fields:
func (_m *ChainVM) Bootstrapped() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Bootstrapping provides a mock function with given fields:
func (_m *ChainVM) Bootstrapping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildBlock provides a mock function with given fields:
func (_m *ChainVM) BuildBlock() (snowman.Block, error) {
	ret := _m.Called()

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func() snowman.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHandlers provides a mock function with given fields:
func (_m *ChainVM) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	ret := _m.Called()

	var r0 map[string]*common.HTTPHandler
	if rf, ok := ret.Get(0).(func() map[string]*common.HTTPHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*common.HTTPHandler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: _a0
func (_m *ChainVM) GetBlock(_a0 ids.ID) (snowman.Block, error) {
	ret := _m.Called(_a0)

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func(ids.ID) snowman.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields:
func (_m *ChainVM) HealthCheck() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs
func (_m *ChainVM) Initialize(ctx *snow.Context, dbManager manager.Manager, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx) error {
	ret := _m.Called(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs)

	var r0 error
	if rf, ok := ret.Get(0).(func(*snow.Context, manager.Manager, []byte, []byte, []byte, chan<- common.Message, []*common.Fx) error); ok {
		r0 = rf(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LastAccepted provides a mock function with given fields:
func (_m *ChainVM) LastAccepted() (ids.ID, error) {
	ret := _m.Called()

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func() ids.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseBlock provides a mock function with given fields: _a0
func (_m *ChainVM) ParseBlock(_a0 []byte) (snowman.Block, error) {
	ret := _m.Called(_a0)

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func([]byte) snowman.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPreference provides a mock function with given fields: _a0
func (_m *ChainVM) SetPreference(_a0 ids.ID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *ChainVM) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
