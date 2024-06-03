package user_test

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/rand"

	"github.com/celestiaorg/celestia-app/v2/app"
	"github.com/celestiaorg/celestia-app/v2/app/encoding"
	"github.com/celestiaorg/celestia-app/v2/pkg/user"
	"github.com/celestiaorg/celestia-app/v2/test/util/blobfactory"
	"github.com/celestiaorg/celestia-app/v2/test/util/testnode"
)

func TestTxClientTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	suite.Run(t, new(TxClientTestSuite))
}

type TxClientTestSuite struct {
	suite.Suite

	ctx      testnode.Context
	encCfg   encoding.Config
	txClient *user.TxClient
}

func (suite *TxClientTestSuite) SetupSuite() {
	suite.encCfg = encoding.MakeConfig(app.ModuleEncodingRegisters...)
	suite.ctx, _, _ = testnode.NewNetwork(suite.T(), testnode.DefaultConfig().WithFundedAccounts("a"))
	_, err := suite.ctx.WaitForHeight(1)
	suite.Require().NoError(err)
	suite.txClient, err = user.SetupTxClient(suite.ctx.GoContext(), suite.ctx.Keyring, suite.ctx.GRPCClient, suite.encCfg, user.WithGasMultiplier(1.2))
	suite.Require().NoError(err)
}

func (suite *TxClientTestSuite) TestSubmitPayForBlob() {
	t := suite.T()
	blobs := blobfactory.ManyRandBlobs(rand.NewRand(), 1e3, 1e4)

	subCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("submit blob without provided fee and gas limit", func(t *testing.T) {
		resp, err := suite.txClient.SubmitPayForBlob(subCtx, blobs)
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
		require.Greater(t, resp.GasWanted, int64(0))
	})

	t.Run("submit blob with provided fee and gas limit", func(t *testing.T) {
		fee := user.SetFee(1e6)
		gas := user.SetGasLimit(1e6)
		resp, err := suite.txClient.SubmitPayForBlob(subCtx, blobs, fee, gas)
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
		require.EqualValues(t, resp.GasWanted, 1e6)
	})
}

func (suite *TxClientTestSuite) TestSubmitTx() {
	t := suite.T()
	fee := user.SetFee(1e6)
	gas := user.SetGasLimit(1e6)
	addr := suite.txClient.DefaultAddress()
	msg := bank.NewMsgSend(addr, testnode.RandomAddress().(sdk.AccAddress), sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 10)))

	t.Run("submit tx without provided fee and gas limit", func(t *testing.T) {
		resp, err := suite.txClient.SubmitTx(suite.ctx.GoContext(), []sdk.Msg{msg})
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
		require.Greater(t, resp.GasWanted, int64(0))
	})

	t.Run("submit tx with provided gas limit", func(t *testing.T) {
		resp, err := suite.txClient.SubmitTx(suite.ctx.GoContext(), []sdk.Msg{msg}, gas)
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
		require.EqualValues(t, resp.GasWanted, 1e6)
	})

	t.Run("submit tx with provided fee", func(t *testing.T) {
		resp, err := suite.txClient.SubmitTx(suite.ctx.GoContext(), []sdk.Msg{msg}, fee)
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
	})

	t.Run("submit tx with provided fee and gas limit", func(t *testing.T) {
		resp, err := suite.txClient.SubmitTx(suite.ctx.GoContext(), []sdk.Msg{msg}, fee, gas)
		require.NoError(t, err)
		require.EqualValues(t, 0, resp.Code)
		require.EqualValues(t, resp.GasWanted, 1e6)
	})
}

