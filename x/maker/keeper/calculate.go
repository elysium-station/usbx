package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	blackfury "github.com/elysium-station/blackfury/types"
	"github.com/elysium-station/blackfury/x/maker/types"
)

func (k Keeper) calculateMintBySwapIn(
	ctx sdk.Context,
	mintOut sdk.Coin,
	backingDenom string,
	fullBacking bool,
) (
	backingIn sdk.Coin,
	furyIn sdk.Coin,
	mintFee sdk.Coin,
	err error,
) {
	backingIn = sdk.NewCoin(backingDenom, sdk.ZeroInt())
	furyIn = sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt())
	mintFee = sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.ZeroInt())

	err = k.checkMintPriceLowerBound(ctx)
	if err != nil {
		return
	}

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	mintFee = computeFee(mintOut, backingParams.MintFee)
	mintTotal := mintOut.Add(mintFee)
	mintTotalInUSD := mintTotal.Amount.ToDec().Mul(blackfury.MicroUSBXTarget)

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}
	poolBacking.BlackMinted = poolBacking.BlackMinted.Add(mintTotal)
	if backingParams.MaxBlackMint != nil && poolBacking.BlackMinted.Amount.GT(*backingParams.MaxBlackMint) {
		err = sdkerrors.Wrapf(types.ErrBlackCeiling, "black over ceiling")
		return
	}

	backingRatio := k.GetBackingRatio(ctx)
	if backingRatio.GTE(sdk.OneDec()) || fullBacking {
		// full/over backing, or user selects full backing
		backingIn.Amount = mintTotalInUSD.QuoRoundUp(backingPrice).RoundInt()
	} else if backingRatio.IsZero() {
		// full algorithmic
		furyIn.Amount = mintTotalInUSD.QuoRoundUp(furyPrice).RoundInt()
	} else {
		// fractional
		backingIn.Amount = mintTotalInUSD.Mul(backingRatio).QuoRoundUp(backingPrice).RoundInt()
		furyIn.Amount = mintTotalInUSD.Mul(sdk.OneDec().Sub(backingRatio)).QuoRoundUp(furyPrice).RoundInt()
	}

	poolBacking.Backing = poolBacking.Backing.Add(backingIn)
	if backingParams.MaxBacking != nil && poolBacking.Backing.Amount.GT(*backingParams.MaxBacking) {
		err = sdkerrors.Wrapf(types.ErrBackingCeiling, "backing over ceiling")
		return
	}

	return
}

