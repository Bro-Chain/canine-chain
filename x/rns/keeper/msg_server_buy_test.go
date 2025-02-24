package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestBuyMsg() {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)
	keeper := suite.rnsKeeper

	// Create name owner account
	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)
	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000))) // Send some coins to their account
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)
	// Create buyer account
	buyer, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)
	coins = sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(1000))) // Send insufficient amount amout
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	// Init rns account and register rns
	rnsName := "Nuggie"
	rnsTLD := "jkl"
	fullName := rnsName + "." + rnsTLD // "Nuggie.jkl"

	keeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), fullName, "{}", "2")
	suite.Require().NoError(err)
	originalNames, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
	suite.Require().True(found)

	// Put it up for sale
	salePrice := "1000000"
	msgList := types.MsgList{
		Creator: nameOwner.String(),
		Name:    fullName,
		Price:   salePrice,
	}
	_, err = msgSrvr.List(ctx, &msgList)
	suite.Require().NoError(err)

	// Test cases
	cases := []struct {
		testName  string
		preRun    func() *types.MsgBuy
		postRun   func()
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "name_not_for_sale",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Name:    "not_for_sale.jkl",
					Creator: buyer.String(),
				}
			},
			expErr:    true,
			expErrMsg: "Name not for sale.",
		},

		{
			testName: "name_not_found",
			preRun: func() *types.MsgBuy {
				// Delete name from KVStore
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				if !found {
					return nil
				}
				keeper.RemoveNames(suite.ctx, names.Name, names.Tld)
				_, found = keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().False(found)
				return &types.MsgBuy{
					Name:    fullName,
					Creator: buyer.String(),
				}
			},
			postRun: func() {
				// add back name
				keeper.SetNames(suite.ctx, originalNames)
				_, found = keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
			},
			expErr:    true,
			expErrMsg: "Name does not exist",
		},

		{
			testName: "name_expired",
			preRun: func() *types.MsgBuy {
				// Make the name expired
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
				names.Expires = suite.ctx.BlockHeight() - 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgBuy{
					Name:    fullName,
					Creator: buyer.String(),
				}
			},
			postRun: func() {
				names, found := keeper.GetNames(suite.ctx, rnsName, rnsTLD)
				suite.Require().True(found)
				names.Expires++
				keeper.SetNames(suite.ctx, names)
			},
			expErr:    true,
			expErrMsg: "expired",
		},

		{
			testName: "owner_cannot_buy_owners_name",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Creator: nameOwner.String(),
					Name:    fullName,
				}
			},
			expErr:    true,
			expErrMsg: "cannot buy your own name",
		},

		{
			testName: "not_enough_balance",
			preRun: func() *types.MsgBuy {
				return &types.MsgBuy{
					Creator: buyer.String(),
					Name:    fullName,
				}
			},
			expErr:    true,
			expErrMsg: "not enough balance",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			msg := tc.preRun()

			err := keeper.BuyName(suite.ctx, msg.Creator, msg.Name)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}

			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}
