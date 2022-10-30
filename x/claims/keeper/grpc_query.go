package keeper

import (
	"neware/x/claims/types"
)

var _ types.QueryServer = Keeper{}
