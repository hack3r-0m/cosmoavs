package keeper

import (
	"context"
	"log"

	avstypes "github.com/Layr-Labs/eigensdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hack3r-0m/cosmoavs/x/statesync/types"
)

var _ types.QueryServer = QueryServer{}

type QueryServer struct {
	Keeper
}

func NewQueryServer(keeper Keeper) types.QueryServer {
	return &QueryServer{Keeper: keeper}
}

func (s QueryServer) OperatorState(ctx context.Context, req *types.QueryOperatorStateRequest) (*types.QueryOperatorStateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if req.Quorum > 255 {
		log.Fatalf("error: quorum overflow")
	}

	quorumNums := make([]avstypes.QuorumNum, 1)
	quorumNums[0] = avstypes.QuorumNum(uint8(req.Quorum))

	operatorSet, err := s.Keeper.avsClient.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{
		Pending: false,
		Context: sdkCtx,
	}, quorumNums, uint32(req.Block))

	if err != nil {
		return nil, err
	}

	var opSet []types.OperatorSet

	for index, operatorState := range operatorSet {
		opSet = append(opSet, types.OperatorSet{
			Operator:   operatorState[index].Operator.Hex(),
			OperatorId: operatorState[index].OperatorId[:],
			Stake:      operatorState[index].Stake.Bytes(),
		})
	}

	return &types.QueryOperatorStateResponse{
		Operators: opSet,
	}, nil
}
