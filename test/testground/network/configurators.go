package network

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/tendermint/tendermint/pkg/trace/schema"
	"github.com/testground/sdk-go/runtime"
)

const (
	TopologyParam         = "topology"
	ConnectAllTopology    = "connect_all"
	ConnectRandomTopology = "connect_random"
	SeedTopology          = "seed"
	SeedGroupID           = "seeds"
)

func DefaultTopologies() []string {
	return []string{
		ConnectAllTopology,
	}
}

// GetConfigurators
func GetConfigurators(runenv *runtime.RunEnv) ([]Configurator, error) {
	topology := runenv.StringParam(TopologyParam)
	if topology == "" {
		topology = ConnectAllTopology
	}
	ops := make([]Configurator, 0)
	switch topology {
	case ConnectAllTopology:
		ops = append(ops, ConnectAll)
	case ConnectRandomTopology:
		ops = append(ops, ConnectRandom(10))
	case SeedTopology:
		// don't do anything since we are manually adding peers to the address book
	default:
		return nil, fmt.Errorf("unknown topology func: %s", topology)
	}

	ops = append(ops, TracingConfigurator(runenv, ParseTracingParams(runenv)))

	return ops, nil
}

// Configurator is a function that arbitrarily modifies the provided node
// configurations. It is used to generate the topology (which nodes are
// connected to which) of the network, along with making other arbitrary changes
// to the configs.
type Configurator func(nodes []RoleConfig) ([]RoleConfig, error)

var _ = Configurator(ConnectAll)

// ConnectAll is a Configurator that connects all nodes to each other via
// persistent peers.
func ConnectAll(nodes []RoleConfig) ([]RoleConfig, error) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GlobalSequence < nodes[j].GlobalSequence
	})
	peerIDs := peerIDs(nodes)

	// For each node, generate the string that excludes its own P2PID
	for i, nodeConfig := range nodes {
		var filteredP2PIDs []string
		for _, pid := range peerIDs {
			if pid != nodeConfig.PeerID {
				filteredP2PIDs = append(filteredP2PIDs, pid)
			}
		}

		// Here you could put the concatenated string into another field in NodeConfig
		// or do whatever you want with it.
		nodeConfig.CmtConfig.P2P.PersistentPeers = strings.Join(filteredP2PIDs, ",")
		nodes[i] = nodeConfig
	}

	return nodes, nil
}

func ConnectRandom(numPeers int) Configurator {
	return func(nodes []RoleConfig) ([]RoleConfig, error) {
		if numPeers >= len(nodes) {
			return nil, errors.New("numPeers should be less than the total number of nodes")
		}

		for i, nodeConfig := range nodes {
			// Shuffle the indexes for each nodeConfig
			indexes := rand.Perm(len(nodes))

			var chosenPeers []string

			for _, idx := range indexes {
				potentialPeer := nodes[idx]

				if len(chosenPeers) >= numPeers {
					break
				}
				if potentialPeer.PeerID != nodeConfig.PeerID {
					chosenPeers = append(chosenPeers, potentialPeer.PeerID)
				}
			}

			nodeConfig.CmtConfig.P2P.PersistentPeers = strings.Join(chosenPeers, ",")
			nodes[i] = nodeConfig
		}

		return nodes, nil
	}
}

// TracingConfigurator is a Configurator that configures tracing for the
// network. It will set the nodes to collect only the round state data, and will
// set the nodes specified in the TracingParams to collect all trace data.
func TracingConfigurator(runenv *runtime.RunEnv, tparams TracingParams) Configurator {
	return func(nodes []RoleConfig) ([]RoleConfig, error) {
		runenv.RecordMessage(fmt.Sprintf("tracing nodes: %+v", tparams))

		// set all of the nodes to collect the round state data. This allows us
		// to measure when exactly each node progresses to the next step of
		// consensus, but we are not overloading the influxdb instance with too
		// much trace data.
		for i := range nodes {
			nodes[i].CmtConfig.Instrumentation.InfluxOrg = "celestia"
			nodes[i].CmtConfig.Instrumentation.InfluxBucket = "testground"
			nodes[i].CmtConfig.Instrumentation.InfluxBatchSize = 200
			nodes[i].CmtConfig.Instrumentation.InfluxURL = tparams.URL
			nodes[i].CmtConfig.Instrumentation.InfluxToken = tparams.Token
			nodes[i].CmtConfig.Instrumentation.InfluxTables = []string{schema.RoundStateTable}
		}

		// Trace all data from these nodes. We might want to make this more
		// configurable in the future.
		for i := 0; i < tparams.Nodes; i++ {
			nodes[i].CmtConfig.Instrumentation.InfluxTables = schema.AllTables()
		}

		return nodes, nil
	}
}
