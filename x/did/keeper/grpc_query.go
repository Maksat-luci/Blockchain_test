package keeper

import (
	"spider/x/did/types"
)

var _ types.QueryServer = Keeper{}
