package keeper

import (
	"neware/x/dex/types"
)

var _ types.QueryServer = Keeper{}
