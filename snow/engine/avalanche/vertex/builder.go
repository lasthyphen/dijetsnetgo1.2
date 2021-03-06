// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/lasthyphen/dijetsnetgo1.2/ids"
	"github.com/lasthyphen/dijetsnetgo1.2/snow/consensus/avalanche"
	"github.com/lasthyphen/dijetsnetgo1.2/snow/consensus/snowstorm"
	"github.com/lasthyphen/dijetsnetgo1.2/utils/hashing"
)

// Builder builds a vertex given a set of parentIDs and transactions.
type Builder interface {
	// Build a new vertex from the contents of a vertex
	BuildVtx(
		parentIDs []ids.ID,
		txs []snowstorm.Tx,
	) (avalanche.Vertex, error)
}

// Build a new stateless vertex from the contents of a vertex
func Build(
	chainID ids.ID,
	height uint64,
	parentIDs []ids.ID,
	txs [][]byte,
) (StatelessVertex, error) {
	ids.SortIDs(parentIDs)
	SortHashOf(txs)

	innerVtx := innerStatelessVertex{
		Version:   codecVersion,
		ChainID:   chainID,
		Height:    height,
		Epoch:     0,
		ParentIDs: parentIDs,
		Txs:       txs,
	}
	if err := innerVtx.Verify(); err != nil {
		return nil, err
	}

	vtxBytes, err := c.Marshal(innerVtx.Version, innerVtx)
	vtx := statelessVertex{
		innerStatelessVertex: innerVtx,
		id:                   hashing.ComputeHash256Array(vtxBytes),
		bytes:                vtxBytes,
	}
	return vtx, err
}
