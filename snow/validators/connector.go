// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/lasthyphen/avalanchego-1.7.3/ids"
	"github.com/lasthyphen/avalanchego-1.7.3/version"
)

// Connector represents a handler that is called when a connection is marked as
// connected or disconnected
type Connector interface {
	Connected(id ids.ShortID, nodeVersion version.Application) error
	Disconnected(id ids.ShortID) error
}
