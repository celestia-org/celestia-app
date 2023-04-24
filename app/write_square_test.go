package app

import (
	"bytes"
	"fmt"
	"testing"

	appns "github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/celestia-app/pkg/shares"
	"github.com/celestiaorg/celestia-app/test/util/blobfactory"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	coretypes "github.com/tendermint/tendermint/types"
)

func Test_finalizeLayout(t *testing.T) {
	ns1 := appns.MustNewV0(bytes.Repeat([]byte{1}, appns.NamespaceVersionZeroIDSize))
	ns2 := appns.MustNewV0(bytes.Repeat([]byte{2}, appns.NamespaceVersionZeroIDSize))
	ns3 := appns.MustNewV0(bytes.Repeat([]byte{3}, appns.NamespaceVersionZeroIDSize))

	type test struct {
		squareSize      uint64
		nonreserveStart int
		blobTxs         []tmproto.BlobTx
		expectedIndexes [][]uint32
	}
	tests := []test{
		{
			squareSize:      4,
			nonreserveStart: 10,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1},
				[][]int{{1}},
			),
			expectedIndexes: [][]uint32{{10}},
		},
		{
			squareSize:      4,
			nonreserveStart: 10,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns1},
				blobfactory.Repeat([]int{100}, 2),
			),
			expectedIndexes: [][]uint32{{10}, {11}},
		},
		{
			squareSize:      4,
			nonreserveStart: 10,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1},
				blobfactory.Repeat([]int{100}, 10),
			),
			expectedIndexes: [][]uint32{{10}, {11}, {12}, {13}, {14}, {15}},
		},
		{
			squareSize:      4,
			nonreserveStart: 7,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1},
				blobfactory.Repeat([]int{100}, 9),
			),
			expectedIndexes: [][]uint32{{7}, {8}, {9}, {10}, {11}, {12}, {13}, {14}, {15}},
		},
		{
			squareSize:      4,
			nonreserveStart: 3,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns1, ns1},
				[][]int{{10000}, {10000}, {1000000}},
			),
			expectedIndexes: [][]uint32{},
		},
		{
			squareSize:      64,
			nonreserveStart: 32,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns1, ns1},
				[][]int{{1000}, {10000}, {100000}},
			),
			expectedIndexes: [][]uint32{{32}, {35}, {56}},
		},
		{
			squareSize:      32,
			nonreserveStart: 32,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns2, ns1, ns1},
				[][]int{{100}, {100}, {100}},
			),
			expectedIndexes: [][]uint32{{34}, {32}, {33}},
		},
		{
			squareSize:      32,
			nonreserveStart: 32,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns2, ns1},
				[][]int{{100}, {900}, {900}}, // 1, 2, 2 shares respectively
			),
			expectedIndexes: [][]uint32{{32}, {35}, {33}},
		},
		{
			squareSize:      4,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns3, ns3, ns2},
				[][]int{{100}, {1000, 1000}, {420}},
			),
			expectedIndexes: [][]uint32{{4}, {6, 9}, {5}},
		},
		{
			// no blob txs should make it in the square
			squareSize:      2,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns2, ns3},
				[][]int{{1000}, {1000}, {1000}},
			),
			expectedIndexes: [][]uint32{},
		},
		{
			// only one blob tx should make it in the square (after reordering)
			squareSize:      4,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns3, ns2, ns1},
				[][]int{{2000}, {2000}, {5000}},
			),
			expectedIndexes: [][]uint32{{4}},
		},
		{
			squareSize:      4,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns3, ns3, ns2, ns1},
				[][]int{{1800, 1000}, {6000}, {1800}},
			),
			// should be ns1 and {ns3, ns3} as ns2 is too large
			expectedIndexes: [][]uint32{{8, 12}, {4}},
		},
		{
			squareSize:      4,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns3, ns3, ns1, ns2, ns2},
				[][]int{{100}, {1400, 900, 200, 200}, {420}},
			),
			expectedIndexes: [][]uint32{{4}, {8, 11, 5, 6}, {7}},
		},
		{
			squareSize:      4,
			nonreserveStart: 4,
			blobTxs: generateBlobTxsWithNamespaces(
				t,
				[]appns.Namespace{ns1, ns3, ns3, ns1, ns2, ns2},
				[][]int{{100}, {900, 1400, 200, 200}, {420}},
			),
			expectedIndexes: [][]uint32{{4}, {8, 10, 5, 6}, {7}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			wrappedPFBs, blobs := finalizeBlobLayout(tt.squareSize, tt.nonreserveStart, tt.blobTxs)
			require.Equal(t, len(wrappedPFBs), len(tt.expectedIndexes))
			for j, pfbBytes := range wrappedPFBs {
				wrappedPFB, isWrappedPFB := coretypes.UnmarshalIndexWrapper(pfbBytes)
				require.True(t, isWrappedPFB)
				require.Equal(t, tt.expectedIndexes[j], wrappedPFB.ShareIndexes, j)
			}

			blockData := tmproto.Data{
				Txs:        wrappedPFBs,
				Blobs:      blobs,
				SquareSize: tt.squareSize,
			}

			coreData, err := coretypes.DataFromProto(&blockData)
			require.NoError(t, err)

			_, err = shares.Split(coreData, true)
			require.NoError(t, err)
		})
	}
}
