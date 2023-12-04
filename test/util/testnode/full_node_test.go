package testnode

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	appns "github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/celestia-app/test/util/genesis"
	blobtypes "github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func TestIntegrationTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping full node integration test in short mode.")
	}
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite

	accounts []string
	cctx     Context
}

func (s *IntegrationTestSuite) SetupSuite() {
	t := s.T()
	s.accounts = RandomAccounts(10)

	ecfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	blobGenState := blobtypes.DefaultGenesis()
	blobGenState.Params.GovMaxSquareSize = uint64(appconsts.DefaultSquareSizeUpperBound)

	cfg := DefaultConfig().
		WithFundedAccounts(s.accounts...).
		WithModifiers(genesis.SetBlobParams(ecfg.Codec, blobGenState.Params))

	cctx, _, _ := NewNetwork(t, cfg)
	s.cctx = cctx
}

func (s *IntegrationTestSuite) Test_verifyTimeIotaMs() {
	require := s.Require()
	err := s.cctx.WaitForNextBlock()
	require.NoError(err)

	var params *coretypes.ResultConsensusParams
	// this query can be flaky with fast block times, so we repeat it multiple
	// times in attempt to decrease flakiness
	for i := 0; i < 100; i++ {
		params, err = s.cctx.Client.ConsensusParams(context.Background(), nil)
		if err == nil && params != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	require.NoError(err)
	require.NotNil(params)
	require.Equal(int64(1), params.ConsensusParams.Block.TimeIotaMs)
}

func (s *IntegrationTestSuite) TestPostData() {
	require := s.Require()
	_, err := s.cctx.PostData(s.accounts[0], flags.BroadcastBlock, appns.RandomBlobNamespace(), tmrand.Bytes(kibibyte))
	require.NoError(err)
}

func (s *IntegrationTestSuite) TestFillBlock() {
	require := s.Require()

	for squareSize := 2; squareSize <= appconsts.DefaultGovMaxSquareSize; squareSize *= 2 {
		resp, err := s.cctx.FillBlock(squareSize, s.accounts[1], flags.BroadcastSync)
		require.NoError(err)

		err = s.cctx.WaitForBlocks(3)
		require.NoError(err, squareSize)

		res, err := QueryWithoutProof(s.cctx.Context, resp.TxHash)
		require.NoError(err, squareSize)
		require.Equal(abci.CodeTypeOK, res.TxResult.Code, squareSize)

		b, err := s.cctx.Client.Block(s.cctx.GoContext(), &res.Height)
		require.NoError(err, squareSize)
		require.Equal(uint64(squareSize), b.Block.SquareSize, squareSize)
	}
}

func (s *IntegrationTestSuite) TestFillBlock_InvalidSquareSizeError() {
	tests := []struct {
		name        string
		squareSize  int
		expectedErr error
	}{
		{
			name:        "when squareSize less than 2",
			squareSize:  0,
			expectedErr: fmt.Errorf("unsupported squareSize: 0"),
		},
		{
			name:        "when squareSize is greater than 2 but not a power of 2",
			squareSize:  18,
			expectedErr: fmt.Errorf("unsupported squareSize: 18"),
		},
		{
			name:       "when squareSize is greater than 2 and a power of 2",
			squareSize: 16,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			_, actualErr := s.cctx.FillBlock(tc.squareSize, s.accounts[2], flags.BroadcastAsync)
			s.Equal(tc.expectedErr, actualErr)
		})
	}
}
