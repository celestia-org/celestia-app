package app_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/celestiaorg/celestia-app/v2/app"
	"github.com/celestiaorg/celestia-app/v2/app/encoding"
	v1 "github.com/celestiaorg/celestia-app/v2/pkg/appconsts/v1"
	v2 "github.com/celestiaorg/celestia-app/v2/pkg/appconsts/v2"
	"github.com/celestiaorg/celestia-app/v2/test/util"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	icahosttypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	version "github.com/tendermint/tendermint/proto/tendermint/version"
	dbm "github.com/tendermint/tm-db"
)

// TestICA verifies that the ICA module's params are overridden during an
// upgrade from v1 -> v2.
func TestICA(t *testing.T) {
	testApp, _ := setupTestApp(t, 3)
	ctx := testApp.NewContext(true, tmproto.Header{
		Version: version.Consensus{
			App: 1,
		},
	})
	testApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		Height:  2,
		Version: version.Consensus{App: 1},
	}})
	require.EqualValues(t, 1, testApp.AppVersion())

	// Query the ICA host module params
	gotBefore, err := testApp.ParamsKeeper.Params(ctx, &proposal.QueryParamsRequest{
		Subspace: icahosttypes.SubModuleName,
		Key:      string(icahosttypes.KeyHostEnabled),
	})
	require.NoError(t, err)
	require.Equal(t, "", gotBefore.Param.Value)

	// Upgrade from v1 -> v2
	testApp.EndBlock(abci.RequestEndBlock{Height: 2})
	testApp.Commit()
	require.EqualValues(t, 2, testApp.AppVersion())

	newCtx := testApp.NewContext(true, tmproto.Header{Version: version.Consensus{App: 2}})
	got, err := testApp.ParamsKeeper.Params(newCtx, &proposal.QueryParamsRequest{
		Subspace: icahosttypes.SubModuleName,
		Key:      string(icahosttypes.KeyHostEnabled),
	})
	require.NoError(t, err)
	require.Equal(t, "true", got.Param.Value)
}

func setupTestApp(t *testing.T, upgradeHeight int64) (*app.App, keyring.Keyring) {
	t.Helper()

	db := dbm.NewMemDB()
	chainID := "test_chain"
	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	testApp := app.New(log.NewNopLogger(), db, nil, true, 0, encCfg, upgradeHeight, util.EmptyAppOptions{})
	genesisState, _, kr := util.GenesisStateWithSingleValidator(testApp, "account")
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)
	infoResp := testApp.Info(abci.RequestInfo{})
	require.EqualValues(t, 0, infoResp.AppVersion)
	cp := app.DefaultInitialConsensusParams()
	abciParams := &abci.ConsensusParams{
		Block: &abci.BlockParams{
			MaxBytes: cp.Block.MaxBytes,
			MaxGas:   cp.Block.MaxGas,
		},
		Evidence:  &cp.Evidence,
		Validator: &cp.Validator,
		Version:   &cp.Version,
	}

	_ = testApp.InitChain(
		abci.RequestInitChain{
			Time:            time.Now(),
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: abciParams,
			AppStateBytes:   stateBytes,
			ChainId:         chainID,
		},
	)

	// assert that the chain starts with version provided in genesis
	infoResp = testApp.Info(abci.RequestInfo{})
	require.EqualValues(t, app.DefaultInitialConsensusParams().Version.AppVersion, infoResp.AppVersion)

	_ = testApp.Commit()
	supportedVersions := []uint64{v1.Version, v2.Version}
	require.Equal(t, supportedVersions, testApp.SupportedVersions())
	return testApp, kr
}
