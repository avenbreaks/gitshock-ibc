package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "neware/testutil/keeper"
	"neware/x/neware/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewareKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