func (suite *TxClientTestSuite) TestConfirmTx() {
	t := suite.T()

	fee := user.SetFee(1e6)
	gas := user.SetGasLimit(1e6)

	t.Run("deadline exceeded when the context times out", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(suite.ctx.GoContext(), time.Second)
		defer cancel()
		_, err := suite.txClient.ConfirmTx(ctx, "E32BD15CAF57AF15D17B0D63CF4E63A9835DD1CEBB059C335C79586BC3013728")
		require.Error(t, err)
		require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
	})

	t.Run("should error when tx is not found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(suite.ctx.GoContext(), 5*time.Second)
		defer cancel()
		_, err := suite.txClient.ConfirmTx(ctx, "not found tx")
		require.Error(t, err)
	})

	t.Run("should success when tx is found immediately", func(t *testing.T) {
		addr := suite.txClient.DefaultAddress()
		msg := bank.NewMsgSend(addr, testnode.RandomAddress().(sdk.AccAddress), sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 10)))
		resp, err := suite.txClient.BroadcastTx(suite.ctx.GoContext(), []sdk.Msg{msg}, fee, gas)
		require.NoError(t, err)
		require.NotNil(t, resp)
		ctx, cancel := context.WithTimeout(suite.ctx.GoContext(), 30*time.Second)
		defer cancel()
		resp, err = suite.txClient.ConfirmTx(ctx, resp.TxHash)
		require.NoError(t, err)
		require.Equal(t, abci.CodeTypeOK, resp.Code)
	})

	t.Run("should error when tx is found with a non-zero error code", func(t *testing.T) {
		balance := suite.queryCurrentBalance(t)
		addr := suite.txClient.DefaultAddress()
		// Create a msg send with out of balance, ensure this tx fails
		msg := bank.NewMsgSend(addr, testnode.RandomAddress().(sdk.AccAddress), sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 1+balance)))
		resp, err := suite.txClient.BroadcastTx(suite.ctx.GoContext(), []sdk.Msg{msg}, fee, gas)
		require.NoError(t, err)
		require.NotNil(t, resp)
		resp, err = suite.txClient.ConfirmTx(suite.ctx.GoContext(), resp.TxHash)
		require.Error(t, err)
		require.NotEqual(t, abci.CodeTypeOK, resp.Code)
	})
}

func (suite *TxClientTestSuite) TestGasEstimation() {
	addr := suite.txClient.DefaultAddress()
	msg := bank.NewMsgSend(addr, testnode.RandomAddress().(sdk.AccAddress), sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 10)))
	gas, err := suite.txClient.EstimateGas(suite.ctx.GoContext(), []sdk.Msg{msg})
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), gas, uint64(0))
}

// TestGasConsumption verifies that the amount deducted from a user's balance is
// based on the fee provided in the tx instead of the gas used by the tx. This
// behavior leads to poor UX because tx submitters must over-estimate the amount
// of gas that their tx will consume and they are not refunded for the excessuite.
func (suite *TxClientTestSuite) TestGasConsumption() {
	t := suite.T()

	utiaToSend := int64(1)
	addr := suite.txClient.DefaultAddress()
	msg := bank.NewMsgSend(addr, testnode.RandomAddress().(sdk.AccAddress), sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, utiaToSend)))

	gasPrice := int64(1)
	gasLimit := uint64(1e6)
	fee := uint64(1e6) // 1 TIA
	// Note: gas price * gas limit = fee amount. So by setting gasLimit and fee
	// to the same value, these options set a gas price of 1utia.
	options := []user.TxOption{user.SetGasLimit(gasLimit), user.SetFee(fee)}

	balanceBefore := suite.queryCurrentBalance(t)
	resp, err := suite.txClient.SubmitTx(suite.ctx.GoContext(), []sdk.Msg{msg}, options...)
	require.NoError(t, err)

	require.EqualValues(t, abci.CodeTypeOK, resp.Code)
	balanceAfter := suite.queryCurrentBalance(t)

	// verify that the amount deducted depends on the fee set in the tx.
	amountDeducted := balanceBefore - balanceAfter - utiaToSend
	require.Equal(t, int64(fee), amountDeducted)

	// verify that the amount deducted does not depend on the actual gas used.
	gasUsedBasedDeduction := resp.GasUsed * gasPrice
	require.NotEqual(t, gasUsedBasedDeduction, amountDeducted)
	// The gas used based deduction should be less than the fee because the fee is 1 TIA.
	require.Less(t, gasUsedBasedDeduction, int64(fee))
}

func (suite *TxClientTestSuite) queryCurrentBalance(t *testing.T) int64 {
	balanceQuery := bank.NewQueryClient(suite.ctx.GRPCClient)
	addr := suite.txClient.DefaultAddress()
	balanceResp, err := balanceQuery.AllBalances(suite.ctx.GoContext(), &bank.QueryAllBalancesRequest{Address: addr.String()})
	require.NoError(t, err)
	return balanceResp.Balances.AmountOf(app.BondDenom).Int64()
}
