package keeper

import (
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/core/header"
	"cosmossdk.io/store/types"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/gogoproto/proto"
	localtypes "github.com/hack3r-0m/cosmoavs/x/statesync/types"
)

func TestQuery(t *testing.T) {
	key := types.NewKVStoreKey("test")

	storeService := runtime.NewKVStoreService(key)

	testCtx := testutil.DefaultContextWithDB(t, key, types.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now()})

	ir, _ := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles:     proto.HybridResolver,
		SigningOptions: signing.Options{},
	})

	k := NewKeeper(codec.NewProtoCodec(ir), storeService)
	genState := DefaultGenesis()

	genState.L1ChainId = 1
	genState.L1StartBlock = 0
	genState.RegistryCoordinator = "0x0baac79acd45a023e19345c352d8a7a83c4e5656"
	genState.OperatorStateRetriever = "0xD5D7fB4647cE79740E6e83819EFDf43fa74F8C31"

	k.InitGenesis(ctx, &genState)

	qs := NewQueryServer(k)

	query := localtypes.QueryOperatorStateRequest{
		Quorum: 1,
		Block:  20108817,
	}

	out, _ := qs.OperatorState(ctx, &query)
	fmt.Printf("%+v\n", out)
}
