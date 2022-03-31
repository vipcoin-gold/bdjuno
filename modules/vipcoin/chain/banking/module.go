/*
 * Copyright 2022 Business Process Technologies. All rights reserved.
 */

package banking

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/juno/v2/modules"

	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/modules/vipcoin/chain/banking/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/banking module
type Module struct {
	cdc    codec.Marshaler
	db     *database.Db
	keeper source.Source
}

// NewModule returns a new Module instance
func NewModule(
	keeper source.Source, cdc codec.Marshaler, db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		keeper: keeper,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vipcoin_banking"
}