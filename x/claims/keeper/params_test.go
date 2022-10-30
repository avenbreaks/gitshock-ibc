package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "neware/testutil/keeper"
	"neware/x/claims/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ClaimsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
