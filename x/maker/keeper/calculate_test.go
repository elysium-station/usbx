package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	blackfury "github.com/elysium-station/blackfury/types"
	"github.com/elysium-station/blackfury/x/maker/types"
)

func (suite *KeeperTestSuite) TestEstimateMintBySwapIn() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateMintBySwapInRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateMintBySwapInResponse
	}{
		{
			name: "black price too low",
			malleate: func() {
				suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(1009, 3))
			},
			req:     &types.EstimateMintBySwapInRequest{BackingDenom: suite.bcDenom},
			expPass: false,
			expErr:  types.ErrBlackPriceTooLow,
		},
		{
			name:    "backing denom not found",
			req:     &types.EstimateMintBySwapInRequest{BackingDenom: "fil"},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateMintBySwapInRequest{BackingDenom: "eth"},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "black over ceiling",
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(2_500000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrBlackCeiling,
		},
		{
			name: "default full backing",
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapInResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1015152)), // 1_000000 * (1+0.005) / 0.99
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5000)),
			},
		},
		{
			name: "user asked full backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
				FullBacking:  true,
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapInResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1015152)), // 1_000000 * (1+0.005) / 0.99
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5000)),
			},
		},
		{
			name: "full algorithmic",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.ZeroDec())
			},
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapInResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10050_000000_000000)), // 1_000000 * (1+0.005) / 10**-10
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5000)),
			},
		},
		{
			name: "fractional",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapInResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(812121)),                     // 1_000000 * (1+0.005) * 0.8 / 0.99
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(2010_000000_000000)), // 1_000000 * (1+0.005) * 0.2 / 10**-10
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5000)),
			},
		},
		{
			name: "backing over ceiling",
			req: &types.EstimateMintBySwapInRequest{
				MintOut:      sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_500000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrBackingCeiling,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(101, 2))
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateMintBySwapIn(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateMintBySwapOut() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateMintBySwapOutRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateMintBySwapOutResponse
	}{
		{
			name: "black price too low",
			malleate: func() {
				suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(989, 3))
			},
			req:     &types.EstimateMintBySwapOutRequest{BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBlackPriceTooLow,
		},
		{
			name:    "backing denom not found",
			req:     &types.EstimateMintBySwapOutRequest{BackingInMax: sdk.NewCoin("fil", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateMintBySwapOutRequest{BackingInMax: sdk.NewCoin("eth", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "default full backing",
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(985050)), // 1_000000 * 0.99 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(4950)),   // 1_000000 * 0.99 * 0.005
			},
		},
		{
			name: "user asked full backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FullBacking:  true,
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(985050)), // 1_000000 * 0.99 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(4950)),   // 1_000000 * 0.99 * 0.005
			},
		},
		{
			name: "full algorithmic",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.ZeroDec())
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryInMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(995000)), // 10**16 * 10**-10 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5000)),   // 10**16 * 10**-10 * 0.005
			},
		},
		{
			name: "zero fury using backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryInMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(1237_500000_000000)), // 500000 * 0.99 / 0.8 * 0.2 / (10**-10)
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(615656)),             // 500000 * 0.99 / 0.8 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(3094)),               // 500000 * 0.99 / 0.8 * 0.005
			},
		},
		{
			name: "fractional using max backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryInMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(1237_500000_000000)), // 500000 * 0.99 / 0.8 * 0.2 / (10**-10)
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(615656)),             // 500000 * 0.99 / 0.8 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(3094)),               // 500000 * 0.99 / 0.8 * 0.005
			},
		},
		{
			name: "zero backing using fury",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(20, 2))
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),
				FuryInMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(252525)), // 10**16 * 10**-10 / 0.8 * 0.2 / 0.99
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1243750)), // 10**16 * 10**-10 / 0.8 * (1 - 0.005))
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6250)),    // 10**16 * 10**-10 / 0.8 * 0.005
			},
		},
		{
			name: "fractional using max fury",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(20, 2))
			},
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryInMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateMintBySwapOutResponse{
				BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(252525)), // 10**16 * 10**-10 / 0.8 * 0.2 / 0.99
				FuryIn:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				MintOut:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1243750)), // 10**16 * 10**-10 / 0.8 * (1 - 0.005)
				MintFee:   sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6250)),    // 10**16 * 10**-10 / 0.8 * 0.005
			},
		},
		{
			name: "black over ceiling",
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(2_500000)),
			},
			expPass: false,
			expErr:  types.ErrBlackCeiling,
		},
		{
			name: "backing over ceiling",
			req: &types.EstimateMintBySwapOutRequest{
				BackingInMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_500000)),
			},
			expPass: false,
			expErr:  types.ErrBackingCeiling,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(101, 2))
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateMintBySwapOut(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateBurnBySwapIn() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateBurnBySwapInRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateBurnBySwapInResponse
	}{
		{
			name: "black price too high",
			malleate: func() {
				suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(991, 3))
			},
			req:     &types.EstimateBurnBySwapInRequest{BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBlackPriceTooHigh,
		},
		{
			name:    "backing denom not found",
			req:     &types.EstimateBurnBySwapInRequest{BackingOutMax: sdk.NewCoin("fil", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateBurnBySwapInRequest{BackingOutMax: sdk.NewCoin("eth", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "moudle backing insufficient",
			malleate: func() {
				suite.app.BankKeeper.BurnCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(suite.bcDenom, sdk.NewInt(1000_000000))))
			},
			req:     &types.EstimateBurnBySwapInRequest{BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000))},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "full backing",
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(995976)), // 1_000000 * 0.99 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(5976)), // 1_000000 * 0.99 / (1-0.006) * 0.006
			},
		},
		{
			name: "full algorithmic",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.ZeroDec())
			},
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryOutMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1006036)), // 10**16 * 10**-10 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6036)), // 10**16 * 10**-10 / (1-0.006) * 0.006
			},
		},
		{
			name: "zero fury using backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryOutMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(622485)), // 500000 * 0.99 / 0.8 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(1237_500000_000000)), // 500000 * 0.99 / 0.8 * 0.2 / (10**-10)
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(3735)),               // 500000 * 0.99 / 0.8 / (1-0.006) * 0.006
			},
		},
		{
			name: "fractional using max backing",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryOutMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(622485)), // 500000 * 0.99 / 0.8 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(1237_500000_000000)), // 500000 * 0.99 / 0.8 * 0.2 / (10**-10)
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(3735)),               // 500000 * 0.99 / 0.8 / (1-0.006) * 0.006
			},
		},
		{
			name: "zero backing using fury",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(20, 2))
			},
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),
				FuryOutMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1257545)), // 10**16 * 10**-10 / 0.8 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(252525)),          // 10**16 * 10**-10 / 0.8 * 0.2 / 0.99
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(7545)), // 10**16 * 10**-10 / 0.8 / (1-0.006) * 0.006
			},
		},
		{
			name: "fractional using max fury",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(20, 2))
			},
			req: &types.EstimateBurnBySwapInRequest{
				BackingOutMax: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000)),
				FuryOutMax:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapInResponse{
				BurnIn:     sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1257545)), // 10**16 * 10**-10 / 0.8 / (1-0.006)
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(252525)),          // 10**16 * 10**-10 / 0.8 * 0.2 / 0.99
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(7545)), // 10**16 * 10**-10 / 0.8 / (1-0.006) * 0.006
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(99, 2))
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateBurnBySwapIn(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateBurnBySwapOut() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateBurnBySwapOutRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateBurnBySwapOutResponse
	}{
		{
			name: "black price too high",
			malleate: func() {
				suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(1011, 3))
			},
			req:     &types.EstimateBurnBySwapOutRequest{BackingDenom: suite.bcDenom},
			expPass: false,
			expErr:  types.ErrBlackPriceTooHigh,
		},
		{
			name:    "backing denom not found",
			req:     &types.EstimateBurnBySwapOutRequest{BackingDenom: "fil"},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateBurnBySwapOutRequest{BackingDenom: "eth"},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "pool backing insufficient",
			malleate: func() {
				poolBacking, found := suite.app.MakerKeeper.GetPoolBacking(suite.ctx, suite.bcDenom)
				suite.Require().True(found)
				poolBacking.Backing.Amount = sdk.ZeroInt()
				suite.app.MakerKeeper.SetPoolBacking(suite.ctx, poolBacking)
			},
			req: &types.EstimateBurnBySwapOutRequest{
				BurnIn:       sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "moudle backing insufficient",
			malleate: func() {
				suite.app.BankKeeper.BurnCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(suite.bcDenom, sdk.NewInt(1000_000000))))
			},
			req: &types.EstimateBurnBySwapOutRequest{
				BurnIn:       sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "full backing",
			req: &types.EstimateBurnBySwapOutRequest{
				BurnIn:       sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapOutResponse{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_004040)), // 1_000000 * (1-0.006) / 0.99
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6000)), // 1_000000 * 0.006
			},
		},
		{
			name: "full algorithmic",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.ZeroDec())
			},
			req: &types.EstimateBurnBySwapOutRequest{
				BurnIn:       sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapOutResponse{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.ZeroInt()),                          // 1_000000 * (1-0.006) / 0.99
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(9940_000000_000000)), // 1_000000 * (1-0.006) / 10**-10
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6000)),               // 1_000000 * 0.006
			},
		},
		{
			name: "fractional",
			malleate: func() {
				suite.app.MakerKeeper.SetBackingRatio(suite.ctx, sdk.NewDecWithPrec(80, 2))
			},
			req: &types.EstimateBurnBySwapOutRequest{
				BurnIn:       sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(1_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateBurnBySwapOutResponse{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(803232)),                     // 1_000000 * (1-0.006) * 0.8 / 0.99
				FuryOut:    sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(19880_00000_000000)), // 1_000000 * (1-0.006) * 0.2 / 10**-10
				BurnFee:    sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6000)),               // 1_000000 * 0.006
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(99, 2))
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateBurnBySwapOut(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateBuyBackingIn() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateBuyBackingInRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateBuyBackingInResponse
	}{
		{
			name:    "backing denom not found",
			req:     &types.EstimateBuyBackingInRequest{BackingOut: sdk.NewCoin("fil", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateBuyBackingInRequest{BackingOut: sdk.NewCoin("eth", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "excess backing insufficient",
			req: &types.EstimateBuyBackingInRequest{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(500000)),
			},
			expPass: false, // 5*10**5 * 0.99 / (1-0.007) > 9*10**6 * 0.99 * 1 - 8.5*10**6
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "pool backing insufficient",
			malleate: func() {
				poolBacking, found := suite.app.MakerKeeper.GetPoolBacking(suite.ctx, suite.bcDenom)
				suite.Require().True(found)
				poolBacking.Backing.Amount = sdk.ZeroInt()
				suite.app.MakerKeeper.SetPoolBacking(suite.ctx, poolBacking)
			},
			req: &types.EstimateBuyBackingInRequest{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(300000)),
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "moudle backing insufficient",
			malleate: func() {
				suite.app.BankKeeper.BurnCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(suite.bcDenom, sdk.NewInt(1000_000000))))
			},
			req: &types.EstimateBuyBackingInRequest{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(300000)),
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "correct",
			req: &types.EstimateBuyBackingInRequest{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(300000)),
			},
			expPass: true,
			expRes: &types.EstimateBuyBackingInResponse{
				FuryIn:     sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(29909_28600_000000)), // [3*10**5 / (1-0.007)] * 0.99 / 10**-10
				BuybackFee: sdk.NewCoin(suite.bcDenom, sdk.NewInt(2115)),                       // [3*10**5 / (1-0.007)] * 0.007
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateBuyBackingIn(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateBuyBackingOut() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateBuyBackingOutRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateBuyBackingOutResponse
	}{
		{
			name:    "backing denom not found",
			req:     &types.EstimateBuyBackingOutRequest{BackingDenom: "fil"},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateBuyBackingOutRequest{BackingDenom: "eth"},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "excess backing insufficient",
			req: &types.EstimateBuyBackingOutRequest{
				BackingDenom: suite.bcDenom,
				FuryIn:       sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(5000_000000_000000)),
			},
			expPass: false, // 0.5*10**16 * 10**-10 > 9*10**6 * 0.99 * 1 - 8.5*10**6
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "pool backing insufficient",
			malleate: func() {
				poolBacking, found := suite.app.MakerKeeper.GetPoolBacking(suite.ctx, suite.bcDenom)
				suite.Require().True(found)
				poolBacking.Backing.Amount = sdk.ZeroInt()
				suite.app.MakerKeeper.SetPoolBacking(suite.ctx, poolBacking)
			},
			req: &types.EstimateBuyBackingOutRequest{
				BackingDenom: suite.bcDenom,
				FuryIn:       sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(3000_000000_000000)),
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "moudle backing insufficient",
			malleate: func() {
				suite.app.BankKeeper.BurnCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(suite.bcDenom, sdk.NewInt(1000_000000))))
			},
			req: &types.EstimateBuyBackingOutRequest{
				BackingDenom: suite.bcDenom,
				FuryIn:       sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(3000_000000_000000)),
			},
			expPass: false,
			expErr:  types.ErrBackingCoinInsufficient,
		},
		{
			name: "correct",
			req: &types.EstimateBuyBackingOutRequest{
				BackingDenom: suite.bcDenom,
				FuryIn:       sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(3000_000000_000000)),
			},
			expPass: true,
			expRes: &types.EstimateBuyBackingOutResponse{
				BackingOut: sdk.NewCoin(suite.bcDenom, sdk.NewInt(300909)), // 0.3*10**16 * 10**-10 / 0.99  * (1-0.007)
				BuybackFee: sdk.NewCoin(suite.bcDenom, sdk.NewInt(2121)),   // 0.3*10**16 * 10**-10 / 0.99  * 0.007
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateBuyBackingOut(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateSellBackingIn() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateSellBackingInRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateSellBackingInResponse
	}{
		{
			name:    "backing denom not found",
			req:     &types.EstimateSellBackingInRequest{BackingDenom: "fil"},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateSellBackingInRequest{BackingDenom: "eth"},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name: "pool backing over ceiling",
			req: &types.EstimateSellBackingInRequest{
				FuryOut:      sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(20000_000000_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrBackingCeiling,
		},
		{
			name: "fury insufficient",
			req: &types.EstimateSellBackingInRequest{
				FuryOut:      sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: false,
			expErr:  types.ErrFuryCoinInsufficient,
		},
		{
			name: "correct",
			malleate: func() {
				totalBacking, found := suite.app.MakerKeeper.GetTotalBacking(suite.ctx)
				suite.Require().True(found)
				totalBacking.BlackMinted.Amount = sdk.NewInt(10_000000)
				suite.app.MakerKeeper.SetTotalBacking(suite.ctx, totalBacking)
			},
			req: &types.EstimateSellBackingInRequest{
				FuryOut:      sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(10000_000000_000000)),
				BackingDenom: suite.bcDenom,
			},
			expPass: true,
			expRes: &types.EstimateSellBackingInResponse{
				BackingIn:   sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_006578)),                 // 1*10**16 / (1+0.0075-0.004) * 10**-10 / 0.99
				SellbackFee: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(39_860488_290982)), // 1*10**16 / (1+0.0075-0.004) * 0.004
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateSellBackingIn(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestEstimateSellBackingOut() {
	testCases := []struct {
		name     string
		malleate func()
		req      *types.EstimateSellBackingOutRequest
		expPass  bool
		expErr   error
		expRes   *types.EstimateSellBackingOutResponse
	}{
		{
			name:    "backing denom not found",
			req:     &types.EstimateSellBackingOutRequest{BackingIn: sdk.NewCoin("fil", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinNotFound,
		},
		{
			name:    "backing denom disabled",
			req:     &types.EstimateSellBackingOutRequest{BackingIn: sdk.NewCoin("eth", sdk.ZeroInt())},
			expPass: false,
			expErr:  types.ErrBackingCoinDisabled,
		},
		{
			name:    "pool backing over ceiling",
			req:     &types.EstimateSellBackingOutRequest{BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(2_000000))},
			expPass: false,
			expErr:  types.ErrBackingCeiling,
		},
		{
			name:    "fury insufficient",
			req:     &types.EstimateSellBackingOutRequest{BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000))},
			expPass: false,
			expErr:  types.ErrFuryCoinInsufficient,
		},
		{
			name: "correct",
			malleate: func() {
				totalBacking, found := suite.app.MakerKeeper.GetTotalBacking(suite.ctx)
				suite.Require().True(found)
				totalBacking.BlackMinted.Amount = sdk.NewInt(10_000000)
				suite.app.MakerKeeper.SetTotalBacking(suite.ctx, totalBacking)
			},
			req:     &types.EstimateSellBackingOutRequest{BackingIn: sdk.NewCoin(suite.bcDenom, sdk.NewInt(1_000000))},
			expPass: true,
			expRes: &types.EstimateSellBackingOutResponse{
				FuryOut:     sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(9934_650000_000000)), // 1*10**6 * 0.99 / 10**-10 * (1+0.0075-0.004)
				SellbackFee: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(39_600000_000000)),   // 1*10**6 * 0.99 / 10**-10 * 0.004
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			suite.setupEstimationTest()
			if tc.malleate != nil {
				tc.malleate()
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.queryClient.EstimateSellBackingOut(ctx, tc.req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRes, res)
			} else {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expErr)
			}
		})
	}
}

