package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) SignContract(goCtx context.Context, msg *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract, found := k.GetContracts(ctx, msg.Cid)
	if !found {
		return nil, fmt.Errorf("Contract not found")
	}

	if contract.Signee != msg.Creator {
		return nil, fmt.Errorf("You do not have permission to approve this contract.")
	}

	eblock, ok := sdk.NewIntFromString(contract.Duration)
	if !ok {
		return nil, fmt.Errorf("Duration failed to convert to int")
	}

	deal := types.ActiveDeals{
		Cid:           contract.Cid,
		Signee:        contract.Signee,
		Miner:         contract.Creator,
		Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:      fmt.Sprintf("%d", ctx.BlockHeight()+eblock.Int64()),
		Filesize:      contract.Filesize,
		Proofverified: "false",
		Blocktoprove:  "0",
		Creator:       msg.Creator,
		Proofsmissed:  "0",
		Merkle:        contract.Merkle,
	}

	k.SetActiveDeals(ctx, deal)
	k.RemoveContracts(ctx, contract.Cid)

	return &types.MsgSignContractResponse{}, nil
}
