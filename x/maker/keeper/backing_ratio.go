package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	blackfury "github.com/elysium-station/blackfury/types"
	"github.com/elysium-station/blackfury/x/maker/types"
)

// AdjustBackingRatio dynamically adjusts the backing ratio, according to black price change.
func (k Keeper) AdjustBackingRatio(ctx sdk.Context) {
	// check cooldown period since last update
	if ctx.BlockHeight()-k.GetBackingRatioLastBlock(ctx) < k.BackingRatioCooldownPeriod(ctx) {
		return
	}

	ratioStep := k.BackingRatioStep(ctx)
	if ratioStep.IsZero() {
		return
	}
	backingRatio := k.GetBackingRatio(ctx)
	priceBand := blackfury.MicroUSBXTarget.Mul(k.BackingRatioPriceBand(ctx))

	blackPrice, err := k.oracleKeeper.GetExchangeRate(ctx, blackfury.MicroUSBXDenom)
	if err != nil {
		panic(err)
	}

	if blackPrice.GT(blackfury.MicroUSBXTarget.Add(priceBand)) {
		// black price is too high
		// decrease backing ratio; min 0%
		backingRatio = sdk.MaxDec(backingRatio.Sub(ratioStep), sdk.ZeroDec())
	} else if blackPrice.LT(blackfury.MicroUSBXTarget.Sub(priceBand)) {
		// black price is too low
		// increase backing ratio; max 100%
		backingRatio = sdk.MinDec(backingRatio.Add(ratioStep), sdk.OneDec())
	}

	// TODO: consider adjusting BR based on total minted Black, even though Black price is within the band

	k.SetBackingRatio(ctx, backingRatio)
	k.SetBackingRatioLastBlock(ctx, ctx.BlockHeight())
}

func (k Keeper) SetBackingRatio(ctx sdk.Context, br sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: br})
	store.Set(types.KeyPrefixBackingRatio, bz)
}

func (k Keeper) GetBackingRatio(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixBackingRatio)
	if bz == nil {
		return sdk.OneDec()
	}
	dp := sdk.DecProto{}
	k.cdc.MustUnmarshal(bz, &dp)
	return dp.Dec
}

func (k Keeper) SetBackingRatioLastBlock(ctx sdk.Context, bh int64) {
	store := ctx.KVStore(k.storeKey)
	if bh < 0 {
		panic("invalid block height")
	}
	bz := sdk.Uint64ToBigEndian(uint64(bh))
	store.Set(types.KeyPrefixBackingRatioLastBlock, bz)
}

func (k Keeper) GetBackingRatioLastBlock(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixBackingRatioLastBlock)
	if bz == nil {
		return 0
	}
	return int64(sdk.BigEndianToUint64(bz))
}