func (suite *KeeperTestSuite) setupEstimationTest() {
	// set prices
	suite.app.OracleKeeper.SetExchangeRate(suite.ctx, suite.bcDenom, sdk.NewDecWithPrec(99, 2))
	suite.app.OracleKeeper.SetExchangeRate(suite.ctx, "eth", sdk.NewDec(1000_000000))
	suite.app.OracleKeeper.SetExchangeRate(suite.ctx, "fil", sdk.NewDec(5_000000))
	suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.AttoFuryDenom, sdk.NewDecWithPrec(100, 12))
	suite.app.OracleKeeper.SetExchangeRate(suite.ctx, blackfury.MicroUSBXDenom, sdk.NewDecWithPrec(1005, 3))

	// set risk params
	brp, brp2 := suite.dummyBackingRiskParams()
	suite.app.MakerKeeper.SetBackingRiskParams(suite.ctx, brp)
	suite.app.MakerKeeper.SetBackingRiskParams(suite.ctx, brp2)

	crp, crp2 := suite.dummyCollateralRiskParams()
	suite.app.MakerKeeper.SetCollateralRiskParams(suite.ctx, crp)
	suite.app.MakerKeeper.SetCollateralRiskParams(suite.ctx, crp2)

	// set pool and total backing
	suite.app.MakerKeeper.SetPoolBacking(suite.ctx, types.PoolBacking{
		BlackMinted:  sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(8_000000)),
		Backing:    sdk.NewCoin(suite.bcDenom, sdk.NewInt(9_000000)),
		FuryBurned: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
	})
	suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(suite.bcDenom, sdk.NewInt(1000_000000))))
	suite.app.MakerKeeper.SetTotalBacking(suite.ctx, types.TotalBacking{
		BlackMinted:  sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(8_500000)),
		FuryBurned: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
	})

	// set account, pool and total collateral
	suite.app.MakerKeeper.SetAccountCollateral(suite.ctx, suite.accAddress, types.AccountCollateral{
		Account:             suite.accAddress.String(),
		Collateral:          sdk.NewCoin(suite.bcDenom, sdk.NewInt(10_000000)),
		BlackDebt:             sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(6_000000)),
		FuryCollateralized:  sdk.NewCoin(blackfury.AttoFuryDenom, sdk.NewInt(3e15)),
		LastInterest:        sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.ZeroInt()),
		LastSettlementBlock: 0,
	})
	suite.app.MakerKeeper.SetPoolCollateral(suite.ctx, types.PoolCollateral{
		Collateral:         sdk.NewCoin(suite.bcDenom, sdk.NewInt(15_000000)),
		BlackDebt:            sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(8_000000)),
		FuryCollateralized: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
	})
	suite.app.MakerKeeper.SetTotalCollateral(suite.ctx, types.TotalCollateral{
		BlackDebt:            sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.NewInt(10_000000)),
		FuryCollateralized: sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt()),
	})
}
