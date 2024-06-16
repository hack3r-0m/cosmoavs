package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc codec.BinaryCodec

	Schema       collections.Schema
	storeService storetypes.KVStoreService

	avsClient *avsregistry.AvsRegistryChainReader
}

func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService) Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:          cdc,
		storeService: storeService,
	}

	schema, err := sb.Build()

	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}
