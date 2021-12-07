package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ChihuahuaChain/chihuahua/testutil/keeper"
	"github.com/ChihuahuaChain/chihuahua/x/chihuahua/keeper"
	"github.com/ChihuahuaChain/chihuahua/x/chihuahua/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChihuahuaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
