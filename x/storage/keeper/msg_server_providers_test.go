package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// testing msg server files for: init_provider, set_provider_ip, set_provider_totalspace

func (suite *KeeperTestSuite) TestMsgInitProvider() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgInitProvider
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgInitProvider {
				return types.NewMsgInitProvider(
					user.String(),
					"127.0.0.1",
					"1000000000",
					"test-key",
				)
			},
			expErr: false,
			name:   "init provider success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.InitProvider(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitProviderResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderIP() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider := types.Providers{
		Address:         user.String(),
		Ip:              "",
		Totalspace:      "1000000",
		Creator:         user.String(),
		BurnedContracts: "0",
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	cases := []struct {
		preRun    func() *types.MsgSetProviderIP
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderIP {
				return types.NewMsgSetProviderIP(
					user.String(),
					"127.0.0.1",
				)
			},
			expErr: false,
			name:   "set provider ip success",
		},
		{
			preRun: func() *types.MsgSetProviderIP {
				return types.NewMsgSetProviderIP(
					"wrong address",
					"127.0.0.1",
				)
			},
			expErr:    true,
			expErrMsg: "Provider not found. Please init your provider.",
			name:      "set provider ip fail",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderIP(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderIPResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderTotalSpace() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider := types.Providers{
		Address:         user.String(),
		Ip:              "127.0.0.1",
		Totalspace:      "",
		Creator:         user.String(),
		BurnedContracts: "0",
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	cases := []struct {
		preRun    func() *types.MsgSetProviderTotalspace
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderTotalspace {
				return types.NewMsgSetProviderTotalspace(
					user.String(),
					"1000000",
				)
			},
			expErr: false,
			name:   "set provider total space success",
		},
		{
			preRun: func() *types.MsgSetProviderTotalspace {
				return types.NewMsgSetProviderTotalspace(
					"wrong address",
					"1000000",
				)
			},
			expErr:    true,
			expErrMsg: "Provider not found. Please init your provider.",
			name:      "set provider total space fail",
		},
		{
			preRun: func() *types.MsgSetProviderTotalspace {
				return types.NewMsgSetProviderTotalspace(
					user.String(),
					"9@!0",
				)
			},
			expErr:    true,
			expErrMsg: "Not a valid total space. Please enter total number of bytes to provide.",
			name:      "invalid space param",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderTotalspace(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderIPResponse{}, *res)

			}
		})
	}
}