func (k Keeper) calculateMintBySwapOut(
	ctx sdk.Context,
	backingInMax sdk.Coin,
	furyInMax sdk.Coin,
	fullBacking bool,
) (
	backingIn sdk.Coin,
	furyIn sdk.Coin,
	mintOut sdk.Coin,
	mintFee sdk.Coin,
	err error,
) {
	backingDenom := backingInMax.Denom

	err = k.checkMintPriceLowerBound(ctx)
	if err != nil {
		return
	}

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in uusd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	backingRatio := k.GetBackingRatio(ctx)

	backingInMaxInUSD := backingPrice.MulInt(backingInMax.Amount)
	furyInMaxInUSD := furyPrice.MulInt(furyInMax.Amount)

	mintTotalInUSD := sdk.ZeroDec()
	backingIn = sdk.NewCoin(backingDenom, sdk.ZeroInt())
	furyIn = sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt())

	if backingRatio.GTE(sdk.OneDec()) || fullBacking {
		// full/over backing, or user selects full backing
		mintTotalInUSD = backingInMaxInUSD
		backingIn.Amount = backingInMax.Amount
	} else if backingRatio.IsZero() {
		// full algorithmic
		mintTotalInUSD = furyInMaxInUSD
		furyIn.Amount = furyInMax.Amount
	} else {
		// fractional
		max1 := backingInMaxInUSD.Quo(backingRatio)
		max2 := furyInMaxInUSD.Quo(sdk.OneDec().Sub(backingRatio))
		if backingInMax.IsPositive() && (furyInMax.IsZero() || max1.LTE(max2)) {
			mintTotalInUSD = max1
			backingIn.Amount = backingInMax.Amount
			furyIn.Amount = mintTotalInUSD.Mul(sdk.OneDec().Sub(backingRatio)).QuoRoundUp(furyPrice).RoundInt()
			if furyInMax.IsPositive() && furyInMax.IsLT(furyIn) {
				furyIn.Amount = furyInMax.Amount
			}
		} else {
			mintTotalInUSD = max2
			furyIn.Amount = furyInMax.Amount
			backingIn.Amount = mintTotalInUSD.Mul(backingRatio).QuoRoundUp(backingPrice).RoundInt()
			if backingInMax.IsPositive() && backingInMax.IsLT(backingIn) {
				backingIn.Amount = backingInMax.Amount
			}
		}
	}

	mintTotal := sdk.NewCoin(blackfury.MicroUSBXDenom, mintTotalInUSD.Quo(blackfury.MicroUSBXTarget).TruncateInt())

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}

	poolBacking.BlackMinted = poolBacking.BlackMinted.AddAmount(mintTotal.Amount)
	if backingParams.MaxBlackMint != nil && poolBacking.BlackMinted.Amount.GT(*backingParams.MaxBlackMint) {
		err = sdkerrors.Wrap(types.ErrBlackCeiling, "")
		return
	}

	poolBacking.Backing = poolBacking.Backing.Add(backingIn)
	if backingParams.MaxBacking != nil && poolBacking.Backing.Amount.GT(*backingParams.MaxBacking) {
		err = sdkerrors.Wrap(types.ErrBackingCeiling, "")
		return
	}

	mintFee = computeFee(mintTotal, backingParams.MintFee)
	mintOut = mintTotal.Sub(mintFee)
	return
}

func (k Keeper) calculateBurnBySwapIn(
	ctx sdk.Context,
	backingOutMax sdk.Coin,
	furyOutMax sdk.Coin,
) (
	burnIn sdk.Coin,
	backingOut sdk.Coin,
	furyOut sdk.Coin,
	burnFee sdk.Coin,
	err error,
) {
	backingDenom := backingOutMax.Denom

	burnIn = sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.ZeroInt())
	backingOut = sdk.NewCoin(backingOutMax.Denom, sdk.ZeroInt())
	furyOut = sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt())
	burnFee = sdk.NewCoin(blackfury.MicroUSBXDenom, sdk.ZeroInt())

	err = k.checkBurnPriceUpperBound(ctx)
	if err != nil {
		return
	}

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	backingOutMaxInUSD := backingPrice.MulInt(backingOutMax.Amount)
	furyOutMaxInUSD := furyPrice.MulInt(furyOutMax.Amount)

	burnActualInUSD := sdk.ZeroDec()
	backingRatio := k.GetBackingRatio(ctx)
	if backingRatio.GTE(sdk.OneDec()) {
		// full/over backing
		burnActualInUSD = backingOutMaxInUSD
		backingOut.Amount = backingOutMax.Amount
	} else if backingRatio.IsZero() {
		// full algorithmic
		burnActualInUSD = furyOutMaxInUSD
		furyOut.Amount = furyOutMax.Amount
	} else {
		// fractional
		burnActualWithBackingInUSD := backingOutMaxInUSD.Quo(backingRatio)
		burnActualWithFuryInUSD := furyOutMaxInUSD.Quo(sdk.OneDec().Sub(backingRatio))
		if furyOutMax.IsZero() || (backingOutMax.IsPositive() && burnActualWithBackingInUSD.LT(burnActualWithFuryInUSD)) {
			burnActualInUSD = burnActualWithBackingInUSD
			backingOut.Amount = backingOutMax.Amount
			furyOut.Amount = burnActualInUSD.Mul(sdk.OneDec().Sub(backingRatio)).QuoRoundUp(furyPrice).RoundInt()
		} else {
			burnActualInUSD = burnActualWithFuryInUSD
			furyOut.Amount = furyOutMax.Amount
			backingOut.Amount = burnActualInUSD.Mul(backingRatio).QuoRoundUp(backingPrice).RoundInt()
		}
	}

	moduleOwnedBacking := k.bankKeeper.GetBalance(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), backingDenom)
	if moduleOwnedBacking.IsLT(backingOut) {
		err = sdkerrors.Wrapf(types.ErrBackingCoinInsufficient, "backing coin out(%s) < balance(%s)", backingOut, moduleOwnedBacking)
		return
	}

	burnFeeRate := sdk.ZeroDec()
	if backingParams.BurnFee != nil {
		burnFeeRate = *backingParams.BurnFee
	}

	burnInValue := burnActualInUSD.Quo(blackfury.MicroUSBXTarget).Quo(sdk.OneDec().Sub(burnFeeRate))
	burnFeeValue := burnInValue.Mul(burnFeeRate)
	burnIn = sdk.NewCoin(blackfury.MicroUSBXDenom, burnInValue.RoundInt())
	burnFee = sdk.NewCoin(blackfury.MicroUSBXDenom, burnFeeValue.RoundInt())
	return
}

