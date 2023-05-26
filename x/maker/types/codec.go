package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintBySwap{}, "blackfury/MsgMintBySwap", nil)
	cdc.RegisterConcrete(&MsgBurnBySwap{}, "blackfury/MsgBurnBySwap", nil)
	cdc.RegisterConcrete(&MsgBuyBacking{}, "blackfury/MsgBuyBacking", nil)
	cdc.RegisterConcrete(&MsgSellBacking{}, "blackfury/MsgSellBacking", nil)
	cdc.RegisterConcrete(&MsgMintByCollateral{}, "blackfury/MsgMintByCollateral", nil)
	cdc.RegisterConcrete(&MsgBurnByCollateral{}, "blackfury/MsgBurnByCollateral", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "blackfury/MsgDepositCollateral", nil)
	cdc.RegisterConcrete(&MsgRedeemCollateral{}, "blackfury/MsgRedeemCollateral", nil)
	cdc.RegisterConcrete(&MsgLiquidateCollateral{}, "blackfury/MsgLiquidateCollateral", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RegisterBackingProposal{},
		&RegisterCollateralProposal{},
		&SetBackingRiskParamsProposal{},
		&SetCollateralRiskParamsProposal{},
		&BatchSetBackingRiskParamsProposal{},
		&BatchSetCollateralRiskParamsProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
}
