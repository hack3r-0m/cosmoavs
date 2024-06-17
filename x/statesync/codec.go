package statesync

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	localtypes "github.com/hack3r-0m/cosmoavs/x/statesync/types"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&localtypes.QueryOperatorStateRequest{},
		&localtypes.QueryOperatorStateResponse{},
	)
}