func (k Keeper) calculateBurnBySwapOut(
	ctx sdk.Context,
	burnIn sdk.Coin,
	backingDenom string,
) (
	backingOut sdk.Coin,
	furyOut sdk.Coin,
	burnFee sdk.Coin,
	err error,
) {
	err = k.checkBurnPriceUpperBound(ctx)
	if err != nil {
		return
	}

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	backingRatio := k.GetBackingRatio(ctx)

	burnFee = computeFee(burnIn, backingParams.BurnFee)
	burnActual := burnIn.Sub(burnFee)
	burnActualInUSD := burnActual.Amount.ToDec().Mul(blackfury.MicroUSBXTarget)

	backingOut = sdk.NewCoin(backingDenom, sdk.ZeroInt())
	furyOut = sdk.NewCoin(blackfury.AttoFuryDenom, sdk.ZeroInt())

	if backingRatio.GTE(sdk.OneDec()) {
		// full/over backing
		backingOut.Amount = burnActualInUSD.QuoTruncate(backingPrice).TruncateInt()
	} else if backingRatio.IsZero() {
		// full algorithmic
		furyOut.Amount = burnActualInUSD.QuoTruncate(furyPrice).TruncateInt()
	} else {
		// fractional
		backingOut.Amount = burnActualInUSD.Mul(backingRatio).QuoTruncate(backingPrice).TruncateInt()
		furyOut.Amount = burnActualInUSD.Mul(sdk.OneDec().Sub(backingRatio)).QuoTruncate(furyPrice).TruncateInt()
	}

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}
	moduleOwnedBacking := k.bankKeeper.GetBalance(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), backingDenom)

	poolBackingBalance := sdk.NewCoin(backingDenom, sdk.MinInt(poolBacking.Backing.Amount, moduleOwnedBacking.Amount))
	if poolBackingBalance.IsLT(backingOut) {
		err = sdkerrors.Wrapf(types.ErrBackingCoinInsufficient, "backing coin out(%s) > balance(%s)", backingOut, poolBackingBalance)
		return
	}

	return
}

func (k Keeper) calculateBuyBackingIn(
	ctx sdk.Context,
	backingOut sdk.Coin,
) (
	furyIn sdk.Coin,
	buybackFee sdk.Coin,
	err error,
) {
	backingDenom := backingOut.Denom

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	excessBackingValue, err := k.getExcessBackingValue(ctx)
	if err != nil {
		return
	}

	backingOutTotal := sdk.NewCoin(backingDenom, backingOut.Amount.ToDec().Quo(sdk.OneDec().Sub(*backingParams.BuybackFee)).TruncateInt())
	furyInValue := backingOutTotal.Amount.ToDec().Mul(backingPrice)

	if furyInValue.GT(excessBackingValue.ToDec()) {
		err = sdkerrors.Wrap(types.ErrBackingCoinInsufficient, "")
		return
	}

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}
	moduleOwnedBacking := k.bankKeeper.GetBalance(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), backingDenom)

	poolBackingBalance := sdk.NewCoin(backingDenom, sdk.MinInt(poolBacking.Backing.Amount, moduleOwnedBacking.Amount))
	if poolBackingBalance.IsLT(backingOutTotal) {
		err = sdkerrors.Wrapf(types.ErrBackingCoinInsufficient, "backing coin out(%s) > balance(%s)", backingOutTotal, poolBackingBalance)
		return
	}

	furyIn = sdk.NewCoin(blackfury.AttoFuryDenom, furyInValue.Quo(furyPrice).RoundInt())
	buybackFee = sdk.NewCoin(backingDenom, backingOutTotal.Amount.ToDec().Mul(*backingParams.BuybackFee).RoundInt())
	return
}

