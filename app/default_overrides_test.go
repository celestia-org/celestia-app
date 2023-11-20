package app

import (
	"testing"
	"time"

	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/assert"
	tmcfg "github.com/tendermint/tendermint/config"
)

// Test_newGovModule verifies that the gov module's genesis state has defaults
// overridden.
func Test_newGovModule(t *testing.T) {
	encCfg := encoding.MakeConfig(ModuleEncodingRegisters...)
	day := time.Duration(time.Hour * 24)
	oneWeek := day * 7

	govModule := newGovModule()
	raw := govModule.DefaultGenesis(encCfg.Codec)
	govGenesisState := govtypes.GenesisState{}

	encCfg.Codec.MustUnmarshalJSON(raw, &govGenesisState)

	want := []types.Coin{{
		Denom:  BondDenom,
		Amount: types.NewInt(10_000_000_000),
	}}

	assert.Equal(t, want, govGenesisState.DepositParams.MinDeposit)
	assert.Equal(t, oneWeek, *govGenesisState.DepositParams.MaxDepositPeriod)
	assert.Equal(t, oneWeek, *govGenesisState.VotingParams.VotingPeriod)
}

// TestDefaultGenesis verifies that the distribution module's genesis state has
// defaults overridden.
func TestDefaultGenesis(t *testing.T) {
	encCfg := encoding.MakeConfig(ModuleEncodingRegisters...)
	dm := distributionModule{}
	raw := dm.DefaultGenesis(encCfg.Codec)
	distributionGenesisState := distributiontypes.GenesisState{}
	encCfg.Codec.MustUnmarshalJSON(raw, &distributionGenesisState)

	// Verify that BaseProposerReward and BonusProposerReward were overridden to 0%.
	assert.Equal(t, types.ZeroDec(), distributionGenesisState.Params.BaseProposerReward)
	assert.Equal(t, types.ZeroDec(), distributionGenesisState.Params.BonusProposerReward)

	// Verify that other params weren't overridden.
	assert.Equal(t, distributiontypes.DefaultParams().WithdrawAddrEnabled, distributionGenesisState.Params.WithdrawAddrEnabled)
	assert.Equal(t, distributiontypes.DefaultParams().CommunityTax, distributionGenesisState.Params.CommunityTax)
}

func TestDefaultAppConfig(t *testing.T) {
	cfg := DefaultAppConfig()

	assert.False(t, cfg.API.Enable)
	assert.False(t, cfg.GRPC.Enable)
	assert.False(t, cfg.GRPCWeb.Enable)

	assert.Equal(t, uint64(1500), cfg.StateSync.SnapshotInterval)
	assert.Equal(t, uint32(2), cfg.StateSync.SnapshotKeepRecent)
	assert.Equal(t, "0.002utia", cfg.MinGasPrices)
}

func TestDefaultConsensusConfig(t *testing.T) {
	got := DefaultConsensusConfig()

	t.Run("mempool overrides", func(t *testing.T) {
		want := tmcfg.MempoolConfig{
			// defaults
			Broadcast:             tmcfg.DefaultMempoolConfig().Broadcast,
			CacheSize:             tmcfg.DefaultMempoolConfig().CacheSize,
			KeepInvalidTxsInCache: tmcfg.DefaultMempoolConfig().KeepInvalidTxsInCache,
			Recheck:               tmcfg.DefaultMempoolConfig().Recheck,
			RootDir:               tmcfg.DefaultMempoolConfig().RootDir,
			Size:                  tmcfg.DefaultMempoolConfig().Size,
			WalPath:               tmcfg.DefaultMempoolConfig().WalPath,

			// overrides
			MaxTxBytes:   7_897_088,
			MaxTxsBytes:  39_485_440,
			TTLDuration:  75 * time.Second,
			TTLNumBlocks: 5,
			Version:      "v1",
		}
		assert.Equal(t, want, *got.Mempool)
	})
}
