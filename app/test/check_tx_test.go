package app_test

import (
	"bytes"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/celestiaorg/celestia-app/pkg/blob"
	appns "github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/celestia-app/pkg/user"
	testutil "github.com/celestiaorg/celestia-app/test/util"
	"github.com/celestiaorg/celestia-app/test/util/blobfactory"
	"github.com/celestiaorg/celestia-app/test/util/testfactory"
	blobtypes "github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	coretypes "github.com/tendermint/tendermint/types"
)

// Here we only need to check the functionality that is added to CheckTx. We
// assume that the rest of CheckTx is tested by the cosmos-sdk.
func TestCheckTx(t *testing.T) {
	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	ns1 := appns.MustNewV0(bytes.Repeat([]byte{1}, appns.NamespaceVersionZeroIDSize))

	accs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	testApp, kr := testutil.SetupTestAppWithGenesisValSet(app.DefaultConsensusParams(), accs...)
	testApp.Commit()

	opts := blobfactory.FeeTxOpts(1e9)

	type test struct {
		name             string
		checkType        abci.CheckTxType
		getTx            func() []byte
		expectedABCICode uint32
	}

	tests := []test{
		{
			name:      "normal transaction, CheckTxType_New",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[0], encCfg.TxConfig, 1)
				btx := blobfactory.RandBlobTxsWithNamespacesAndSigner(
					signer,
					[]appns.Namespace{ns1},
					[]int{100},
				)[0]
				return btx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "normal transaction, CheckTxType_Recheck",
			checkType: abci.CheckTxType_Recheck,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[1], encCfg.TxConfig, 2)
				btx := blobfactory.RandBlobTxsWithNamespacesAndSigner(
					signer,
					[]appns.Namespace{ns1},
					[]int{100},
				)[0]
				return btx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "invalid transaction, mismatched namespace",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[2], encCfg.TxConfig, 3)
				btx := blobfactory.RandBlobTxsWithNamespacesAndSigner(
					signer,
					[]appns.Namespace{ns1},
					[]int{100},
				)[0]

				dtx, _ := blob.UnmarshalBlobTx(btx)
				dtx.Blobs[0].NamespaceId = appns.RandomBlobNamespace().ID
				bbtx, err := blob.MarshalBlobTx(dtx.Tx, dtx.Blobs[0])
				require.NoError(t, err)
				return bbtx
			},
			expectedABCICode: blobtypes.ErrNamespaceMismatch.ABCICode(),
		},
		{
			name:      "PFB with no blob, CheckTxType_New",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[3], encCfg.TxConfig, 4)
				btx := blobfactory.RandBlobTxsWithNamespacesAndSigner(
					signer,
					[]appns.Namespace{ns1},
					[]int{100},
				)[0]
				dtx, _ := coretypes.UnmarshalBlobTx(btx)
				return dtx.Tx
			},
			expectedABCICode: blobtypes.ErrNoBlobs.ABCICode(),
		},
		{
			name:      "normal blobTx w/ multiple blobs, CheckTxType_New",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[4], encCfg.TxConfig, 5)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 10_000, 10)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "1,000 byte blob",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[5], encCfg.TxConfig, 6)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 1_000, 1)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "10,000 byte blob",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[6], encCfg.TxConfig, 7)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 10_000, 1)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "100,000 byte blob",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[7], encCfg.TxConfig, 8)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 100_000, 1)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "1,000,000 byte blob",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[8], encCfg.TxConfig, 9)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 1_000_000, 1)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: abci.CodeTypeOK,
		},
		{
			name:      "10,000,000 byte blob",
			checkType: abci.CheckTxType_New,
			getTx: func() []byte {
				signer := createSigner(t, kr, accs[9], encCfg.TxConfig, 10)
				_, blobs := blobfactory.RandMsgPayForBlobsWithSigner(tmrand.NewRand(), signer.Address().String(), 10_000_000, 1)
				tx, err := signer.CreatePayForBlob(blobs, opts...)
				require.NoError(t, err)
				return tx
			},
			expectedABCICode: blobtypes.ErrTotalBlobSizeTooLarge.ABCICode(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testApp.CheckTx(abci.RequestCheckTx{Type: tt.checkType, Tx: tt.getTx()})
			assert.Equal(t, tt.expectedABCICode, resp.Code, resp.Log)
		})
	}
}

func createSigner(t *testing.T, kr keyring.Keyring, accountName string, enc client.TxConfig, accNum uint64) *user.Signer {
	addr := testfactory.GetAddress(kr, accountName)
	signer, err := user.NewSigner(kr, nil, addr, enc, testutil.ChainID, accNum, 0)
	require.NoError(t, err)
	return signer
}
