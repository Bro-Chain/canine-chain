package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPayData(goCtx context.Context, req *types.QueryPayDataRequest) (*types.QueryPayDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	paid, _, block := k.GetPaidAmount(ctx, req.Address, ctx.BlockHeight())
	blocks := ctx.BlockHeight()
	if block != nil {
		bnum, ok := sdk.NewIntFromString(block.Blocknum)
		if ok {
			blocks = bnum.Int64()
		}
	}

	blocks -= ctx.BlockHeight()

	return &types.QueryPayDataResponse{BlocksRemaining: blocks, Bytes: paid}, nil
}
