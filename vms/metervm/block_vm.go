// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/lasthyphen/avalanchego-1.7.3/api/metrics"
	"github.com/lasthyphen/avalanchego-1.7.3/database/manager"
	"github.com/lasthyphen/avalanchego-1.7.3/ids"
	"github.com/lasthyphen/avalanchego-1.7.3/snow"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowman"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/common"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/snowman/block"
	"github.com/lasthyphen/avalanchego-1.7.3/utils/timer/mockable"
)

var _ block.ChainVM = &blockVM{}

func NewBlockVM(vm block.ChainVM) block.ChainVM {
	return &blockVM{
		ChainVM: vm,
	}
}

type blockVM struct {
	block.ChainVM
	blockMetrics
	clock mockable.Clock
}

func (vm *blockVM) Initialize(
	ctx *snow.Context,
	db manager.Manager,
	genesisBytes,
	upgradeBytes,
	configBytes []byte,
	toEngine chan<- common.Message,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	registerer := prometheus.NewRegistry()
	_, supportsBatchedFetching := vm.ChainVM.(block.BatchedChainVM)
	if err := vm.blockMetrics.Initialize(supportsBatchedFetching, "", registerer); err != nil {
		return err
	}

	optionalGatherer := metrics.NewOptionalGatherer()
	multiGatherer := metrics.NewMultiGatherer()
	if err := multiGatherer.Register("metervm", registerer); err != nil {
		return err
	}
	if err := multiGatherer.Register("", optionalGatherer); err != nil {
		return err
	}
	if err := ctx.Metrics.Register(multiGatherer); err != nil {
		return err
	}
	ctx.Metrics = optionalGatherer

	return vm.ChainVM.Initialize(ctx, db, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
}

func (vm *blockVM) BuildBlock() (snowman.Block, error) {
	start := vm.clock.Time()
	blk, err := vm.ChainVM.BuildBlock()
	end := vm.clock.Time()
	duration := float64(end.Sub(start))
	if err != nil {
		vm.blockMetrics.buildBlockErr.Observe(duration)
		return nil, err
	}
	vm.blockMetrics.buildBlock.Observe(duration)
	return &meterBlock{
		Block: blk,
		vm:    vm,
	}, nil
}

func (vm *blockVM) ParseBlock(b []byte) (snowman.Block, error) {
	start := vm.clock.Time()
	blk, err := vm.ChainVM.ParseBlock(b)
	end := vm.clock.Time()
	duration := float64(end.Sub(start))
	if err != nil {
		vm.blockMetrics.parseBlockErr.Observe(duration)
		return nil, err
	}
	vm.blockMetrics.parseBlock.Observe(duration)
	return &meterBlock{
		Block: blk,
		vm:    vm,
	}, nil
}

func (vm *blockVM) GetBlock(id ids.ID) (snowman.Block, error) {
	start := vm.clock.Time()
	blk, err := vm.ChainVM.GetBlock(id)
	end := vm.clock.Time()
	duration := float64(end.Sub(start))
	if err != nil {
		vm.blockMetrics.getBlockErr.Observe(duration)
		return nil, err
	}
	vm.blockMetrics.getBlock.Observe(duration)
	return &meterBlock{
		Block: blk,
		vm:    vm,
	}, nil
}

func (vm *blockVM) SetPreference(id ids.ID) error {
	start := vm.clock.Time()
	err := vm.ChainVM.SetPreference(id)
	end := vm.clock.Time()
	vm.blockMetrics.setPreference.Observe(float64(end.Sub(start)))
	return err
}

func (vm *blockVM) LastAccepted() (ids.ID, error) {
	start := vm.clock.Time()
	lastAcceptedID, err := vm.ChainVM.LastAccepted()
	end := vm.clock.Time()
	vm.blockMetrics.lastAccepted.Observe(float64(end.Sub(start)))
	return lastAcceptedID, err
}
