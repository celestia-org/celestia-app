package testnode

import (
	"context"
	"net"
	"testing"

	"github.com/celestiaorg/celestia-app/test/util/genesis"
	"github.com/stretchr/testify/require"
)

// NewNetwork starts a single validator celestia-app network using the provided
// configurations. Configured accounts will be funded and their keys can be
// accessed in keyring returned client.Context. All rpc, p2p, and grpc addresses
// in the provided configs are overwritten to use open ports. The node can be
// accessed via the returned client.Context or via the returned rpc and grpc
// addresses. Configured genesis options will be applied after all accounts have
// been initialized.
func NewNetwork(t testing.TB, cfg *Config) (cctx Context, rpcAddr, grpcAddr string) {
	t.Helper()

	// initialize the genesis file and validator files for the first validator.
	baseDir, err := genesis.InitFiles(t.TempDir(), cfg.TmConfig, cfg.Genesis, 0)
	require.NoError(t, err)

	tmNode, app, err := NewCometNode(baseDir, &cfg.UniversalTestingConfig)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})

	cctx = NewContext(ctx, cfg.Genesis.Keyring(), cfg.TmConfig, cfg.Genesis.ChainID, cfg.AppConfig.API.Address)

	cctx, stopNode, err := StartNode(tmNode, cctx)
	require.NoError(t, err)

	cctx, cleanupGRPC, err := StartGRPCServer(app, cfg.AppConfig, cctx)
	require.NoError(t, err)

	apiServer, err := StartAPIServer(app, *cfg.AppConfig, cctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		t.Log("tearing down testnode")
		err := stopNode()
		if err != nil {
			// the test has already completed so log the error instead of
			// failing the test.
			t.Logf("error stopping node %v", err)
		}
		err = cleanupGRPC()
		if err != nil {
			// the test has already completed so just log the error instead of
			// failing the test.
			t.Logf("error when cleaning up GRPC %v", err)
		}
		err = apiServer.Close()
		if err != nil {
			// the test has already completed so just log the error instead of
			// failing the test.
			t.Logf("error when closing API server %v", err)
		}
	})

	return cctx, cfg.TmConfig.RPC.ListenAddress, cfg.AppConfig.GRPC.Address
}

// getFreePort returns a free port and optionally an error.
func getFreePort() (int, error) {
	a, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// mustGetFreePort returns a free port. Panics if no free ports are available or
// an error is encountered.
func mustGetFreePort() int {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}
	return port
}
