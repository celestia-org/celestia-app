package errors_test

import (
	"fmt"
	"testing"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	apperr "github.com/celestiaorg/celestia-app/app/errors"
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/celestia-app/pkg/user"
	testutil "github.com/celestiaorg/celestia-app/test/util"
	"github.com/celestiaorg/celestia-app/test/util/testfactory"
	blob "github.com/celestiaorg/celestia-app/x/blob/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// This will detect any changes to the DeductFeeDecorator which may cause a
// different error message that does not match the regexp.
func TestNonceMismatchIntegration(t *testing.T) {
	account := "test"
	testApp, kr := testutil.SetupTestAppWithGenesisValSet(app.DefaultConsensusParams(), account)
	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	minGasPrice, err := sdk.ParseDecCoins(fmt.Sprintf("%v%s", appconsts.DefaultMinGasPrice, app.BondDenom))
	require.NoError(t, err)
	ctx := testApp.NewContext(true, tmproto.Header{}).WithMinGasPrices(minGasPrice)
	addr := testfactory.GetAddress(kr, account)
	enc := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	acc := testutil.DirectQueryAccount(testApp, addr)
	// set the sequence to an incorrect value
	signer, err := user.NewSigner(kr, nil, addr, enc.TxConfig, testutil.ChainID, acc.GetAccountNumber(), 2)
	require.NoError(t, err)

	b, err := blob.NewBlob(namespace.RandomNamespace(), []byte("hello world"), 0)
	require.NoError(t, err)

	msg, err := blob.NewMsgPayForBlobs(signer.Address().String(), b)
	require.NoError(t, err)

	tx, err := signer.CreateTx([]sdk.Msg{msg})
	require.NoError(t, err)
	sdkTx, err := enc.TxConfig.TxDecoder()(tx)
	require.NoError(t, err)

	decorator := ante.NewSigVerificationDecorator(testApp.AccountKeeper, encCfg.TxConfig.SignModeHandler())
	anteHandler := sdk.ChainAnteDecorators(decorator)

	// We set simulate to true here to bypass having to initialize the
	// accounts public key.
	_, err = anteHandler(ctx, sdkTx, true)
	require.True(t, apperr.IsNonceMismatch(err), err)
	expectedNonce, err := apperr.ParseNonceMismatch(err)
	require.NoError(t, err)
	require.EqualValues(t, 0, expectedNonce, err)
}
