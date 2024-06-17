package statesync

import (
	"context"
	"encoding/json"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hack3r-0m/cosmoavs/x/statesync/keeper"
	types "github.com/hack3r-0m/cosmoavs/x/statesync/types"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.HasName        = AppModule{}
	_ module.AppModuleBasic = AppModule{}
	_ module.HasGenesis     = AppModule{}
	_ appmodule.AppModule   = AppModule{}
)

const ConsensusVersion = 1

// AppModule implements the AppModule interface for the epochs module.
type AppModule struct {
	cdc    codec.Codec
	keeper *keeper.Keeper
}

// IsOnePerModuleType implements appmodule.AppModule.
func (am AppModule) IsOnePerModuleType() {
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, keeper *keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

func NewAppModuleBasic(m AppModule) module.AppModuleBasic {
	return module.CoreAppModuleBasicAdaptor(m.Name(), m)
}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// Name returns the epochs module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInterfaces registers interfaces and implementations
func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic.
func (am AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterQueryServer(registrar, keeper.NewQueryServer(*am.keeper))
	return nil
}

func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genesis := keeper.DefaultGenesis()
	return cdc.MustMarshalJSON(&genesis)
}

// InitGenesis performs genesis initialization
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	am.keeper.InitGenesis(ctx, &genesisState)
}

// ValidateGenesis performs genesis state validation
func (AppModule) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gs := am.keeper.ExportGenesis(sdkCtx)

	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements HasConsensusVersion
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }
