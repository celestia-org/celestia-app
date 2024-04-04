package module_test

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"

	"github.com/celestiaorg/celestia-app/v2/app/module"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/tests/mocks"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestManagerOrderSetters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)
	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)

	mockAppModule1.EXPECT().Name().Times(6).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(6).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	require.Equal(t, []string{"module1", "module2"}, mm.OrderInitGenesis)
	mm.SetOrderInitGenesis("module2", "module1")
	require.Equal(t, []string{"module2", "module1"}, mm.OrderInitGenesis)

	require.Equal(t, []string{"module1", "module2"}, mm.OrderExportGenesis)
	mm.SetOrderExportGenesis("module2", "module1")
	require.Equal(t, []string{"module2", "module1"}, mm.OrderExportGenesis)

	require.Equal(t, []string{"module1", "module2"}, mm.OrderBeginBlockers)
	mm.SetOrderBeginBlockers("module2", "module1")
	require.Equal(t, []string{"module2", "module1"}, mm.OrderBeginBlockers)

	require.Equal(t, []string{"module1", "module2"}, mm.OrderEndBlockers)
	mm.SetOrderEndBlockers("module2", "module1")
	require.Equal(t, []string{"module2", "module1"}, mm.OrderEndBlockers)
}

func TestManager_RegisterInvariants(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(2).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	// test RegisterInvariants
	mockInvariantRegistry := mocks.NewMockInvariantRegistry(mockCtrl)
	mockAppModule1.EXPECT().RegisterInvariants(gomock.Eq(mockInvariantRegistry)).Times(1)
	mockAppModule2.EXPECT().RegisterInvariants(gomock.Eq(mockInvariantRegistry)).Times(1)
	mm.RegisterInvariants(mockInvariantRegistry)
}

func TestManager_RegisterQueryServices(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(3).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(2).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(3).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(2).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	msgRouter := mocks.NewMockServer(mockCtrl)
	queryRouter := mocks.NewMockServer(mockCtrl)
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	cfg := module.NewConfigurator(cdc, msgRouter, queryRouter)
	mockAppModule1.EXPECT().RegisterServices(gomock.Any()).Times(1)
	mockAppModule2.EXPECT().RegisterServices(gomock.Any()).Times(1)

	mm.RegisterServices(cfg)
}

func TestManager_InitGenesis(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(2).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	ctx := sdk.NewContext(nil, tmproto.Header{}, false, log.NewNopLogger())
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	genesisData := map[string]json.RawMessage{"module1": json.RawMessage(`{"key": "value"}`)}

	// this should panic since the validator set is empty even after init genesis
	mockAppModule1.EXPECT().InitGenesis(gomock.Eq(ctx), gomock.Eq(cdc), gomock.Eq(genesisData["module1"])).Times(1).Return(nil)
	require.Panics(t, func() { mm.InitGenesis(ctx, cdc, genesisData, 1) })

	// test panic
	genesisData = map[string]json.RawMessage{
		"module1": json.RawMessage(`{"key": "value"}`),
		"module2": json.RawMessage(`{"key": "value"}`),
	}
	mockAppModule1.EXPECT().InitGenesis(gomock.Eq(ctx), gomock.Eq(cdc), gomock.Eq(genesisData["module1"])).Times(1).Return([]abci.ValidatorUpdate{{}})
	mockAppModule2.EXPECT().InitGenesis(gomock.Eq(ctx), gomock.Eq(cdc), gomock.Eq(genesisData["module2"])).Times(1).Return([]abci.ValidatorUpdate{{}})
	require.Panics(t, func() { mm.InitGenesis(ctx, cdc, genesisData, 1) })
}

func TestManager_ExportGenesis(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(2).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	ctx := sdk.Context{}
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	mockAppModule1.EXPECT().ExportGenesis(gomock.Eq(ctx), gomock.Eq(cdc)).Times(1).Return(json.RawMessage(`{"key1": "value1"}`))
	mockAppModule2.EXPECT().ExportGenesis(gomock.Eq(ctx), gomock.Eq(cdc)).Times(1).Return(json.RawMessage(`{"key2": "value2"}`))

	want := map[string]json.RawMessage{
		"module1": json.RawMessage(`{"key1": "value1"}`),
		"module2": json.RawMessage(`{"key2": "value2"}`),
	}
	require.Equal(t, want, mm.ExportGenesis(ctx, cdc, 1))
}

func TestManager_BeginBlock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(2).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	req := abci.RequestBeginBlock{Hash: []byte("test")}

	mockAppModule1.EXPECT().BeginBlock(gomock.Any(), gomock.Eq(req)).Times(1)
	mockAppModule2.EXPECT().BeginBlock(gomock.Any(), gomock.Eq(req)).Times(1)
	ctx := sdk.NewContext(nil, tmproto.Header{
		Version: tmversion.Consensus{App: 1},
	}, false, log.NewNopLogger())
	mm.BeginBlock(ctx, req)
}

func TestManager_EndBlock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("module1")
	mockAppModule1.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mockAppModule2.EXPECT().Name().Times(2).Return("module2")
	mockAppModule2.EXPECT().ConsensusVersion().Times(1).Return(uint64(1))
	mm, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 1, ToVersion: 1},
	})
	require.NoError(t, err)
	require.NotNil(t, mm)
	require.Equal(t, 2, len(mm.ModuleNames(1)))

	req := abci.RequestEndBlock{Height: 10}

	mockAppModule1.EXPECT().EndBlock(gomock.Any(), gomock.Eq(req)).Times(1).Return([]abci.ValidatorUpdate{{}})
	mockAppModule2.EXPECT().EndBlock(gomock.Any(), gomock.Eq(req)).Times(1)
	ctx := sdk.NewContext(nil, tmproto.Header{
		Version: tmversion.Consensus{App: 1},
	}, false, log.NewNopLogger())
	ret := mm.EndBlock(ctx, req)
	require.Equal(t, []abci.ValidatorUpdate{{}}, ret.ValidatorUpdates)

	// test panic
	mockAppModule1.EXPECT().EndBlock(gomock.Any(), gomock.Eq(req)).Times(1).Return([]abci.ValidatorUpdate{{}})
	mockAppModule2.EXPECT().EndBlock(gomock.Any(), gomock.Eq(req)).Times(1).Return([]abci.ValidatorUpdate{{}})
	require.Panics(t, func() { mm.EndBlock(ctx, req) })
}

func TestManager_UpgradeSchedule(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)

	mockAppModule1 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule2 := mocks.NewMockAppModule(mockCtrl)
	mockAppModule1.EXPECT().Name().Times(2).Return("blob")
	mockAppModule2.EXPECT().Name().Times(2).Return("blob")
	mockAppModule1.EXPECT().ConsensusVersion().Times(2).Return(uint64(3))
	mockAppModule2.EXPECT().ConsensusVersion().Times(2).Return(uint64(2))
	_, err := module.NewManager([]module.VersionedModule{
		{Module: mockAppModule1, FromVersion: 1, ToVersion: 1},
		{Module: mockAppModule2, FromVersion: 2, ToVersion: 2},
	})
	require.Error(t, err)
}