func (k Keeper) calculateBuyBackingOut(
	ctx sdk.Context,
	furyIn sdk.Coin,
	backingDenom string,
) (
	backingOut sdk.Coin,
	buybackFee sdk.Coin,
	err error,
) {
	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	excessBackingValue, err := k.getExcessBackingValue(ctx)
	if err != nil {
		return
	}

	furyInValue := furyIn.Amount.ToDec().Mul(furyPrice)
	if furyInValue.GT(excessBackingValue.ToDec()) {
		err = sdkerrors.Wrap(types.ErrBackingCoinInsufficient, "")
		return
	}

	backingOutTotal := sdk.NewCoin(backingDenom, furyInValue.Quo(backingPrice).TruncateInt())

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}
	moduleOwnedBacking := k.bankKeeper.GetBalance(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), backingDenom)

	poolBackingBalance := sdk.NewCoin(backingDenom, sdk.MinInt(poolBacking.Backing.Amount, moduleOwnedBacking.Amount))
	if poolBackingBalance.IsLT(backingOutTotal) {
		err = sdkerrors.Wrapf(types.ErrBackingCoinInsufficient, "backing coin out(%s) > balance(%s)", backingOutTotal, poolBackingBalance)
		return
	}

	buybackFee = computeFee(backingOutTotal, backingParams.BuybackFee)
	backingOut = backingOutTotal.Sub(buybackFee)
	return
}

func (k Keeper) calculateSellBackingIn(
	ctx sdk.Context,
	furyOut sdk.Coin,
	backingDenom string,
) (
	backingIn sdk.Coin,
	rebackFee sdk.Coin,
	err error,
) {
	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}

	excessBackingValue, err := k.getExcessBackingValue(ctx)
	if err != nil {
		return
	}
	missingBackingValue := excessBackingValue.Neg()
	availableFuryMint := missingBackingValue.ToDec().Quo(furyPrice)

	bonusRatio := k.RebackBonus(ctx)

	furyMint := furyOut.Amount.ToDec().Quo(sdk.OneDec().Add(bonusRatio).Sub(*backingParams.RebackFee))

	backingIn = sdk.NewCoin(backingDenom, furyMint.Mul(furyPrice).Quo(backingPrice).RoundInt())
	rebackFee = sdk.NewCoin(blackfury.AttoFuryDenom, furyMint.Mul(*backingParams.RebackFee).RoundInt())

	poolBacking.Backing = poolBacking.Backing.Add(backingIn)
	if backingParams.MaxBacking != nil && poolBacking.Backing.Amount.GT(*backingParams.MaxBacking) {
		err = sdkerrors.Wrap(types.ErrBackingCeiling, "")
		return
	}
	if furyMint.GT(availableFuryMint) {
		err = sdkerrors.Wrap(types.ErrFuryCoinInsufficient, "")
		return
	}

	return
}

