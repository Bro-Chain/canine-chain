package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) SetProviderKeybase(goCtx context.Context, msg *types.MsgSetProviderKeybase) (*types.MsgSetProviderKeybaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	provider.KeybaseIdentity = msg.Keybase

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderKeybaseResponse{}, nil
}
