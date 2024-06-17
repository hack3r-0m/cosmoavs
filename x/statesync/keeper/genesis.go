package keeper

import (
	"log"
	"os"

	"cosmossdk.io/collections"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hack3r-0m/cosmoavs/x/statesync/types"
	genState "github.com/hack3r-0m/cosmoavs/x/statesync/types"
)

func DefaultGenesis() genState.GenesisState {
	return genState.GenesisState{
		L1ChainId:              1,
		L1StartBlock:           0,
		RegistryCoordinator:    "",
		OperatorStateRetriever: "",
	}
}

func setL1ChainID(ctx sdk.Context, k *Keeper, value uint64) {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixL1ChainID, types.KeyL1ChainID, collections.Uint64Value)

	err := item.Set(ctx, value)
	if err != nil {
		log.Fatalf("error %v", err)
	}

}

func getL1ChainID(ctx sdk.Context, k *Keeper) uint64 {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixL1ChainID, types.KeyL1ChainID, collections.Uint64Value)

	value, err := item.Get(ctx)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return value
}

func setL1StartBlock(ctx sdk.Context, k *Keeper, value uint64) {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixL1StartBlock, types.KeyL1StartBlock, collections.Uint64Value)

	err := item.Set(ctx, value)
	if err != nil {
		log.Fatalf("error %v", err)
	}
}

func getL1StartBlock(ctx sdk.Context, k *Keeper) uint64 {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixL1StartBlock, types.KeyL1StartBlock, collections.Uint64Value)

	value, err := item.Get(ctx)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return value
}

func setRegistryCoordinator(ctx sdk.Context, k *Keeper, value string) {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixRegistryCoordinator, types.KeyRegistryCoordinator, collections.StringValue)

	err := item.Set(ctx, value)
	if err != nil {
		log.Fatalf("error %v", err)
	}

}

func getRegistryCoordinator(ctx sdk.Context, k *Keeper) string {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixRegistryCoordinator, types.KeyRegistryCoordinator, collections.StringValue)

	value, err := item.Get(ctx)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return value
}

func setOperatorStateRetriever(ctx sdk.Context, k *Keeper, value string) {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixOperatorStateRetriever, types.KeyOperatorStateRetriever, collections.StringValue)

	err := item.Set(ctx, value)
	if err != nil {
		log.Fatalf("error %v", err)
	}
}

func getOperatorStateRetriever(ctx sdk.Context, k *Keeper) string {
	sb := collections.NewSchemaBuilder(k.storeService)
	item := collections.NewItem(sb, types.KeyPrefixOperatorStateRetriever, types.KeyOperatorStateRetriever, collections.StringValue)

	value, err := item.Get(ctx)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	return value
}

func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	rpcUrl := os.Getenv("RPC_URL")
	if rpcUrl == "" {
		log.Fatal("RPC_URL env not set")
	}

	client, err := eth.NewClient(rpcUrl)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	logger, err := logging.NewZapLogger(logging.LogLevel("development"))
	if err != nil {
		log.Fatalf("error %v", err)
	}

	reader, err := avsregistry.BuildAvsRegistryChainReader(
		common.HexToAddress(genState.RegistryCoordinator),
		common.HexToAddress(genState.OperatorStateRetriever),
		client,
		logger,
	)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	k.avsClient = reader

	setL1ChainID(ctx, k, genState.L1ChainId)
	setL1StartBlock(ctx, k, genState.L1StartBlock)
	setRegistryCoordinator(ctx, k, genState.RegistryCoordinator)
	setOperatorStateRetriever(ctx, k, genState.OperatorStateRetriever)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *genState.GenesisState {
	return &genState.GenesisState{
		L1ChainId:              getL1ChainID(ctx, &k),
		L1StartBlock:           getL1StartBlock(ctx, &k),
		RegistryCoordinator:    getRegistryCoordinator(ctx, &k),
		OperatorStateRetriever: getOperatorStateRetriever(ctx, &k),
	}
}
