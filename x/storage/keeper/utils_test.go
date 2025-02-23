package keeper_test

import (
	module "github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestGetPaidAmount() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name     string
		preRun   func() (string, int64)
		paidAmt  int64
		free     bool
		payBlock *types.PayBlocks
	}{
		{
			name: "no_payblock",
			preRun: func() (string, int64) {
				suite.ctx = suite.ctx.WithBlockHeight(100)
				return "cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl", 1
			},
			paidAmt:  module.TwoGigs,
			free:     true,
			payBlock: nil,
		},

		{
			name: "Invalid_endblock_blocknum",
			preRun: func() (string, int64) {
				pbs := types.PayBlocks{
					Blockid:   "123",
					Bytes:     "100000000000",
					Blocktype: module.StartBlockType,
					Blocknum:  "aabbcc",
				}
				pbe := types.PayBlocks{
					// endblock blockid starts with '.'
					Blockid:   ".123",
					Bytes:     "100000000000",
					Blocktype: module.EndBlockType,
					Blocknum:  "aabbcc",
				}
				sKeeper.SetPayBlocks(suite.ctx, pbs)
				sKeeper.SetPayBlocks(suite.ctx, pbe)
				return "123", 0
			},
			paidAmt:  module.TwoGigs,
			free:     true,
			payBlock: nil,
		},

		{
			name: "invalid_eblock_bytes",
			preRun: func() (string, int64) {
				eBlock, found := sKeeper.GetPayBlocks(suite.ctx, ".123")
				suite.Require().True(found)
				eBlock.Blocknum = "10"
				eBlock.Bytes = "aabbcc"
				sKeeper.SetPayBlocks(suite.ctx, eBlock)
				return "123", 100
			},
			paidAmt: module.TwoGigs,
			free:    true,
			payBlock: &types.PayBlocks{
				Blockid:   ".123",
				Bytes:     "aabbcc",
				Blocktype: module.EndBlockType,
				Blocknum:  "10",
			},
		},

		{
			name: "one_month_grace_period",
			preRun: func() (string, int64) {
				eBlock, found := sKeeper.GetPayBlocks(suite.ctx, ".123")
				suite.Require().True(found)
				eBlock.Bytes = "10000000000"
				sKeeper.SetPayBlocks(suite.ctx, eBlock)
				return "123", 100
			},
			paidAmt:  10000000000,
			free:     true,
			payBlock: nil,
		},

		{
			name: "highest_block_zero",
			preRun: func() (string, int64) {
				suite.ctx = suite.ctx.WithBlockHeight(0)
				return "123", 1
			},
			paidAmt: module.TwoGigs,
			free:    true,
			payBlock: &types.PayBlocks{
				Blockid:   ".123",
				Bytes:     "10000000000",
				Blocktype: module.EndBlockType,
				Blocknum:  "10",
			},
		},

		{
			name: "highestblock_has_invalid_bytes",
			preRun: func() (string, int64) {
				suite.ctx = suite.ctx.WithBlockHeight(1000)
				pbs := types.PayBlocks{
					Blockid:   "123100",
					Bytes:     "aabbcc",
					Blocktype: module.StartBlockType,
					Blocknum:  "100",
				}
				pbe := types.PayBlocks{
					Blockid:   ".123",
					Bytes:     "100000000000",
					Blocktype: module.EndBlockType,
					Blocknum:  "200000",
				}
				sKeeper.SetPayBlocks(suite.ctx, pbs)
				sKeeper.SetPayBlocks(suite.ctx, pbe)
				return "123", 10000
			},
			paidAmt: module.TwoGigs,
			free:    true,
			payBlock: &types.PayBlocks{
				Blockid:   ".123",
				Bytes:     "100000000000",
				Blocktype: module.EndBlockType,
				Blocknum:  "200000",
			},
		},

		{
			name: "successful_paid_amount_returned",
			preRun: func() (string, int64) {
				pbs, found := sKeeper.GetPayBlocks(suite.ctx, "123100")
				suite.Require().True(found)
				pbs.Bytes = "100000000000"
				sKeeper.SetPayBlocks(suite.ctx, pbs)
				return "123", 10000
			},
			paidAmt: 100000000000,
			free:    false,
			payBlock: &types.PayBlocks{
				Blockid:   ".123",
				Bytes:     "100000000000",
				Blocktype: module.EndBlockType,
				Blocknum:  "200000",
			},
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			addr, blockh := tc.preRun()
			rPaidAmt, rFree, rPayBlock := sKeeper.GetPaidAmount(suite.ctx, addr, blockh)

			suite.Require().Equal(tc.paidAmt, rPaidAmt)
			suite.Require().Equal(tc.free, rFree)
			suite.Require().Equal(tc.payBlock, rPayBlock)
		})
	}
}

func (suite *KeeperTestSuite) TestCreatePayBlock() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name      string
		preRun    func() (string, int64, int64)
		check     func()
		expErr    bool
		expErrMsg string
	}{
		{
			name: "normal_payblock",
			preRun: func() (string, int64, int64) {
				suite.ctx = suite.ctx.WithBlockHeight(1)
				return "a", 10000, 10000
			},
			check: func() {
				pbs, found := sKeeper.GetPayBlocks(suite.ctx, "a1")
				suite.Require().True(found)
				suite.Require().Equal("a1", pbs.Blockid)
				suite.Require().Equal("10000", pbs.Bytes)
				suite.Require().Equal(module.StartBlockType, pbs.Blocktype)
				suite.Require().Equal("1", pbs.Blocknum)
				pbe, found := sKeeper.GetPayBlocks(suite.ctx, ".a")
				suite.Require().True(found)
				suite.Require().Equal(".a", pbe.Blockid)
				suite.Require().Equal("10000", pbe.Bytes)
				suite.Require().Equal(module.EndBlockType, pbe.Blocktype)
				suite.Require().Equal("10001", pbe.Blocknum)
			},
			expErr: false,
		},

		{
			name: "buying_within_existing_storage_window_payblock",
			preRun: func() (string, int64, int64) {
				return "a", 10000, 10000
			},
			check:     func() {},
			expErr:    true,
			expErrMsg: "can't buy storage within another storage window",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			suite.Require().NotNil(tc.preRun)
			addr, length, bytes := tc.preRun()
			err := sKeeper.CreatePayBlock(suite.ctx, addr, length, bytes)
			tc.check()

			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetProviderUsing() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name      string
		preRun    func() string
		expReturn int64
	}{
		{
			name: "No_provider_found",
			preRun: func() string {
				return "a"
			},
			expReturn: 0,
		},

		{
			name: "invalid_active_deal_file_size",
			preRun: func() string {
				ad := types.ActiveDeals{
					Provider: "a",
					Filesize: "aaaaa",
					Cid:      "abc",
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return "a"
			},
			expReturn: 0,
		},

		{
			name: "valid_active_deal_file_size",
			preRun: func() string {
				ad := types.ActiveDeals{
					Provider: "a",
					Filesize: "100000",
					Cid:      "bbb",
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return "a"
			},
			expReturn: 100000,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			suite.Require().NotNil(tc.preRun)
			provider := tc.preRun()
			result := sKeeper.GetProviderUsing(suite.ctx, provider)

			suite.Require().Equal(tc.expReturn, result)
		})
	}
}