func (k Keeper) calculateSellBackingOut(
	ctx sdk.Context,
	backingIn sdk.Coin,
) (
	furyOut sdk.Coin,
	rebackFee sdk.Coin,
	err error,
) {
	backingDenom := backingIn.Denom

	backingParams, err := k.getAvailableBackingParams(ctx, backingDenom)
	if err != nil {
		return
	}

	// get prices in usd
	backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, backingDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	_, poolBacking, err := k.getBacking(ctx, backingDenom)
	if err != nil {
		return
	}

	poolBacking.Backing = poolBacking.Backing.Add(backingIn)
	if backingParams.MaxBacking != nil && poolBacking.Backing.Amount.GT(*backingParams.MaxBacking) {
		err = sdkerrors.Wrap(types.ErrBackingCeiling, "")
		return
	}

	excessBackingValue, err := k.getExcessBackingValue(ctx)
	if err != nil {
		return
	}
	missingBackingValue := excessBackingValue.Neg()
	availableFuryMint := missingBackingValue.ToDec().Quo(furyPrice)

	bonusRatio := k.RebackBonus(ctx)
	furyMint := sdk.NewCoin(blackfury.AttoFuryDenom, backingIn.Amount.ToDec().Mul(backingPrice).Quo(furyPrice).TruncateInt())
	bonus := computeFee(furyMint, &bonusRatio)
	rebackFee = computeFee(furyMint, backingParams.RebackFee)

	if furyMint.Amount.ToDec().GT(availableFuryMint) {
		err = sdkerrors.Wrap(types.ErrFuryCoinInsufficient, "")
		return
	}

	furyOut = furyMint.Add(bonus).Sub(rebackFee)
	return
}

func (k Keeper) calculateMintByCollateral(
	ctx sdk.Context,
	account sdk.AccAddress,
	collateralDenom string,
	mintOut sdk.Coin,
) (
	mintFee sdk.Coin,
	totalColl types.TotalCollateral,
	poolColl types.PoolCollateral,
	accColl types.AccountCollateral,
	err error,
) {
	collateralParams, err := k.getAvailableCollateralParams(ctx, collateralDenom)
	if err != nil {
		return
	}

	// get prices in usd
	collateralPrice, err := k.oracleKeeper.GetExchangeRate(ctx, collateralDenom)
	if err != nil {
		return
	}
	furyPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.AttoFuryDenom)
	if err != nil {
		return
	}

	totalColl, poolColl, accColl, err = k.getCollateral(ctx, account, collateralDenom)
	if err != nil {
		return
	}

	// settle interest fee
	settleInterestFee(ctx, &accColl, &poolColl, &totalColl, *collateralParams.InterestFee)

	// compute mint total
	mintFee = computeFee(mintOut, collateralParams.MintFee)
	mintTotal := mintOut.Add(mintFee)

	// update black debt
	accColl.BlackDebt = accColl.BlackDebt.Add(mintTotal)
	poolColl.BlackDebt = poolColl.BlackDebt.Add(mintTotal)
	totalColl.BlackDebt = totalColl.BlackDebt.Add(mintTotal)

	if collateralParams.MaxBlackMint != nil && poolColl.BlackDebt.Amount.GT(*collateralParams.MaxBlackMint) {
		err = sdkerrors.Wrapf(types.ErrBlackCeiling, "")
		return
	}

	collateralValue := accColl.Collateral.Amount.ToDec().Mul(collateralPrice)
	furyCollateralizedValue := accColl.FuryCollateralized.Amount.ToDec().Mul(furyPrice)
	if !collateralValue.IsPositive() {
		err = sdkerrors.Wrapf(types.ErrAccountInsufficientCollateral, "")
		return
	}

	actualCatalyticRatio := sdk.MinDec(furyCollateralizedValue.Quo(collateralValue), *collateralParams.CatalyticFuryRatio)

	// actualCatalyticRatio / catalyticRatio = (availableLTV - basicLTV) / (maxLTV - basicLTV)
	availableLTV := *collateralParams.BasicLoanToValue
	if collateralParams.CatalyticFuryRatio.IsPositive() {
		availableLTV = availableLTV.Add(actualCatalyticRatio.Mul(collateralParams.LoanToValue.Sub(*collateralParams.BasicLoanToValue)).Quo(*collateralParams.CatalyticFuryRatio))
	}
	availableDebtMax := collateralValue.Mul(availableLTV).Quo(blackfury.MicroUSBXTarget).TruncateInt()

	if availableDebtMax.LT(accColl.BlackDebt.Amount) {
		err = sdkerrors.Wrapf(types.ErrAccountInsufficientCollateral, "")
		return
	}

	return
}

