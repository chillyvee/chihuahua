package chihuahua_test

import (
	"testing"

	keepertest "github.com/ChihuahuaChain/chihuahua/testutil/keeper"
	"github.com/ChihuahuaChain/chihuahua/x/chihuahua"
	"github.com/ChihuahuaChain/chihuahua/x/chihuahua/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChihuahuaKeeper(t)
	chihuahua.InitGenesis(ctx, *k, genesisState)
	got := chihuahua.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
