package keeper

import (
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
)

func TestGenesis(t *testing.T) {
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

	if genState.L1ChainId != 1 {
		t.Errorf("DefaultGenesis L1ChainId should be 1, got %d", genState.L1ChainId)
	}

	genState.L1ChainId = 1
	genState.L1StartBlock = 0
	genState.RegistryCoordinator = "0x0baac79acd45a023e19345c352d8a7a83c4e5656"
	genState.OperatorStateRetriever = "0xD5D7fB4647cE79740E6e83819EFDf43fa74F8C31"

	k.InitGenesis(ctx, &genState)

	if getL1ChainID(ctx, &k) != 1 {
		t.Errorf("%d %d", getL1ChainID(ctx, &k), genState.L1ChainId)
	}

	if getL1StartBlock(ctx, &k) != 0 {
		t.Errorf("%d %d", getL1StartBlock(ctx, &k), genState.L1StartBlock)
	}

	if getRegistryCoordinator(ctx, &k) != genState.RegistryCoordinator {
		t.Errorf("%s %s", getRegistryCoordinator(ctx, &k), genState.RegistryCoordinator)
	}

	if getOperatorStateRetriever(ctx, &k) != genState.OperatorStateRetriever {
		t.Errorf("%s %s", getOperatorStateRetriever(ctx, &k), genState.OperatorStateRetriever)
	}

	exportedGenState := k.ExportGenesis(ctx)

	if exportedGenState.L1ChainId != genState.L1ChainId {
		t.Errorf("%d %d", exportedGenState.L1ChainId, genState.L1ChainId)
	}

	if exportedGenState.L1StartBlock != genState.L1StartBlock {
		t.Errorf("%d %d", exportedGenState.L1StartBlock, genState.L1StartBlock)
	}

	if exportedGenState.RegistryCoordinator != genState.RegistryCoordinator {
		t.Errorf("%s %s", exportedGenState.RegistryCoordinator, genState.RegistryCoordinator)
	}

	if exportedGenState.OperatorStateRetriever != genState.OperatorStateRetriever {
		t.Errorf("%s %s", exportedGenState.OperatorStateRetriever, genState.OperatorStateRetriever)
	}

}