func computeFee(coin sdk.Coin, rate *sdk.Dec) sdk.Coin {
	amt := sdk.ZeroInt()
	if rate != nil {
		amt = coin.Amount.ToDec().Mul(*rate).RoundInt()
	}
	return sdk.NewCoin(coin.Denom, amt)
}

func (k Keeper) checkMintPriceLowerBound(ctx sdk.Context) error {
	blackPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.MicroUSBXDenom)
	if err != nil {
		return err
	}
	// market price must be >= target price + mint bias
	mintPriceLowerBound := blackfury.MicroUSBXTarget.Mul(sdk.OneDec().Add(k.MintPriceBias(ctx)))
	if blackPrice.LT(mintPriceLowerBound) {
		return sdkerrors.Wrapf(types.ErrBlackPriceTooLow, "%s price too low: %s", blackfury.MicroUSBXDenom, blackPrice)
	}
	return nil
}

func (k Keeper) checkBurnPriceUpperBound(ctx sdk.Context) error {
	blackPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.MicroUSBXDenom)
	if err != nil {
		return err
	}
	// market price must be <= target price - burn bias
	burnPriceUpperBound := blackfury.MicroUSBXTarget.Mul(sdk.OneDec().Sub(k.BurnPriceBias(ctx)))
	if blackPrice.GT(burnPriceUpperBound) {
		return sdkerrors.Wrapf(types.ErrBlackPriceTooHigh, "%s price too high: %s", blackfury.MicroUSBXDenom, blackPrice)
	}
	return nil
}

func (k Keeper) getAvailableBackingParams(ctx sdk.Context, backingDenom string) (backingParams types.BackingRiskParams, err error) {
	backingParams, found := k.GetBackingRiskParams(ctx, backingDenom)
	if !found {
		err = sdkerrors.Wrapf(types.ErrBackingCoinNotFound, "backing coin denomination not found: %s", backingDenom)
		return
	}
	if !backingParams.Enabled {
		err = sdkerrors.Wrapf(types.ErrBackingCoinDisabled, "backing coin disabled: %s", backingDenom)
		return
	}
	return
}

func (k Keeper) getAvailableCollateralParams(ctx sdk.Context, collateralDenom string) (collateralParams types.CollateralRiskParams, err error) {
	collateralParams, found := k.GetCollateralRiskParams(ctx, collateralDenom)
	if !found {
		err = sdkerrors.Wrapf(types.ErrCollateralCoinNotFound, "collateral coin denomination not found: %s", collateralDenom)
		return
	}
	if !collateralParams.Enabled {
		err = sdkerrors.Wrapf(types.ErrCollateralCoinDisabled, "collateral coin disabled: %s", collateralDenom)
		return
	}
	return
}

func (k Keeper) getExcessBackingValue(ctx sdk.Context) (excessBackingValue sdk.Int, err error) {
	totalBacking, found := k.GetTotalBacking(ctx)
	if !found {
		err = sdkerrors.Wrapf(types.ErrBackingCoinNotFound, "total backing not found")
		return
	}

	backingRatio := k.GetBackingRatio(ctx)
	requiredBackingValue := totalBacking.BlackMinted.Amount.ToDec().Mul(backingRatio).Ceil().TruncateInt()
	if requiredBackingValue.IsNegative() {
		requiredBackingValue = sdk.ZeroInt()
	}

	totalBackingValue, err := k.totalBackingInUSD(ctx)
	if err != nil {
		return
	}

	// may be negative
	excessBackingValue = totalBackingValue.Sub(requiredBackingValue)
	return
}

func (k Keeper) totalBackingInUSD(ctx sdk.Context) (sdk.Int, error) {
	totalBackingValue := sdk.ZeroDec()
	for _, pool := range k.GetAllPoolBacking(ctx) {
		// get price in usd
		backingPrice, err := k.oracleKeeper.GetExchangeRate(ctx, pool.Backing.Denom)
		if err != nil {
			return sdk.Int{}, err
		}
		totalBackingValue = totalBackingValue.Add(pool.Backing.Amount.ToDec().Mul(backingPrice))
	}
	return totalBackingValue.TruncateInt(), nil
}
