package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "spider/testutil/keeper"
	"spider/x/did/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DidKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
