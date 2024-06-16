package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name.
	ModuleName = "cosmoavs/statesync"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName
)

const (
	KeyL1ChainID              = "L1ChainID"
	KeyL1StartBlock           = "L1StartBlock"
	KeyRegistryCoordinator    = "RegistryCoordinator"
	KeyOperatorStateRetriever = "OperatorStateRetriever"
)

var (
	KeyPrefixL1ChainID              = collections.NewPrefix(ModuleName + KeyL1ChainID)
	KeyPrefixL1StartBlock           = collections.NewPrefix(ModuleName + KeyL1StartBlock)
	KeyPrefixRegistryCoordinator    = collections.NewPrefix(ModuleName + KeyRegistryCoordinator)
	KeyPrefixOperatorStateRetriever = collections.NewPrefix(ModuleName + KeyOperatorStateRetriever)
)
