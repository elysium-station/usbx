<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [blackfury/bank/v1beta1/bank.proto](#blackfury/bank/v1beta1/bank.proto)
    - [SetDenomMetadataProposal](#blackfury.bank.v1beta1.SetDenomMetadataProposal)
  
- [blackfury/erc20/v1/erc20.proto](#blackfury/erc20/v1/erc20.proto)
    - [TokenPair](#blackfury.erc20.v1.TokenPair)
  
    - [Owner](#blackfury.erc20.v1.Owner)
  
- [blackfury/erc20/v1/genesis.proto](#blackfury/erc20/v1/genesis.proto)
    - [GenesisState](#blackfury.erc20.v1.GenesisState)
    - [Params](#blackfury.erc20.v1.Params)
  
- [blackfury/erc20/v1/query.proto](#blackfury/erc20/v1/query.proto)
    - [QueryParamsRequest](#blackfury.erc20.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.erc20.v1.QueryParamsResponse)
    - [QueryTokenPairRequest](#blackfury.erc20.v1.QueryTokenPairRequest)
    - [QueryTokenPairResponse](#blackfury.erc20.v1.QueryTokenPairResponse)
    - [QueryTokenPairsRequest](#blackfury.erc20.v1.QueryTokenPairsRequest)
    - [QueryTokenPairsResponse](#blackfury.erc20.v1.QueryTokenPairsResponse)
  
    - [Query](#blackfury.erc20.v1.Query)
  
- [blackfury/erc20/v1/tx.proto](#blackfury/erc20/v1/tx.proto)
    - [Msg](#blackfury.erc20.v1.Msg)
  
- [blackfury/gauge/v1/event.proto](#blackfury/gauge/v1/event.proto)
- [blackfury/gauge/v1/gauge.proto](#blackfury/gauge/v1/gauge.proto)
    - [Checkpoint](#blackfury.gauge.v1.Checkpoint)
    - [Reward](#blackfury.gauge.v1.Reward)
    - [UserReward](#blackfury.gauge.v1.UserReward)
  
- [blackfury/gauge/v1/genesis.proto](#blackfury/gauge/v1/genesis.proto)
    - [GenesisState](#blackfury.gauge.v1.GenesisState)
    - [Params](#blackfury.gauge.v1.Params)
  
- [blackfury/gauge/v1/query.proto](#blackfury/gauge/v1/query.proto)
    - [QueryParamsRequest](#blackfury.gauge.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.gauge.v1.QueryParamsResponse)
  
    - [Query](#blackfury.gauge.v1.Query)
  
- [blackfury/gauge/v1/tx.proto](#blackfury/gauge/v1/tx.proto)
    - [Msg](#blackfury.gauge.v1.Msg)
  
- [blackfury/maker/v1/genesis.proto](#blackfury/maker/v1/genesis.proto)
    - [GenesisState](#blackfury.maker.v1.GenesisState)
    - [Params](#blackfury.maker.v1.Params)
  
- [blackfury/maker/v1/maker.proto](#blackfury/maker/v1/maker.proto)
    - [AccountBacking](#blackfury.maker.v1.AccountBacking)
    - [AccountCollateral](#blackfury.maker.v1.AccountCollateral)
    - [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams)
    - [BatchBackingRiskParams](#blackfury.maker.v1.BatchBackingRiskParams)
    - [BatchCollateralRiskParams](#blackfury.maker.v1.BatchCollateralRiskParams)
    - [BatchSetBackingRiskParamsProposal](#blackfury.maker.v1.BatchSetBackingRiskParamsProposal)
    - [BatchSetCollateralRiskParamsProposal](#blackfury.maker.v1.BatchSetCollateralRiskParamsProposal)
    - [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams)
    - [PoolBacking](#blackfury.maker.v1.PoolBacking)
    - [PoolCollateral](#blackfury.maker.v1.PoolCollateral)
    - [RegisterBackingProposal](#blackfury.maker.v1.RegisterBackingProposal)
    - [RegisterCollateralProposal](#blackfury.maker.v1.RegisterCollateralProposal)
    - [SetBackingRiskParamsProposal](#blackfury.maker.v1.SetBackingRiskParamsProposal)
    - [SetCollateralRiskParamsProposal](#blackfury.maker.v1.SetCollateralRiskParamsProposal)
    - [TotalBacking](#blackfury.maker.v1.TotalBacking)
    - [TotalCollateral](#blackfury.maker.v1.TotalCollateral)
  
- [blackfury/maker/v1/query.proto](#blackfury/maker/v1/query.proto)
    - [EstimateBurnBySwapInRequest](#blackfury.maker.v1.EstimateBurnBySwapInRequest)
    - [EstimateBurnBySwapInResponse](#blackfury.maker.v1.EstimateBurnBySwapInResponse)
    - [EstimateBurnBySwapOutRequest](#blackfury.maker.v1.EstimateBurnBySwapOutRequest)
    - [EstimateBurnBySwapOutResponse](#blackfury.maker.v1.EstimateBurnBySwapOutResponse)
    - [EstimateBuyBackingInRequest](#blackfury.maker.v1.EstimateBuyBackingInRequest)
    - [EstimateBuyBackingInResponse](#blackfury.maker.v1.EstimateBuyBackingInResponse)
    - [EstimateBuyBackingOutRequest](#blackfury.maker.v1.EstimateBuyBackingOutRequest)
    - [EstimateBuyBackingOutResponse](#blackfury.maker.v1.EstimateBuyBackingOutResponse)
    - [EstimateMintBySwapInRequest](#blackfury.maker.v1.EstimateMintBySwapInRequest)
    - [EstimateMintBySwapInResponse](#blackfury.maker.v1.EstimateMintBySwapInResponse)
    - [EstimateMintBySwapOutRequest](#blackfury.maker.v1.EstimateMintBySwapOutRequest)
    - [EstimateMintBySwapOutResponse](#blackfury.maker.v1.EstimateMintBySwapOutResponse)
    - [EstimateSellBackingInRequest](#blackfury.maker.v1.EstimateSellBackingInRequest)
    - [EstimateSellBackingInResponse](#blackfury.maker.v1.EstimateSellBackingInResponse)
    - [EstimateSellBackingOutRequest](#blackfury.maker.v1.EstimateSellBackingOutRequest)
    - [EstimateSellBackingOutResponse](#blackfury.maker.v1.EstimateSellBackingOutResponse)
    - [QueryAllBackingPoolsRequest](#blackfury.maker.v1.QueryAllBackingPoolsRequest)
    - [QueryAllBackingPoolsResponse](#blackfury.maker.v1.QueryAllBackingPoolsResponse)
    - [QueryAllBackingRiskParamsRequest](#blackfury.maker.v1.QueryAllBackingRiskParamsRequest)
    - [QueryAllBackingRiskParamsResponse](#blackfury.maker.v1.QueryAllBackingRiskParamsResponse)
    - [QueryAllCollateralPoolsRequest](#blackfury.maker.v1.QueryAllCollateralPoolsRequest)
    - [QueryAllCollateralPoolsResponse](#blackfury.maker.v1.QueryAllCollateralPoolsResponse)
    - [QueryAllCollateralRiskParamsRequest](#blackfury.maker.v1.QueryAllCollateralRiskParamsRequest)
    - [QueryAllCollateralRiskParamsResponse](#blackfury.maker.v1.QueryAllCollateralRiskParamsResponse)
    - [QueryBackingPoolRequest](#blackfury.maker.v1.QueryBackingPoolRequest)
    - [QueryBackingPoolResponse](#blackfury.maker.v1.QueryBackingPoolResponse)
    - [QueryBackingRatioRequest](#blackfury.maker.v1.QueryBackingRatioRequest)
    - [QueryBackingRatioResponse](#blackfury.maker.v1.QueryBackingRatioResponse)
    - [QueryCollateralOfAccountRequest](#blackfury.maker.v1.QueryCollateralOfAccountRequest)
    - [QueryCollateralOfAccountResponse](#blackfury.maker.v1.QueryCollateralOfAccountResponse)
    - [QueryCollateralPoolRequest](#blackfury.maker.v1.QueryCollateralPoolRequest)
    - [QueryCollateralPoolResponse](#blackfury.maker.v1.QueryCollateralPoolResponse)
    - [QueryParamsRequest](#blackfury.maker.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.maker.v1.QueryParamsResponse)
    - [QueryTotalBackingRequest](#blackfury.maker.v1.QueryTotalBackingRequest)
    - [QueryTotalBackingResponse](#blackfury.maker.v1.QueryTotalBackingResponse)
    - [QueryTotalCollateralRequest](#blackfury.maker.v1.QueryTotalCollateralRequest)
    - [QueryTotalCollateralResponse](#blackfury.maker.v1.QueryTotalCollateralResponse)
  
    - [Query](#blackfury.maker.v1.Query)
  
- [blackfury/maker/v1/tx.proto](#blackfury/maker/v1/tx.proto)
    - [MsgBurnByCollateral](#blackfury.maker.v1.MsgBurnByCollateral)
    - [MsgBurnByCollateralResponse](#blackfury.maker.v1.MsgBurnByCollateralResponse)
    - [MsgBurnBySwap](#blackfury.maker.v1.MsgBurnBySwap)
    - [MsgBurnBySwapResponse](#blackfury.maker.v1.MsgBurnBySwapResponse)
    - [MsgBuyBacking](#blackfury.maker.v1.MsgBuyBacking)
    - [MsgBuyBackingResponse](#blackfury.maker.v1.MsgBuyBackingResponse)
    - [MsgDepositCollateral](#blackfury.maker.v1.MsgDepositCollateral)
    - [MsgDepositCollateralResponse](#blackfury.maker.v1.MsgDepositCollateralResponse)
    - [MsgLiquidateCollateral](#blackfury.maker.v1.MsgLiquidateCollateral)
    - [MsgLiquidateCollateralResponse](#blackfury.maker.v1.MsgLiquidateCollateralResponse)
    - [MsgMintByCollateral](#blackfury.maker.v1.MsgMintByCollateral)
    - [MsgMintByCollateralResponse](#blackfury.maker.v1.MsgMintByCollateralResponse)
    - [MsgMintBySwap](#blackfury.maker.v1.MsgMintBySwap)
    - [MsgMintBySwapResponse](#blackfury.maker.v1.MsgMintBySwapResponse)
    - [MsgRedeemCollateral](#blackfury.maker.v1.MsgRedeemCollateral)
    - [MsgRedeemCollateralResponse](#blackfury.maker.v1.MsgRedeemCollateralResponse)
    - [MsgSellBacking](#blackfury.maker.v1.MsgSellBacking)
    - [MsgSellBackingResponse](#blackfury.maker.v1.MsgSellBackingResponse)
  
    - [Msg](#blackfury.maker.v1.Msg)
  
- [blackfury/oracle/v1/oracle.proto](#blackfury/oracle/v1/oracle.proto)
    - [AggregateExchangeRatePrevote](#blackfury.oracle.v1.AggregateExchangeRatePrevote)
    - [AggregateExchangeRateVote](#blackfury.oracle.v1.AggregateExchangeRateVote)
    - [ExchangeRateTuple](#blackfury.oracle.v1.ExchangeRateTuple)
    - [Params](#blackfury.oracle.v1.Params)
    - [RegisterTargetProposal](#blackfury.oracle.v1.RegisterTargetProposal)
    - [TargetParams](#blackfury.oracle.v1.TargetParams)
  
    - [TargetSource](#blackfury.oracle.v1.TargetSource)
  
- [blackfury/oracle/v1/genesis.proto](#blackfury/oracle/v1/genesis.proto)
    - [FeederDelegation](#blackfury.oracle.v1.FeederDelegation)
    - [GenesisState](#blackfury.oracle.v1.GenesisState)
    - [MissCounter](#blackfury.oracle.v1.MissCounter)
  
- [blackfury/oracle/v1/query.proto](#blackfury/oracle/v1/query.proto)
    - [QueryActivesRequest](#blackfury.oracle.v1.QueryActivesRequest)
    - [QueryActivesResponse](#blackfury.oracle.v1.QueryActivesResponse)
    - [QueryAggregatePrevoteRequest](#blackfury.oracle.v1.QueryAggregatePrevoteRequest)
    - [QueryAggregatePrevoteResponse](#blackfury.oracle.v1.QueryAggregatePrevoteResponse)
    - [QueryAggregatePrevotesRequest](#blackfury.oracle.v1.QueryAggregatePrevotesRequest)
    - [QueryAggregatePrevotesResponse](#blackfury.oracle.v1.QueryAggregatePrevotesResponse)
    - [QueryAggregateVoteRequest](#blackfury.oracle.v1.QueryAggregateVoteRequest)
    - [QueryAggregateVoteResponse](#blackfury.oracle.v1.QueryAggregateVoteResponse)
    - [QueryAggregateVotesRequest](#blackfury.oracle.v1.QueryAggregateVotesRequest)
    - [QueryAggregateVotesResponse](#blackfury.oracle.v1.QueryAggregateVotesResponse)
    - [QueryExchangeRateRequest](#blackfury.oracle.v1.QueryExchangeRateRequest)
    - [QueryExchangeRateResponse](#blackfury.oracle.v1.QueryExchangeRateResponse)
    - [QueryExchangeRatesRequest](#blackfury.oracle.v1.QueryExchangeRatesRequest)
    - [QueryExchangeRatesResponse](#blackfury.oracle.v1.QueryExchangeRatesResponse)
    - [QueryFeederDelegationRequest](#blackfury.oracle.v1.QueryFeederDelegationRequest)
    - [QueryFeederDelegationResponse](#blackfury.oracle.v1.QueryFeederDelegationResponse)
    - [QueryMissCounterRequest](#blackfury.oracle.v1.QueryMissCounterRequest)
    - [QueryMissCounterResponse](#blackfury.oracle.v1.QueryMissCounterResponse)
    - [QueryParamsRequest](#blackfury.oracle.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.oracle.v1.QueryParamsResponse)
    - [QueryTargetsRequest](#blackfury.oracle.v1.QueryTargetsRequest)
    - [QueryTargetsResponse](#blackfury.oracle.v1.QueryTargetsResponse)
    - [QueryVoteTargetsRequest](#blackfury.oracle.v1.QueryVoteTargetsRequest)
    - [QueryVoteTargetsResponse](#blackfury.oracle.v1.QueryVoteTargetsResponse)
  
    - [Query](#blackfury.oracle.v1.Query)
  
- [blackfury/oracle/v1/tx.proto](#blackfury/oracle/v1/tx.proto)
    - [MsgAggregateExchangeRatePrevote](#blackfury.oracle.v1.MsgAggregateExchangeRatePrevote)
    - [MsgAggregateExchangeRatePrevoteResponse](#blackfury.oracle.v1.MsgAggregateExchangeRatePrevoteResponse)
    - [MsgAggregateExchangeRateVote](#blackfury.oracle.v1.MsgAggregateExchangeRateVote)
    - [MsgAggregateExchangeRateVoteResponse](#blackfury.oracle.v1.MsgAggregateExchangeRateVoteResponse)
    - [MsgDelegateFeedConsent](#blackfury.oracle.v1.MsgDelegateFeedConsent)
    - [MsgDelegateFeedConsentResponse](#blackfury.oracle.v1.MsgDelegateFeedConsentResponse)
  
    - [Msg](#blackfury.oracle.v1.Msg)
  
- [blackfury/staking/v1/query.proto](#blackfury/staking/v1/query.proto)
- [blackfury/staking/v1/staking.proto](#blackfury/staking/v1/staking.proto)
    - [VeDelegation](#blackfury.staking.v1.VeDelegation)
    - [VeRedelegation](#blackfury.staking.v1.VeRedelegation)
    - [VeRedelegationEntry](#blackfury.staking.v1.VeRedelegationEntry)
    - [VeRedelegationEntryShares](#blackfury.staking.v1.VeRedelegationEntryShares)
    - [VeShares](#blackfury.staking.v1.VeShares)
    - [VeTokens](#blackfury.staking.v1.VeTokens)
    - [VeUnbondingDelegation](#blackfury.staking.v1.VeUnbondingDelegation)
    - [VeUnbondingDelegationEntry](#blackfury.staking.v1.VeUnbondingDelegationEntry)
    - [VeUnbondingDelegationEntryBalances](#blackfury.staking.v1.VeUnbondingDelegationEntryBalances)
    - [VeValidator](#blackfury.staking.v1.VeValidator)
  
- [blackfury/staking/v1/tx.proto](#blackfury/staking/v1/tx.proto)
    - [MsgVeDelegate](#blackfury.staking.v1.MsgVeDelegate)
    - [MsgVeDelegateResponse](#blackfury.staking.v1.MsgVeDelegateResponse)
  
    - [Msg](#blackfury.staking.v1.Msg)
  
- [blackfury/ve/v1/event.proto](#blackfury/ve/v1/event.proto)
    - [EventCreate](#blackfury.ve.v1.EventCreate)
    - [EventDeposit](#blackfury.ve.v1.EventDeposit)
    - [EventExtendTime](#blackfury.ve.v1.EventExtendTime)
    - [EventMerge](#blackfury.ve.v1.EventMerge)
    - [EventWithdraw](#blackfury.ve.v1.EventWithdraw)
  
- [blackfury/ve/v1/genesis.proto](#blackfury/ve/v1/genesis.proto)
    - [GenesisState](#blackfury.ve.v1.GenesisState)
    - [Params](#blackfury.ve.v1.Params)
  
- [blackfury/ve/v1/query.proto](#blackfury/ve/v1/query.proto)
    - [QueryParamsRequest](#blackfury.ve.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.ve.v1.QueryParamsResponse)
    - [QueryTotalVotingPowerRequest](#blackfury.ve.v1.QueryTotalVotingPowerRequest)
    - [QueryTotalVotingPowerResponse](#blackfury.ve.v1.QueryTotalVotingPowerResponse)
    - [QueryVeNftRequest](#blackfury.ve.v1.QueryVeNftRequest)
    - [QueryVeNftResponse](#blackfury.ve.v1.QueryVeNftResponse)
    - [QueryVeNftsRequest](#blackfury.ve.v1.QueryVeNftsRequest)
    - [QueryVeNftsResponse](#blackfury.ve.v1.QueryVeNftsResponse)
    - [QueryVotingPowerRequest](#blackfury.ve.v1.QueryVotingPowerRequest)
    - [QueryVotingPowerResponse](#blackfury.ve.v1.QueryVotingPowerResponse)
  
    - [Query](#blackfury.ve.v1.Query)
  
- [blackfury/ve/v1/tx.proto](#blackfury/ve/v1/tx.proto)
    - [MsgCreate](#blackfury.ve.v1.MsgCreate)
    - [MsgCreateResponse](#blackfury.ve.v1.MsgCreateResponse)
    - [MsgDeposit](#blackfury.ve.v1.MsgDeposit)
    - [MsgDepositResponse](#blackfury.ve.v1.MsgDepositResponse)
    - [MsgExtendTime](#blackfury.ve.v1.MsgExtendTime)
    - [MsgExtendTimeResponse](#blackfury.ve.v1.MsgExtendTimeResponse)
    - [MsgMerge](#blackfury.ve.v1.MsgMerge)
    - [MsgMergeResponse](#blackfury.ve.v1.MsgMergeResponse)
    - [MsgWithdraw](#blackfury.ve.v1.MsgWithdraw)
    - [MsgWithdrawResponse](#blackfury.ve.v1.MsgWithdrawResponse)
  
    - [Msg](#blackfury.ve.v1.Msg)
  
- [blackfury/ve/v1/ve.proto](#blackfury/ve/v1/ve.proto)
    - [Checkpoint](#blackfury.ve.v1.Checkpoint)
    - [LockedBalance](#blackfury.ve.v1.LockedBalance)
  
- [blackfury/vesting/v1/genesis.proto](#blackfury/vesting/v1/genesis.proto)
    - [AllocationAddresses](#blackfury.vesting.v1.AllocationAddresses)
    - [AllocationAmounts](#blackfury.vesting.v1.AllocationAmounts)
    - [GenesisState](#blackfury.vesting.v1.GenesisState)
    - [Params](#blackfury.vesting.v1.Params)
  
- [blackfury/vesting/v1/vesting.proto](#blackfury/vesting/v1/vesting.proto)
    - [Airdrop](#blackfury.vesting.v1.Airdrop)
  
- [blackfury/vesting/v1/query.proto](#blackfury/vesting/v1/query.proto)
    - [QueryAirdropRequest](#blackfury.vesting.v1.QueryAirdropRequest)
    - [QueryAirdropResponse](#blackfury.vesting.v1.QueryAirdropResponse)
    - [QueryAirdropsRequest](#blackfury.vesting.v1.QueryAirdropsRequest)
    - [QueryAirdropsResponse](#blackfury.vesting.v1.QueryAirdropsResponse)
    - [QueryParamsRequest](#blackfury.vesting.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.vesting.v1.QueryParamsResponse)
  
    - [Query](#blackfury.vesting.v1.Query)
  
- [blackfury/vesting/v1/tx.proto](#blackfury/vesting/v1/tx.proto)
    - [MsgAddAirdrops](#blackfury.vesting.v1.MsgAddAirdrops)
    - [MsgAddAirdropsResponse](#blackfury.vesting.v1.MsgAddAirdropsResponse)
    - [MsgExecuteAirdrops](#blackfury.vesting.v1.MsgExecuteAirdrops)
    - [MsgExecuteAirdropsResponse](#blackfury.vesting.v1.MsgExecuteAirdropsResponse)
    - [MsgSetAllocationAddress](#blackfury.vesting.v1.MsgSetAllocationAddress)
    - [MsgSetAllocationAddressResponse](#blackfury.vesting.v1.MsgSetAllocationAddressResponse)
  
    - [Msg](#blackfury.vesting.v1.Msg)
  
- [blackfury/voter/v1/genesis.proto](#blackfury/voter/v1/genesis.proto)
    - [GenesisState](#blackfury.voter.v1.GenesisState)
    - [Params](#blackfury.voter.v1.Params)
  
- [blackfury/voter/v1/query.proto](#blackfury/voter/v1/query.proto)
    - [QueryParamsRequest](#blackfury.voter.v1.QueryParamsRequest)
    - [QueryParamsResponse](#blackfury.voter.v1.QueryParamsResponse)
  
    - [Query](#blackfury.voter.v1.Query)
  
- [blackfury/voter/v1/tx.proto](#blackfury/voter/v1/tx.proto)
    - [Msg](#blackfury.voter.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="blackfury/bank/v1beta1/bank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/bank/v1beta1/bank.proto



<a name="blackfury.bank.v1beta1.SetDenomMetadataProposal"></a>

### SetDenomMetadataProposal
SetDenomMetaDataProposal is a gov Content type to register a DenomMetaData


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `metadata` | [cosmos.bank.v1beta1.Metadata](#cosmos.bank.v1beta1.Metadata) |  | token pair of Cosmos native denom and ERC20 token address |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/erc20/v1/erc20.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/erc20/v1/erc20.proto



<a name="blackfury.erc20.v1.TokenPair"></a>

### TokenPair
TokenPair defines an instance that records pairing consisting of a Cosmos
native Coin and an ERC20 token address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20_address` | [string](#string) |  | address of ERC20 contract token |
| `denom` | [string](#string) |  | cosmos base denomination to be mapped to |
| `contract_owner` | [Owner](#blackfury.erc20.v1.Owner) |  | ERC20 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address) |





 <!-- end messages -->


<a name="blackfury.erc20.v1.Owner"></a>

### Owner
Owner enumerates the ownership of a ERC20 contract.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNER_UNSPECIFIED | 0 | OWNER_UNSPECIFIED defines an invalid/undefined owner. |
| OWNER_MODULE | 1 | OWNER_MODULE erc20 is owned by the erc20 module account. |
| OWNER_EXTERNAL | 2 | EXTERNAL erc20 is owned by an external account. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/erc20/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/erc20/v1/genesis.proto



<a name="blackfury.erc20.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.erc20.v1.Params) |  | module parameters |
| `token_pairs` | [TokenPair](#blackfury.erc20.v1.TokenPair) | repeated | registered token pairs |






<a name="blackfury.erc20.v1.Params"></a>

### Params
Params defines the erc20 module params





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/erc20/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/erc20/v1/query.proto



<a name="blackfury.erc20.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="blackfury.erc20.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.erc20.v1.Params) |  |  |






<a name="blackfury.erc20.v1.QueryTokenPairRequest"></a>

### QueryTokenPairRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC20 or the Cosmos base denomination |






<a name="blackfury.erc20.v1.QueryTokenPairResponse"></a>

### QueryTokenPairResponse
QueryTokenPairResponse is the response type for the Query/TokenPair RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pair` | [TokenPair](#blackfury.erc20.v1.TokenPair) |  |  |






<a name="blackfury.erc20.v1.QueryTokenPairsRequest"></a>

### QueryTokenPairsRequest
QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="blackfury.erc20.v1.QueryTokenPairsResponse"></a>

### QueryTokenPairsResponse
QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pairs` | [TokenPair](#blackfury.erc20.v1.TokenPair) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.erc20.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TokenPairs` | [QueryTokenPairsRequest](#blackfury.erc20.v1.QueryTokenPairsRequest) | [QueryTokenPairsResponse](#blackfury.erc20.v1.QueryTokenPairsResponse) | Retrieves registered token pairs | GET|/blackfury/erc20/v1/token_pairs|
| `TokenPair` | [QueryTokenPairRequest](#blackfury.erc20.v1.QueryTokenPairRequest) | [QueryTokenPairResponse](#blackfury.erc20.v1.QueryTokenPairResponse) | Retrieves a registered token pair | GET|/blackfury/erc20/v1/token_pairs/{token}|
| `Params` | [QueryParamsRequest](#blackfury.erc20.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.erc20.v1.QueryParamsResponse) | Params retrieves the erc20 module params | GET|/blackfury/erc20/v1/params|

 <!-- end services -->



<a name="blackfury/erc20/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/erc20/v1/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.erc20.v1.Msg"></a>

### Msg
Msg defines the erc20 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

 <!-- end services -->



<a name="blackfury/gauge/v1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/gauge/v1/event.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/gauge/v1/gauge.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/gauge/v1/gauge.proto



<a name="blackfury.gauge.v1.Checkpoint"></a>

### Checkpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `timestamp` | [uint64](#uint64) |  | unix timestamp |
| `amount` | [string](#string) |  |  |






<a name="blackfury.gauge.v1.Reward"></a>

### Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | reward coin denom |
| `rate` | [string](#string) |  | reward amount per second |
| `finish_time` | [uint64](#uint64) |  | reward finish unix time |
| `last_update_time` | [uint64](#uint64) |  | unix time of last reward update |
| `cumulative_per_ticket` | [string](#string) |  | cumulative reward per voting ticket |
| `accrued_amount` | [string](#string) |  | accrued reward amount which has not been used for distributing rateably |






<a name="blackfury.gauge.v1.UserReward"></a>

### UserReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | reward coin denom |
| `ve_id` | [uint64](#uint64) |  | ve id |
| `last_claim_time` | [uint64](#uint64) |  | last claim unix time |
| `cumulative_per_ticket` | [string](#string) |  | cumulative reward per voting ticket |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/gauge/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/gauge/v1/genesis.proto



<a name="blackfury.gauge.v1.GenesisState"></a>

### GenesisState
GenesisState defines the gauge module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.gauge.v1.Params) |  |  |






<a name="blackfury.gauge.v1.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/gauge/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/gauge/v1/query.proto



<a name="blackfury.gauge.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.gauge.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.gauge.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.gauge.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#blackfury.gauge.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.gauge.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfuryzone/blackfury/gauge/params|

 <!-- end services -->



<a name="blackfury/gauge/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/gauge/v1/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.gauge.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

 <!-- end services -->



<a name="blackfury/maker/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/maker/v1/genesis.proto



<a name="blackfury.maker.v1.GenesisState"></a>

### GenesisState
GenesisState defines the maker module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.maker.v1.Params) |  |  |
| `backing_ratio` | [string](#string) |  |  |






<a name="blackfury.maker.v1.Params"></a>

### Params
Params defines the parameters for the maker module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_ratio_step` | [string](#string) |  | step of adjusting backing ratio |
| `backing_ratio_price_band` | [string](#string) |  | price band for adjusting backing ratio |
| `backing_ratio_cooldown_period` | [int64](#int64) |  | cooldown period for adjusting backing ratio |
| `mint_price_bias` | [string](#string) |  | mint Black price bias ratio |
| `burn_price_bias` | [string](#string) |  | burn Black price bias ratio |
| `reback_bonus` | [string](#string) |  | reback bonus ratio |
| `liquidation_commission_fee` | [string](#string) |  | liquidation commission fee ratio |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/maker/v1/maker.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/maker/v1/maker.proto



<a name="blackfury.maker.v1.AccountBacking"></a>

### AccountBacking







<a name="blackfury.maker.v1.AccountCollateral"></a>

### AccountCollateral



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account` | [string](#string) |  | account who owns collateral |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | existing collateral |
| `black_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | remaining black debt, including minted by collateral, mint fee, last interest |
| `fury_collateralized` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total collateralized fury |
| `last_interest` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | remaining interest debt at last settlement |
| `last_settlement_block` | [int64](#int64) |  | the block of last settlement |






<a name="blackfury.maker.v1.BackingRiskParams"></a>

### BackingRiskParams
BackingRiskParams represents an object of backing coin risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_denom` | [string](#string) |  | backing coin denom |
| `enabled` | [bool](#bool) |  | whether enabled |
| `max_backing` | [string](#string) |  | maximum total backing amount |
| `max_black_mint` | [string](#string) |  | maximum mintable Black amount |
| `mint_fee` | [string](#string) |  | mint fee rate |
| `burn_fee` | [string](#string) |  | burn fee rate |
| `buyback_fee` | [string](#string) |  | buyback fee rate |
| `reback_fee` | [string](#string) |  | reback fee rate |






<a name="blackfury.maker.v1.BatchBackingRiskParams"></a>

### BatchBackingRiskParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `risk_params` | [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams) | repeated | batch of collateral risk params |






<a name="blackfury.maker.v1.BatchCollateralRiskParams"></a>

### BatchCollateralRiskParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `risk_params` | [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams) | repeated | batch of collateral risk params |






<a name="blackfury.maker.v1.BatchSetBackingRiskParamsProposal"></a>

### BatchSetBackingRiskParamsProposal
BatchSetBackingRiskParamsProposal is a gov Content type to batch set backing
coin risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams) | repeated | batch of collateral risk params |






<a name="blackfury.maker.v1.BatchSetCollateralRiskParamsProposal"></a>

### BatchSetCollateralRiskParamsProposal
BatchSetCollateralRiskParamsProposal is a gov Content type to batch set
collateral risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams) | repeated | batch of collateral risk params |






<a name="blackfury.maker.v1.CollateralRiskParams"></a>

### CollateralRiskParams
CollateralRiskParams represents an object of collateral risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_denom` | [string](#string) |  | collateral coin denom |
| `enabled` | [bool](#bool) |  | whether enabled |
| `max_collateral` | [string](#string) |  | maximum total collateral amount; empty means no limit |
| `max_black_mint` | [string](#string) |  | maximum total mintable Black amount; empty means no limit |
| `liquidation_threshold` | [string](#string) |  | ratio at which a position is defined as undercollateralized |
| `loan_to_value` | [string](#string) |  | maximum ratio of maximum amount of currency that can be borrowed with a specific collateral |
| `basic_loan_to_value` | [string](#string) |  | basic ratio of maximum amount of currency that can be borrowed with a specific collateral |
| `catalytic_fury_ratio` | [string](#string) |  | catalytic ratio of collateralized Fury to asset, to maximize the LTV in [basic-LTV, LTV] |
| `liquidation_fee` | [string](#string) |  | liquidation fee rate, i.e., the discount a liquidator gets when buying collateral flagged for a liquidation |
| `mint_fee` | [string](#string) |  | mint fee rate, i.e., extra fee debt |
| `interest_fee` | [string](#string) |  | annual interest fee rate (APR) |






<a name="blackfury.maker.v1.PoolBacking"></a>

### PoolBacking



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `black_minted` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total minted black; negative value means burned black |
| `backing` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total backing |
| `fury_burned` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total burned fury; negative value means minted fury |






<a name="blackfury.maker.v1.PoolCollateral"></a>

### PoolCollateral



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total collateral |
| `black_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total existing black debt, including minted by collateral, mint fee, last interest |
| `fury_collateralized` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total collateralized fury |






<a name="blackfury.maker.v1.RegisterBackingProposal"></a>

### RegisterBackingProposal
RegisterBackingProposal is a gov Content type to register eligible
strong-backing asset with backing risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams) |  | backing risk params |






<a name="blackfury.maker.v1.RegisterCollateralProposal"></a>

### RegisterCollateralProposal
RegisterCollateralProposal is a gov Content type to register eligible
collateral with collateral risk parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams) |  | collateral risk params |






<a name="blackfury.maker.v1.SetBackingRiskParamsProposal"></a>

### SetBackingRiskParamsProposal
SetBackingRiskParamsProposal is a gov Content type to set backing coin risk
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams) |  | backing risk params |






<a name="blackfury.maker.v1.SetCollateralRiskParamsProposal"></a>

### SetCollateralRiskParamsProposal
SetCollateralRiskParamsProposal is a gov Content type to set collateral risk
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `risk_params` | [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams) |  | collateral risk params |






<a name="blackfury.maker.v1.TotalBacking"></a>

### TotalBacking



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_value` | [string](#string) |  | total backing value in uUSD |
| `black_minted` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total minted black; negative value means burned black |
| `fury_burned` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total burned fury; negative value means minted fury |






<a name="blackfury.maker.v1.TotalCollateral"></a>

### TotalCollateral



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `black_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total existing black debt, including minted by collateral, mint fee, last interest |
| `fury_collateralized` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total collateralized fury |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/maker/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/maker/v1/query.proto



<a name="blackfury.maker.v1.EstimateBurnBySwapInRequest"></a>

### EstimateBurnBySwapInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateBurnBySwapInResponse"></a>

### EstimateBurnBySwapInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `burn_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `burn_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateBurnBySwapOutRequest"></a>

### EstimateBurnBySwapOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `burn_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.EstimateBurnBySwapOutResponse"></a>

### EstimateBurnBySwapOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `burn_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateBuyBackingInRequest"></a>

### EstimateBuyBackingInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateBuyBackingInResponse"></a>

### EstimateBuyBackingInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `buyback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateBuyBackingOutRequest"></a>

### EstimateBuyBackingOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.EstimateBuyBackingOutResponse"></a>

### EstimateBuyBackingOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `buyback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateMintBySwapInRequest"></a>

### EstimateMintBySwapInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_denom` | [string](#string) |  |  |
| `full_backing` | [bool](#bool) |  |  |






<a name="blackfury.maker.v1.EstimateMintBySwapInResponse"></a>

### EstimateMintBySwapInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateMintBySwapOutRequest"></a>

### EstimateMintBySwapOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `full_backing` | [bool](#bool) |  |  |






<a name="blackfury.maker.v1.EstimateMintBySwapOutResponse"></a>

### EstimateMintBySwapOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateSellBackingInRequest"></a>

### EstimateSellBackingInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.EstimateSellBackingInResponse"></a>

### EstimateSellBackingInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sellback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateSellBackingOutRequest"></a>

### EstimateSellBackingOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.EstimateSellBackingOutResponse"></a>

### EstimateSellBackingOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sellback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.QueryAllBackingPoolsRequest"></a>

### QueryAllBackingPoolsRequest







<a name="blackfury.maker.v1.QueryAllBackingPoolsResponse"></a>

### QueryAllBackingPoolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_pools` | [PoolBacking](#blackfury.maker.v1.PoolBacking) | repeated |  |






<a name="blackfury.maker.v1.QueryAllBackingRiskParamsRequest"></a>

### QueryAllBackingRiskParamsRequest







<a name="blackfury.maker.v1.QueryAllBackingRiskParamsResponse"></a>

### QueryAllBackingRiskParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `risk_params` | [BackingRiskParams](#blackfury.maker.v1.BackingRiskParams) | repeated |  |






<a name="blackfury.maker.v1.QueryAllCollateralPoolsRequest"></a>

### QueryAllCollateralPoolsRequest







<a name="blackfury.maker.v1.QueryAllCollateralPoolsResponse"></a>

### QueryAllCollateralPoolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_pools` | [PoolCollateral](#blackfury.maker.v1.PoolCollateral) | repeated |  |






<a name="blackfury.maker.v1.QueryAllCollateralRiskParamsRequest"></a>

### QueryAllCollateralRiskParamsRequest







<a name="blackfury.maker.v1.QueryAllCollateralRiskParamsResponse"></a>

### QueryAllCollateralRiskParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `risk_params` | [CollateralRiskParams](#blackfury.maker.v1.CollateralRiskParams) | repeated |  |






<a name="blackfury.maker.v1.QueryBackingPoolRequest"></a>

### QueryBackingPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.QueryBackingPoolResponse"></a>

### QueryBackingPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_pool` | [PoolBacking](#blackfury.maker.v1.PoolBacking) |  |  |






<a name="blackfury.maker.v1.QueryBackingRatioRequest"></a>

### QueryBackingRatioRequest







<a name="blackfury.maker.v1.QueryBackingRatioResponse"></a>

### QueryBackingRatioResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_ratio` | [string](#string) |  |  |
| `last_update_block` | [int64](#int64) |  |  |






<a name="blackfury.maker.v1.QueryCollateralOfAccountRequest"></a>

### QueryCollateralOfAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account` | [string](#string) |  |  |
| `collateral_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.QueryCollateralOfAccountResponse"></a>

### QueryCollateralOfAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account_collateral` | [AccountCollateral](#blackfury.maker.v1.AccountCollateral) |  |  |






<a name="blackfury.maker.v1.QueryCollateralPoolRequest"></a>

### QueryCollateralPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_denom` | [string](#string) |  |  |






<a name="blackfury.maker.v1.QueryCollateralPoolResponse"></a>

### QueryCollateralPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_pool` | [PoolCollateral](#blackfury.maker.v1.PoolCollateral) |  |  |






<a name="blackfury.maker.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.maker.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.maker.v1.Params) |  | params holds all the parameters of this module. |






<a name="blackfury.maker.v1.QueryTotalBackingRequest"></a>

### QueryTotalBackingRequest







<a name="blackfury.maker.v1.QueryTotalBackingResponse"></a>

### QueryTotalBackingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_backing` | [TotalBacking](#blackfury.maker.v1.TotalBacking) |  |  |






<a name="blackfury.maker.v1.QueryTotalCollateralRequest"></a>

### QueryTotalCollateralRequest







<a name="blackfury.maker.v1.QueryTotalCollateralResponse"></a>

### QueryTotalCollateralResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_collateral` | [TotalCollateral](#blackfury.maker.v1.TotalCollateral) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.maker.v1.Query"></a>

### Query
Query defines the maker gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `AllBackingRiskParams` | [QueryAllBackingRiskParamsRequest](#blackfury.maker.v1.QueryAllBackingRiskParamsRequest) | [QueryAllBackingRiskParamsResponse](#blackfury.maker.v1.QueryAllBackingRiskParamsResponse) | AllBackingRiskParams queries risk params of all the backing pools. | GET|/blackfury/maker/v1/all_backing_risk_params|
| `AllCollateralRiskParams` | [QueryAllCollateralRiskParamsRequest](#blackfury.maker.v1.QueryAllCollateralRiskParamsRequest) | [QueryAllCollateralRiskParamsResponse](#blackfury.maker.v1.QueryAllCollateralRiskParamsResponse) | AllCollateralRiskParams queries risk params of all the collateral pools. | GET|/blackfury/maker/v1/all_collateral_risk_params|
| `AllBackingPools` | [QueryAllBackingPoolsRequest](#blackfury.maker.v1.QueryAllBackingPoolsRequest) | [QueryAllBackingPoolsResponse](#blackfury.maker.v1.QueryAllBackingPoolsResponse) | AllBackingPools queries all the backing pools. | GET|/blackfury/maker/v1/all_backing_pools|
| `AllCollateralPools` | [QueryAllCollateralPoolsRequest](#blackfury.maker.v1.QueryAllCollateralPoolsRequest) | [QueryAllCollateralPoolsResponse](#blackfury.maker.v1.QueryAllCollateralPoolsResponse) | AllCollateralPools queries all the collateral pools. | GET|/blackfury/maker/v1/all_collateral_pools|
| `BackingPool` | [QueryBackingPoolRequest](#blackfury.maker.v1.QueryBackingPoolRequest) | [QueryBackingPoolResponse](#blackfury.maker.v1.QueryBackingPoolResponse) | BackingPool queries a backing pool. | GET|/blackfury/maker/v1/backing_pool|
| `CollateralPool` | [QueryCollateralPoolRequest](#blackfury.maker.v1.QueryCollateralPoolRequest) | [QueryCollateralPoolResponse](#blackfury.maker.v1.QueryCollateralPoolResponse) | CollateralPool queries a collateral pool. | GET|/blackfury/maker/v1/collateral_pool|
| `CollateralOfAccount` | [QueryCollateralOfAccountRequest](#blackfury.maker.v1.QueryCollateralOfAccountRequest) | [QueryCollateralOfAccountResponse](#blackfury.maker.v1.QueryCollateralOfAccountResponse) | CollateralOfAccount queries the collateral of an account. | GET|/blackfury/maker/v1/collateral_account|
| `TotalBacking` | [QueryTotalBackingRequest](#blackfury.maker.v1.QueryTotalBackingRequest) | [QueryTotalBackingResponse](#blackfury.maker.v1.QueryTotalBackingResponse) | TotalBacking queries the total backing. | GET|/blackfury/maker/v1/total_backing|
| `TotalCollateral` | [QueryTotalCollateralRequest](#blackfury.maker.v1.QueryTotalCollateralRequest) | [QueryTotalCollateralResponse](#blackfury.maker.v1.QueryTotalCollateralResponse) | TotalCollateral queries the total collateral. | GET|/blackfury/maker/v1/total_collateral|
| `BackingRatio` | [QueryBackingRatioRequest](#blackfury.maker.v1.QueryBackingRatioRequest) | [QueryBackingRatioResponse](#blackfury.maker.v1.QueryBackingRatioResponse) | BackingRatio queries the backing ratio. | GET|/blackfury/maker/v1/backing_ratio|
| `Params` | [QueryParamsRequest](#blackfury.maker.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.maker.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfury/maker/v1/params|
| `EstimateMintBySwapIn` | [EstimateMintBySwapInRequest](#blackfury.maker.v1.EstimateMintBySwapInRequest) | [EstimateMintBySwapInResponse](#blackfury.maker.v1.EstimateMintBySwapInResponse) | EstimateMintBySwapIn estimates input of minting by swap. | GET|/blackfury/maker/v1/estimate_mint_by_swap_in|
| `EstimateMintBySwapOut` | [EstimateMintBySwapOutRequest](#blackfury.maker.v1.EstimateMintBySwapOutRequest) | [EstimateMintBySwapOutResponse](#blackfury.maker.v1.EstimateMintBySwapOutResponse) | EstimateMintBySwapOut estimates output of minting by swap. | GET|/blackfury/maker/v1/estimate_mint_by_swap_out|
| `EstimateBurnBySwapIn` | [EstimateBurnBySwapInRequest](#blackfury.maker.v1.EstimateBurnBySwapInRequest) | [EstimateBurnBySwapInResponse](#blackfury.maker.v1.EstimateBurnBySwapInResponse) | EstimateBurnBySwapIn estimates input of burning by swap. | GET|/blackfury/maker/v1/estimate_burn_by_swap_in|
| `EstimateBurnBySwapOut` | [EstimateBurnBySwapOutRequest](#blackfury.maker.v1.EstimateBurnBySwapOutRequest) | [EstimateBurnBySwapOutResponse](#blackfury.maker.v1.EstimateBurnBySwapOutResponse) | EstimateBurnBySwapOut estimates output of burning by swap. | GET|/blackfury/maker/v1/estimate_burn_by_swap_out|
| `EstimateBuyBackingIn` | [EstimateBuyBackingInRequest](#blackfury.maker.v1.EstimateBuyBackingInRequest) | [EstimateBuyBackingInResponse](#blackfury.maker.v1.EstimateBuyBackingInResponse) | EstimateBuyBackingIn estimates inpput of buying backing assets. | GET|/blackfury/maker/v1/estimate_buy_backing_in|
| `EstimateBuyBackingOut` | [EstimateBuyBackingOutRequest](#blackfury.maker.v1.EstimateBuyBackingOutRequest) | [EstimateBuyBackingOutResponse](#blackfury.maker.v1.EstimateBuyBackingOutResponse) | EstimateBuyBackingOut estimates output of buying backing assets. | GET|/blackfury/maker/v1/estimate_buy_backing_out|
| `EstimateSellBackingIn` | [EstimateSellBackingInRequest](#blackfury.maker.v1.EstimateSellBackingInRequest) | [EstimateSellBackingInResponse](#blackfury.maker.v1.EstimateSellBackingInResponse) | EstimateSellBackingIn estimates input of selling backing assets. | GET|/blackfury/maker/v1/estimate_sell_backing_in|
| `EstimateSellBackingOut` | [EstimateSellBackingOutRequest](#blackfury.maker.v1.EstimateSellBackingOutRequest) | [EstimateSellBackingOutResponse](#blackfury.maker.v1.EstimateSellBackingOutResponse) | EstimateSellBackingOut estimates output of selling backing assets. | GET|/blackfury/maker/v1/estimate_sell_backing_out|

 <!-- end services -->



<a name="blackfury/maker/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/maker/v1/tx.proto



<a name="blackfury.maker.v1.MsgBurnByCollateral"></a>

### MsgBurnByCollateral
MsgBurnByCollateral represents a message to burn Black stablecoins by unlocking
collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_denom` | [string](#string) |  |  |
| `repay_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgBurnByCollateralResponse"></a>

### MsgBurnByCollateralResponse
MsgBurnByCollateralResponse defines the Msg/BurnByCollateral response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `repay_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgBurnBySwap"></a>

### MsgBurnBySwap
MsgBurnBySwap represents a message to burn Black stablecoins by swapping.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `burn_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_out_min` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out_min` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgBurnBySwapResponse"></a>

### MsgBurnBySwapResponse
MsgBurnBySwapResponse defines the Msg/BurnBySwap response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `burn_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgBuyBacking"></a>

### MsgBuyBacking
MsgBuyBacking represents a message to buy strong-backing assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `backing_out_min` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgBuyBackingResponse"></a>

### MsgBuyBackingResponse
MsgBuyBackingResponse defines the Msg/BuyBacking response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `buyback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgDepositCollateral"></a>

### MsgDepositCollateral
MsgDepositCollateral represents a message to deposit collateral assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `collateral_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgDepositCollateralResponse"></a>

### MsgDepositCollateralResponse
MsgDepositCollateralResponse defines the Msg/DepositCollateral response type.






<a name="blackfury.maker.v1.MsgLiquidateCollateral"></a>

### MsgLiquidateCollateral
MsgLiquidateCollateral represents a message to liquidates collateral assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `debtor` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `repay_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgLiquidateCollateralResponse"></a>

### MsgLiquidateCollateralResponse
MsgLiquidateCollateralResponse defines the Msg/LiquidateCollateral response
type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `repay_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgMintByCollateral"></a>

### MsgMintByCollateral
MsgMintByCollateral represents a message to mint Black stablecoins by locking
collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `collateral_denom` | [string](#string) |  |  |
| `mint_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgMintByCollateralResponse"></a>

### MsgMintByCollateralResponse
MsgMintByCollateralResponse defines the Msg/MintByCollateral response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgMintBySwap"></a>

### MsgMintBySwap
MsgMintBySwap represents a message to mint Black stablecoins by swapping.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `backing_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in_max` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_out_min` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `full_backing` | [bool](#bool) |  |  |






<a name="blackfury.maker.v1.MsgMintBySwapResponse"></a>

### MsgMintBySwapResponse
MsgMintBySwapResponse defines the Msg/MintBySwap response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mint_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgRedeemCollateral"></a>

### MsgRedeemCollateral
MsgRedeemCollateral represents a message to redeem collateral assets and
collateralized Fury coins.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `collateral_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgRedeemCollateralResponse"></a>

### MsgRedeemCollateralResponse
MsgRedeemCollateralResponse defines the Msg/RedeemCollateral response type.






<a name="blackfury.maker.v1.MsgSellBacking"></a>

### MsgSellBacking
MsgSellBacking represents a message to sell strong-backing
assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `backing_in` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fury_out_min` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.maker.v1.MsgSellBackingResponse"></a>

### MsgSellBackingResponse
MsgSellBackingResponse defines the Msg/SellBacking response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fury_out` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `reback_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.maker.v1.Msg"></a>

### Msg
Msg defines the maker Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintBySwap` | [MsgMintBySwap](#blackfury.maker.v1.MsgMintBySwap) | [MsgMintBySwapResponse](#blackfury.maker.v1.MsgMintBySwapResponse) | MintBySwap mints Black stablecoins by swapping in strong-backing assets and Fury coins. | GET|/blackfury/maker/v1/tx/mint_by_swap|
| `BurnBySwap` | [MsgBurnBySwap](#blackfury.maker.v1.MsgBurnBySwap) | [MsgBurnBySwapResponse](#blackfury.maker.v1.MsgBurnBySwapResponse) | BurnBySwap burns Black stablecoins by swapping out strong-backing assets and Fury coins. | GET|/blackfury/maker/v1/tx/burn_by_swap|
| `BuyBacking` | [MsgBuyBacking](#blackfury.maker.v1.MsgBuyBacking) | [MsgBuyBackingResponse](#blackfury.maker.v1.MsgBuyBackingResponse) | BuyBacking buys strong-backing assets by spending Fury coins. | GET|/blackfury/maker/v1/tx/buy_backing|
| `SellBacking` | [MsgSellBacking](#blackfury.maker.v1.MsgSellBacking) | [MsgSellBackingResponse](#blackfury.maker.v1.MsgSellBackingResponse) | SellBacking sells strong-backing assets by earning Fury coins. | GET|/blackfury/maker/v1/tx/sell_backing|
| `MintByCollateral` | [MsgMintByCollateral](#blackfury.maker.v1.MsgMintByCollateral) | [MsgMintByCollateralResponse](#blackfury.maker.v1.MsgMintByCollateralResponse) | MintByCollateral mints Black stablecoins by locking collateral assets and spending Fury coins. | GET|/blackfury/maker/v1/tx/mint_by_collateral|
| `BurnByCollateral` | [MsgBurnByCollateral](#blackfury.maker.v1.MsgBurnByCollateral) | [MsgBurnByCollateralResponse](#blackfury.maker.v1.MsgBurnByCollateralResponse) | BurnByCollateral burns Black stablecoins by unlocking collateral assets and earning Fury coins. | GET|/blackfury/maker/v1/tx/burn_by_collateral|
| `DepositCollateral` | [MsgDepositCollateral](#blackfury.maker.v1.MsgDepositCollateral) | [MsgDepositCollateralResponse](#blackfury.maker.v1.MsgDepositCollateralResponse) | DepositCollateral deposits collateral assets. | GET|/blackfury/maker/v1/tx/deposit_collateral|
| `RedeemCollateral` | [MsgRedeemCollateral](#blackfury.maker.v1.MsgRedeemCollateral) | [MsgRedeemCollateralResponse](#blackfury.maker.v1.MsgRedeemCollateralResponse) | RedeemCollateral redeems collateral assets and collateralized Fury coins. | GET|/blackfury/maker/v1/tx/redeem_collateral|
| `LiquidateCollateral` | [MsgLiquidateCollateral](#blackfury.maker.v1.MsgLiquidateCollateral) | [MsgLiquidateCollateralResponse](#blackfury.maker.v1.MsgLiquidateCollateralResponse) | LiquidateCollateral liquidates collateral assets which is undercollateralized. | GET|/blackfury/maker/v1/tx/liquidate_collateral|

 <!-- end services -->



<a name="blackfury/oracle/v1/oracle.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/oracle/v1/oracle.proto



<a name="blackfury.oracle.v1.AggregateExchangeRatePrevote"></a>

### AggregateExchangeRatePrevote
AggregateExchangeRatePrevote represents the aggregate prevoting on the
ExchangeRateVote. The purpose of aggregate prevoting is to hide vote exchange
rates with hash which is formatted as hex string in SHA256("{salt}:{exchange
rate}{denom},...,{exchange rate}{denom}:{voter}")


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [string](#string) |  |  |
| `voter` | [string](#string) |  |  |
| `submit_block` | [uint64](#uint64) |  |  |






<a name="blackfury.oracle.v1.AggregateExchangeRateVote"></a>

### AggregateExchangeRateVote
AggregateExchangeRateVote represents the voting on
the exchange rates of various assets denominated in uUSD.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exchange_rate_tuples` | [ExchangeRateTuple](#blackfury.oracle.v1.ExchangeRateTuple) | repeated |  |
| `voter` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.ExchangeRateTuple"></a>

### ExchangeRateTuple
ExchangeRateTuple stores interpreted exchange rates data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `exchange_rate` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.Params"></a>

### Params
Params defines the parameters for the oracle module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote_period` | [uint64](#uint64) |  |  |
| `vote_threshold` | [string](#string) |  |  |
| `reward_band` | [string](#string) |  |  |
| `reward_distribution_window` | [uint64](#uint64) |  |  |
| `slash_fraction` | [string](#string) |  |  |
| `slash_window` | [uint64](#uint64) |  |  |
| `min_valid_per_window` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.RegisterTargetProposal"></a>

### RegisterTargetProposal
RegisterTargetProposal is a gov Content type to register eligible
target asset which will be price quoted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `target_params` | [TargetParams](#blackfury.oracle.v1.TargetParams) |  | target params |






<a name="blackfury.oracle.v1.TargetParams"></a>

### TargetParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | coin denom |
| `source` | [TargetSource](#blackfury.oracle.v1.TargetSource) |  | quotation source |
| `source_dex_contract` | [string](#string) |  | quotation source DEX contract address |





 <!-- end messages -->


<a name="blackfury.oracle.v1.TargetSource"></a>

### TargetSource
TargetSource enumerates the quotation source of a target asset.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TARGET_SOURCE_UNSPECIFIED | 0 | TARGET_SOURCE_UNSPECIFIED defines an invalid/undefined target source. |
| TARGET_SOURCE_VALIDATORS | 1 | TARGET_SOURCE_VALIDATORS target quotation source is from validators. |
| TARGET_SOURCE_DEX | 2 | TARGET_SOURCE_DEX target quotation source is from on-chain DEX. |
| TARGET_SOURCE_INTERCHAIN_DEX | 3 | TARGET_SOURCE_INTERCHAIN_DEX target quotation source is from inter-chain DEX. |
| TARGET_SOURCE_INTERCHAIN_ORACLE | 4 | TARGET_SOURCE_INTERCHAIN_ORACLE target quotation source is from inter-chain oracle. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/oracle/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/oracle/v1/genesis.proto



<a name="blackfury.oracle.v1.FeederDelegation"></a>

### FeederDelegation
FeederDelegation is the address for where oracle feeder authority are
delegated to. By default this struct is only used at genesis to feed in
default feeder addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `feeder_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.GenesisState"></a>

### GenesisState
GenesisState defines the oracle module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.oracle.v1.Params) |  |  |
| `feeder_delegations` | [FeederDelegation](#blackfury.oracle.v1.FeederDelegation) | repeated |  |
| `exchange_rates` | [ExchangeRateTuple](#blackfury.oracle.v1.ExchangeRateTuple) | repeated |  |
| `miss_counters` | [MissCounter](#blackfury.oracle.v1.MissCounter) | repeated |  |
| `aggregate_exchange_rate_prevotes` | [AggregateExchangeRatePrevote](#blackfury.oracle.v1.AggregateExchangeRatePrevote) | repeated |  |
| `aggregate_exchange_rate_votes` | [AggregateExchangeRateVote](#blackfury.oracle.v1.AggregateExchangeRateVote) | repeated |  |






<a name="blackfury.oracle.v1.MissCounter"></a>

### MissCounter
MissCounter defines an miss counter and validator address pair used in
oracle module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |
| `miss_counter` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/oracle/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/oracle/v1/query.proto



<a name="blackfury.oracle.v1.QueryActivesRequest"></a>

### QueryActivesRequest
QueryActivesRequest is the request type for the Query/Actives RPC method.






<a name="blackfury.oracle.v1.QueryActivesResponse"></a>

### QueryActivesResponse
QueryActivesResponse is response type for the
Query/Actives RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `actives` | [string](#string) | repeated | actives defines a list of the denomination which oracle prices aggreed upon. |






<a name="blackfury.oracle.v1.QueryAggregatePrevoteRequest"></a>

### QueryAggregatePrevoteRequest
QueryAggregatePrevoteRequest is the request type for the
Query/AggregatePrevote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator defines the validator address to query for. |






<a name="blackfury.oracle.v1.QueryAggregatePrevoteResponse"></a>

### QueryAggregatePrevoteResponse
QueryAggregatePrevoteResponse is response type for the
Query/AggregatePrevote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `aggregate_prevote` | [AggregateExchangeRatePrevote](#blackfury.oracle.v1.AggregateExchangeRatePrevote) |  | aggregate_prevote defines oracle aggregate prevote submitted by a validator in the current vote period. |






<a name="blackfury.oracle.v1.QueryAggregatePrevotesRequest"></a>

### QueryAggregatePrevotesRequest
QueryAggregatePrevotesRequest is the request type for the
Query/AggregatePrevotes RPC method.






<a name="blackfury.oracle.v1.QueryAggregatePrevotesResponse"></a>

### QueryAggregatePrevotesResponse
QueryAggregatePrevotesResponse is response type for the
Query/AggregatePrevotes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `aggregate_prevotes` | [AggregateExchangeRatePrevote](#blackfury.oracle.v1.AggregateExchangeRatePrevote) | repeated | aggregate_prevotes defines all oracle aggregate prevotes submitted in the current vote period. |






<a name="blackfury.oracle.v1.QueryAggregateVoteRequest"></a>

### QueryAggregateVoteRequest
QueryAggregateVoteRequest is the request type for the Query/AggregateVote RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator defines the validator address to query for. |






<a name="blackfury.oracle.v1.QueryAggregateVoteResponse"></a>

### QueryAggregateVoteResponse
QueryAggregateVoteResponse is response type for the
Query/AggregateVote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `aggregate_vote` | [AggregateExchangeRateVote](#blackfury.oracle.v1.AggregateExchangeRateVote) |  | aggregate_vote defines oracle aggregate vote submitted by a validator in the current vote period. |






<a name="blackfury.oracle.v1.QueryAggregateVotesRequest"></a>

### QueryAggregateVotesRequest
QueryAggregateVotesRequest is the request type for the Query/AggregateVotes
RPC method.






<a name="blackfury.oracle.v1.QueryAggregateVotesResponse"></a>

### QueryAggregateVotesResponse
QueryAggregateVotesResponse is response type for the
Query/AggregateVotes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `aggregate_votes` | [AggregateExchangeRateVote](#blackfury.oracle.v1.AggregateExchangeRateVote) | repeated | aggregate_votes defines all oracle aggregate votes submitted in the current vote period. |






<a name="blackfury.oracle.v1.QueryExchangeRateRequest"></a>

### QueryExchangeRateRequest
QueryExchangeRateRequest is the request type for the Query/ExchangeRate RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom defines the denomination to query for. |






<a name="blackfury.oracle.v1.QueryExchangeRateResponse"></a>

### QueryExchangeRateResponse
QueryExchangeRateResponse is response type for the
Query/ExchangeRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exchange_rate` | [string](#string) |  | exchange_rate defines the exchange rate of the denom asset denominated in uUSD. |






<a name="blackfury.oracle.v1.QueryExchangeRatesRequest"></a>

### QueryExchangeRatesRequest
QueryExchangeRatesRequest is the request type for the Query/ExchangeRates RPC
method.






<a name="blackfury.oracle.v1.QueryExchangeRatesResponse"></a>

### QueryExchangeRatesResponse
QueryExchangeRatesResponse is response type for the
Query/ExchangeRates RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exchange_rates` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | exchange_rates defines a list of the exchange rate for all whitelisted denoms. |






<a name="blackfury.oracle.v1.QueryFeederDelegationRequest"></a>

### QueryFeederDelegationRequest
QueryFeederDelegationRequest is the request type for the
Query/FeederDelegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator defines the validator address to query for. |






<a name="blackfury.oracle.v1.QueryFeederDelegationResponse"></a>

### QueryFeederDelegationResponse
QueryFeederDelegationResponse is response type for the
Query/FeederDelegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `feeder_addr` | [string](#string) |  | feeder_addr defines the feeder delegation of a validator. |






<a name="blackfury.oracle.v1.QueryMissCounterRequest"></a>

### QueryMissCounterRequest
QueryMissCounterRequest is the request type for the Query/MissCounter RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator defines the validator address to query for. |






<a name="blackfury.oracle.v1.QueryMissCounterResponse"></a>

### QueryMissCounterResponse
QueryMissCounterResponse is response type for the
Query/MissCounter RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `miss_counter` | [uint64](#uint64) |  | miss_counter defines the oracle miss counter of a validator. |






<a name="blackfury.oracle.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.oracle.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.oracle.v1.Params) |  | params holds all the parameters of this module. |






<a name="blackfury.oracle.v1.QueryTargetsRequest"></a>

### QueryTargetsRequest
QueryTargetsRequest is the request type for the Query/Targets RPC
method.






<a name="blackfury.oracle.v1.QueryTargetsResponse"></a>

### QueryTargetsResponse
QueryTargetsResponse is response type for the
Query/Targets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `targets` | [string](#string) | repeated | targets defines a list of the denomination which will be fed with price quotation (including voting targets). |






<a name="blackfury.oracle.v1.QueryVoteTargetsRequest"></a>

### QueryVoteTargetsRequest
QueryVoteTargetsRequest is the request type for the Query/VoteTargets RPC
method.






<a name="blackfury.oracle.v1.QueryVoteTargetsResponse"></a>

### QueryVoteTargetsResponse
QueryVoteTargetsResponse is response type for the
Query/VoteTargets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote_targets` | [string](#string) | repeated | vote_targets defines a list of the denomination in which everyone should vote in the current vote period. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.oracle.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ExchangeRate` | [QueryExchangeRateRequest](#blackfury.oracle.v1.QueryExchangeRateRequest) | [QueryExchangeRateResponse](#blackfury.oracle.v1.QueryExchangeRateResponse) | ExchangeRate returns exchange rate of a denom. | GET|/blackfury/oracle/v1/denoms/{denom}/exchange_rate|
| `ExchangeRates` | [QueryExchangeRatesRequest](#blackfury.oracle.v1.QueryExchangeRatesRequest) | [QueryExchangeRatesResponse](#blackfury.oracle.v1.QueryExchangeRatesResponse) | ExchangeRates returns exchange rates of all denoms. | GET|/blackfury/oracle/v1/denoms/exchange_rates|
| `Actives` | [QueryActivesRequest](#blackfury.oracle.v1.QueryActivesRequest) | [QueryActivesResponse](#blackfury.oracle.v1.QueryActivesResponse) | Actives returns all active denoms. | GET|/blackfury/oracle/v1/denoms/actives|
| `VoteTargets` | [QueryVoteTargetsRequest](#blackfury.oracle.v1.QueryVoteTargetsRequest) | [QueryVoteTargetsResponse](#blackfury.oracle.v1.QueryVoteTargetsResponse) | VoteTargets returns all vote target denoms. | GET|/blackfury/oracle/v1/denoms/vote_targets|
| `Targets` | [QueryTargetsRequest](#blackfury.oracle.v1.QueryTargetsRequest) | [QueryTargetsResponse](#blackfury.oracle.v1.QueryTargetsResponse) | Targets returns all target denoms (including vote targets). | GET|/blackfury/oracle/v1/denoms/targets|
| `FeederDelegation` | [QueryFeederDelegationRequest](#blackfury.oracle.v1.QueryFeederDelegationRequest) | [QueryFeederDelegationResponse](#blackfury.oracle.v1.QueryFeederDelegationResponse) | FeederDelegation returns feeder delegation of a validator. | GET|/blackfury/oracle/v1/validators/{validator_addr}/feeder|
| `MissCounter` | [QueryMissCounterRequest](#blackfury.oracle.v1.QueryMissCounterRequest) | [QueryMissCounterResponse](#blackfury.oracle.v1.QueryMissCounterResponse) | MissCounter returns oracle miss counter of a validator. | GET|/blackfury/oracle/v1/validators/{validator_addr}/miss|
| `AggregatePrevote` | [QueryAggregatePrevoteRequest](#blackfury.oracle.v1.QueryAggregatePrevoteRequest) | [QueryAggregatePrevoteResponse](#blackfury.oracle.v1.QueryAggregatePrevoteResponse) | AggregatePrevote returns an aggregate prevote of a validator. | GET|/blackfury/oracle/v1/validators/{validator_addr}/aggregate_prevote|
| `AggregatePrevotes` | [QueryAggregatePrevotesRequest](#blackfury.oracle.v1.QueryAggregatePrevotesRequest) | [QueryAggregatePrevotesResponse](#blackfury.oracle.v1.QueryAggregatePrevotesResponse) | AggregatePrevotes returns aggregate prevotes of all validators. | GET|/blackfury/oracle/v1/validators/aggregate_prevotes|
| `AggregateVote` | [QueryAggregateVoteRequest](#blackfury.oracle.v1.QueryAggregateVoteRequest) | [QueryAggregateVoteResponse](#blackfury.oracle.v1.QueryAggregateVoteResponse) | AggregateVote returns an aggregate vote of a validator. | GET|/blackfury/oracle/v1/valdiators/{validator_addr}/aggregate_vote|
| `AggregateVotes` | [QueryAggregateVotesRequest](#blackfury.oracle.v1.QueryAggregateVotesRequest) | [QueryAggregateVotesResponse](#blackfury.oracle.v1.QueryAggregateVotesResponse) | AggregateVotes returns aggregate votes of all validators. | GET|/blackfury/oracle/v1/validators/aggregate_votes|
| `Params` | [QueryParamsRequest](#blackfury.oracle.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.oracle.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfuryzone/blackfury/oracle/params|

 <!-- end services -->



<a name="blackfury/oracle/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/oracle/v1/tx.proto



<a name="blackfury.oracle.v1.MsgAggregateExchangeRatePrevote"></a>

### MsgAggregateExchangeRatePrevote
MsgAggregateExchangeRatePrevote defines a message to submit
aggregate exchange rate prevote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [string](#string) |  |  |
| `feeder` | [string](#string) |  |  |
| `validator` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.MsgAggregateExchangeRatePrevoteResponse"></a>

### MsgAggregateExchangeRatePrevoteResponse
MsgAggregateExchangeRatePrevoteResponse defines the
MsgAggregateExchangeRatePrevote response type.






<a name="blackfury.oracle.v1.MsgAggregateExchangeRateVote"></a>

### MsgAggregateExchangeRateVote
MsgAggregateExchangeRateVote defines a message to submit
aggregate exchange rate vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `salt` | [string](#string) |  |  |
| `exchange_rates` | [string](#string) |  |  |
| `feeder` | [string](#string) |  |  |
| `validator` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.MsgAggregateExchangeRateVoteResponse"></a>

### MsgAggregateExchangeRateVoteResponse
MsgAggregateExchangeRateVoteResponse defines the MsgAggregateExchangeRateVote
response type.






<a name="blackfury.oracle.v1.MsgDelegateFeedConsent"></a>

### MsgDelegateFeedConsent
MsgDelegateFeedConsent defines a message to
delegate oracle voting rights to another address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator` | [string](#string) |  |  |
| `delegate` | [string](#string) |  |  |






<a name="blackfury.oracle.v1.MsgDelegateFeedConsentResponse"></a>

### MsgDelegateFeedConsentResponse
MsgDelegateFeedConsentResponse defines the MsgDelegateFeedConsent response
type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.oracle.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `AggregateExchangeRatePrevote` | [MsgAggregateExchangeRatePrevote](#blackfury.oracle.v1.MsgAggregateExchangeRatePrevote) | [MsgAggregateExchangeRatePrevoteResponse](#blackfury.oracle.v1.MsgAggregateExchangeRatePrevoteResponse) | AggregateExchangeRatePrevote submits aggregate exchange rate prevote. | GET|/blackfury/oracle/v1/tx/aggregate_exchange_rate_prevote|
| `AggregateExchangeRateVote` | [MsgAggregateExchangeRateVote](#blackfury.oracle.v1.MsgAggregateExchangeRateVote) | [MsgAggregateExchangeRateVoteResponse](#blackfury.oracle.v1.MsgAggregateExchangeRateVoteResponse) | AggregateExchangeRateVote submits aggregate exchange rate vote. | GET|/blackfury/oracle/v1/tx/aggregate_exchange_rate_vote|
| `DelegateFeedConsent` | [MsgDelegateFeedConsent](#blackfury.oracle.v1.MsgDelegateFeedConsent) | [MsgDelegateFeedConsentResponse](#blackfury.oracle.v1.MsgDelegateFeedConsentResponse) | DelegateFeedConsent sets the feeder delegation. | GET|/blackfury/oracle/v1/tx/delegate_feed_consent|

 <!-- end services -->



<a name="blackfury/staking/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/staking/v1/query.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/staking/v1/staking.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/staking/v1/staking.proto



<a name="blackfury.staking.v1.VeDelegation"></a>

### VeDelegation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `ve_shares` | [VeShares](#blackfury.staking.v1.VeShares) | repeated |  |






<a name="blackfury.staking.v1.VeRedelegation"></a>

### VeRedelegation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |
| `entries` | [VeRedelegationEntry](#blackfury.staking.v1.VeRedelegationEntry) | repeated |  |






<a name="blackfury.staking.v1.VeRedelegationEntry"></a>

### VeRedelegationEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_shares` | [VeRedelegationEntryShares](#blackfury.staking.v1.VeRedelegationEntryShares) | repeated |  |






<a name="blackfury.staking.v1.VeRedelegationEntryShares"></a>

### VeRedelegationEntryShares



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [uint64](#uint64) |  |  |
| `initial_balance` | [string](#string) |  |  |
| `shares_dst` | [string](#string) |  |  |






<a name="blackfury.staking.v1.VeShares"></a>

### VeShares



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [uint64](#uint64) |  |  |
| `tokens_may_unsettled` | [string](#string) |  |  |
| `shares` | [string](#string) |  |  |






<a name="blackfury.staking.v1.VeTokens"></a>

### VeTokens



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [uint64](#uint64) |  |  |
| `tokens` | [string](#string) |  |  |






<a name="blackfury.staking.v1.VeUnbondingDelegation"></a>

### VeUnbondingDelegation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `entries` | [VeUnbondingDelegationEntry](#blackfury.staking.v1.VeUnbondingDelegationEntry) | repeated |  |






<a name="blackfury.staking.v1.VeUnbondingDelegationEntry"></a>

### VeUnbondingDelegationEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_balances` | [VeUnbondingDelegationEntryBalances](#blackfury.staking.v1.VeUnbondingDelegationEntryBalances) | repeated |  |






<a name="blackfury.staking.v1.VeUnbondingDelegationEntryBalances"></a>

### VeUnbondingDelegationEntryBalances



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [uint64](#uint64) |  |  |
| `initial_balance` | [string](#string) |  |  |
| `balance` | [string](#string) |  |  |






<a name="blackfury.staking.v1.VeValidator"></a>

### VeValidator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  |  |
| `ve_delegator_shares` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/staking/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/staking/v1/tx.proto



<a name="blackfury.staking.v1.MsgVeDelegate"></a>

### MsgVeDelegate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.staking.v1.MsgVeDelegateResponse"></a>

### MsgVeDelegateResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.staking.v1.Msg"></a>

### Msg
Msg defines the staking Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `VeDelegate` | [MsgVeDelegate](#blackfury.staking.v1.MsgVeDelegate) | [MsgVeDelegateResponse](#blackfury.staking.v1.MsgVeDelegateResponse) | VeDelegate defines a method for performing a delegation of ve-locked coins from a delegator to a validator. | GET|/blackfury/staking/v1/tx/ve_delegate|

 <!-- end services -->



<a name="blackfury/ve/v1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/ve/v1/event.proto



<a name="blackfury.ve.v1.EventCreate"></a>

### EventCreate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `receiver` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `unlock_time` | [uint64](#uint64) |  |  |






<a name="blackfury.ve.v1.EventDeposit"></a>

### EventDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="blackfury.ve.v1.EventExtendTime"></a>

### EventExtendTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `unlock_time` | [uint64](#uint64) |  |  |






<a name="blackfury.ve.v1.EventMerge"></a>

### EventMerge



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `from_ve_id` | [string](#string) |  |  |
| `to_ve_id` | [string](#string) |  |  |






<a name="blackfury.ve.v1.EventWithdraw"></a>

### EventWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/ve/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/ve/v1/genesis.proto



<a name="blackfury.ve.v1.GenesisState"></a>

### GenesisState
GenesisState defines the ve module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.ve.v1.Params) |  |  |






<a name="blackfury.ve.v1.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `lock_denom` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/ve/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/ve/v1/query.proto



<a name="blackfury.ve.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.ve.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.ve.v1.Params) |  | params holds all the parameters of this module. |






<a name="blackfury.ve.v1.QueryTotalVotingPowerRequest"></a>

### QueryTotalVotingPowerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `at_time` | [uint64](#uint64) |  |  |
| `at_block` | [int64](#int64) |  |  |






<a name="blackfury.ve.v1.QueryTotalVotingPowerResponse"></a>

### QueryTotalVotingPowerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `power` | [string](#string) |  |  |






<a name="blackfury.ve.v1.QueryVeNftRequest"></a>

### QueryVeNftRequest
QueryVeNftRequest is the request type for the Query/VeNft RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |






<a name="blackfury.ve.v1.QueryVeNftResponse"></a>

### QueryVeNftResponse
QueryVeNftResponse is the response type for the Query/VeNft RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft` | [cosmos.nft.v1beta1.NFT](#cosmos.nft.v1beta1.NFT) |  |  |






<a name="blackfury.ve.v1.QueryVeNftsRequest"></a>

### QueryVeNftsRequest
QueryVeNftsRequest is the request type for the Query/VeNfts RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="blackfury.ve.v1.QueryVeNftsResponse"></a>

### QueryVeNftsResponse
QueryVeNftsResponse is the response type for the Query/VeNfts RPC methods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [cosmos.nft.v1beta1.NFT](#cosmos.nft.v1beta1.NFT) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="blackfury.ve.v1.QueryVotingPowerRequest"></a>

### QueryVotingPowerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [string](#string) |  |  |
| `at_time` | [uint64](#uint64) |  |  |
| `at_block` | [int64](#int64) |  |  |






<a name="blackfury.ve.v1.QueryVotingPowerResponse"></a>

### QueryVotingPowerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `power` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.ve.v1.Query"></a>

### Query
Query defines the ve gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TotalVotingPower` | [QueryTotalVotingPowerRequest](#blackfury.ve.v1.QueryTotalVotingPowerRequest) | [QueryTotalVotingPowerResponse](#blackfury.ve.v1.QueryTotalVotingPowerResponse) | TotalVotingPower queries the total voting power. | GET|/blackfury/ve/v1/total_voting_power|
| `VotingPower` | [QueryVotingPowerRequest](#blackfury.ve.v1.QueryVotingPowerRequest) | [QueryVotingPowerResponse](#blackfury.ve.v1.QueryVotingPowerResponse) | VotingPower queries the voting power of a veNFT. | GET|/blackfury/ve/v1/voting_power/{ve_id}|
| `VeNfts` | [QueryVeNftsRequest](#blackfury.ve.v1.QueryVeNftsRequest) | [QueryVeNftsResponse](#blackfury.ve.v1.QueryVeNftsResponse) | VeNfts queries all veNFTs of a given owner. | GET|/blackfury/ve/v1/venfts|
| `VeNft` | [QueryVeNftRequest](#blackfury.ve.v1.QueryVeNftRequest) | [QueryVeNftResponse](#blackfury.ve.v1.QueryVeNftResponse) | VeNft queries an veNFT based on its id. | GET|/blackfury/ve/v1/venfts/{id}|
| `Params` | [QueryParamsRequest](#blackfury.ve.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.ve.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfury/ve/v1/params|

 <!-- end services -->



<a name="blackfury/ve/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/ve/v1/tx.proto



<a name="blackfury.ve.v1.MsgCreate"></a>

### MsgCreate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lock_duration` | [uint64](#uint64) |  |  |






<a name="blackfury.ve.v1.MsgCreateResponse"></a>

### MsgCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ve_id` | [string](#string) |  |  |
| `unlock_time` | [uint64](#uint64) |  |  |






<a name="blackfury.ve.v1.MsgDeposit"></a>

### MsgDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount to deposit, must be greater than 0 |






<a name="blackfury.ve.v1.MsgDepositResponse"></a>

### MsgDepositResponse







<a name="blackfury.ve.v1.MsgExtendTime"></a>

### MsgExtendTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |
| `lock_duration` | [uint64](#uint64) |  | Locking duration, must be greater than current locking duration |






<a name="blackfury.ve.v1.MsgExtendTimeResponse"></a>

### MsgExtendTimeResponse







<a name="blackfury.ve.v1.MsgMerge"></a>

### MsgMerge



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `from_ve_id` | [string](#string) |  |  |
| `to_ve_id` | [string](#string) |  |  |






<a name="blackfury.ve.v1.MsgMergeResponse"></a>

### MsgMergeResponse







<a name="blackfury.ve.v1.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `ve_id` | [string](#string) |  |  |






<a name="blackfury.ve.v1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.ve.v1.Msg"></a>

### Msg
Msg defines the ve Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Create` | [MsgCreate](#blackfury.ve.v1.MsgCreate) | [MsgCreateResponse](#blackfury.ve.v1.MsgCreateResponse) | Create creates a veNFT. | GET|/blackfury/ve/v1/tx/create|
| `Deposit` | [MsgDeposit](#blackfury.ve.v1.MsgDeposit) | [MsgDepositResponse](#blackfury.ve.v1.MsgDepositResponse) | Deposit deposits some coin amount for a veNFT. | GET|/blackfury/ve/v1/tx/deposit|
| `ExtendTime` | [MsgExtendTime](#blackfury.ve.v1.MsgExtendTime) | [MsgExtendTimeResponse](#blackfury.ve.v1.MsgExtendTimeResponse) | ExtendTime extends locking duration for a veNFT. | GET|/blackfury/ve/v1/tx/extend_time|
| `Merge` | [MsgMerge](#blackfury.ve.v1.MsgMerge) | [MsgMergeResponse](#blackfury.ve.v1.MsgMergeResponse) | Merge merges a veNFT (burn it) to another veNFT. | GET|/blackfury/ve/v1/tx/merge|
| `Withdraw` | [MsgWithdraw](#blackfury.ve.v1.MsgWithdraw) | [MsgWithdrawResponse](#blackfury.ve.v1.MsgWithdrawResponse) | Withdraw withdraws all coin amount of a veNFT. | GET|/blackfury/ve/v1/tx/withdraw|

 <!-- end services -->



<a name="blackfury/ve/v1/ve.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/ve/v1/ve.proto



<a name="blackfury.ve.v1.Checkpoint"></a>

### Checkpoint
Checkpoint defines a checkpoint of voting power.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bias` | [string](#string) |  | voting power at checkpoint |
| `slope` | [string](#string) |  | weight decay slope so voting power at time t: bias - slope * (t - timestamp) |
| `timestamp` | [uint64](#uint64) |  | unix timestamp at checkpoint |
| `block` | [int64](#int64) |  | block height at checkpoint |






<a name="blackfury.ve.v1.LockedBalance"></a>

### LockedBalance
LockedBalance represents locked amount and unlock time of a ve.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [string](#string) |  | locked amount |
| `end` | [uint64](#uint64) |  | unlocking unix time |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/vesting/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/vesting/v1/genesis.proto



<a name="blackfury.vesting.v1.AllocationAddresses"></a>

### AllocationAddresses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `team_vesting_addr` | [string](#string) |  |  |
| `strategic_reserve_custodian_addr` | [string](#string) |  |  |






<a name="blackfury.vesting.v1.AllocationAmounts"></a>

### AllocationAmounts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_amount` | [string](#string) |  |  |
| `airdrop_amount` | [string](#string) |  |  |
| `ve_vesting_amount` | [string](#string) |  |  |
| `staking_reward_amount` | [string](#string) |  |  |
| `community_pool_amount` | [string](#string) |  |  |
| `strategic_reserve_amount` | [string](#string) |  |  |
| `team_vesting_amount` | [string](#string) |  |  |






<a name="blackfury.vesting.v1.GenesisState"></a>

### GenesisState
GenesisState defines the vesting module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.vesting.v1.Params) |  |  |
| `allocation_addresses` | [AllocationAddresses](#blackfury.vesting.v1.AllocationAddresses) |  |  |






<a name="blackfury.vesting.v1.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allocation` | [AllocationAmounts](#blackfury.vesting.v1.AllocationAmounts) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/vesting/v1/vesting.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/vesting/v1/vesting.proto



<a name="blackfury.vesting.v1.Airdrop"></a>

### Airdrop



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `target_addr` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/vesting/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/vesting/v1/query.proto



<a name="blackfury.vesting.v1.QueryAirdropRequest"></a>

### QueryAirdropRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `target_addr` | [string](#string) |  |  |
| `completed` | [bool](#bool) |  |  |






<a name="blackfury.vesting.v1.QueryAirdropResponse"></a>

### QueryAirdropResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `airdrop` | [Airdrop](#blackfury.vesting.v1.Airdrop) |  |  |






<a name="blackfury.vesting.v1.QueryAirdropsRequest"></a>

### QueryAirdropsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completed` | [bool](#bool) |  | pagination defines an optional pagination for the request. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="blackfury.vesting.v1.QueryAirdropsResponse"></a>

### QueryAirdropsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `airdrops` | [Airdrop](#blackfury.vesting.v1.Airdrop) | repeated | airdrops contains all the queried airdrops. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="blackfury.vesting.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.vesting.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.vesting.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.vesting.v1.Query"></a>

### Query
Query defines the vesting gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Airdrops` | [QueryAirdropsRequest](#blackfury.vesting.v1.QueryAirdropsRequest) | [QueryAirdropsResponse](#blackfury.vesting.v1.QueryAirdropsResponse) | Airdrops queries airdrop targets. | GET|/blackfury/vesting/v1/airdrops|
| `Airdrop` | [QueryAirdropRequest](#blackfury.vesting.v1.QueryAirdropRequest) | [QueryAirdropResponse](#blackfury.vesting.v1.QueryAirdropResponse) | Airdrops queries airdrop target for given address. | GET|/blackfury/vesting/v1/airdrops/{target_addr}|
| `Params` | [QueryParamsRequest](#blackfury.vesting.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.vesting.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfury/vesting/v1/params|

 <!-- end services -->



<a name="blackfury/vesting/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/vesting/v1/tx.proto



<a name="blackfury.vesting.v1.MsgAddAirdrops"></a>

### MsgAddAirdrops
MsgAddAirdrops represents a message to add airdrop targets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `airdrops` | [Airdrop](#blackfury.vesting.v1.Airdrop) | repeated |  |






<a name="blackfury.vesting.v1.MsgAddAirdropsResponse"></a>

### MsgAddAirdropsResponse
MsgMintBySwapResponse defines the Msg/AddAirdrops response type.






<a name="blackfury.vesting.v1.MsgExecuteAirdrops"></a>

### MsgExecuteAirdrops



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `max_count` | [uint64](#uint64) |  | max count of airdrops performed this time |






<a name="blackfury.vesting.v1.MsgExecuteAirdropsResponse"></a>

### MsgExecuteAirdropsResponse







<a name="blackfury.vesting.v1.MsgSetAllocationAddress"></a>

### MsgSetAllocationAddress
MsgSetAllocationAddress represents a message to set allocation address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `team_vesting_addr` | [string](#string) |  |  |
| `strategic_reserve_custodian_addr` | [string](#string) |  |  |






<a name="blackfury.vesting.v1.MsgSetAllocationAddressResponse"></a>

### MsgSetAllocationAddressResponse
MsgSetAllocationAddressResponse defines the Msg/SetAllocationAddress response
type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.vesting.v1.Msg"></a>

### Msg
Msg defines the vesting Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `AddAirdrops` | [MsgAddAirdrops](#blackfury.vesting.v1.MsgAddAirdrops) | [MsgAddAirdropsResponse](#blackfury.vesting.v1.MsgAddAirdropsResponse) | AddAirdrops adds airdrop targets. Should only be called by core team multisig. | GET|/blackfury/vesting/v1/tx/add_airdrops|
| `ExecuteAirdrops` | [MsgExecuteAirdrops](#blackfury.vesting.v1.MsgExecuteAirdrops) | [MsgExecuteAirdropsResponse](#blackfury.vesting.v1.MsgExecuteAirdropsResponse) | ExecuteAirdrops performs airdrops. Should only be called by core team multisig. | GET|/blackfury/vesting/v1/tx/exec_airdrops|
| `SetAllocationAddress` | [MsgSetAllocationAddress](#blackfury.vesting.v1.MsgSetAllocationAddress) | [MsgSetAllocationAddressResponse](#blackfury.vesting.v1.MsgSetAllocationAddressResponse) | SetAllocationAddress sets allocation address of team vesting or strategic_reserve_custodian. | GET|/blackfury/vesting/v1/tx/set_allocation_address|

 <!-- end services -->



<a name="blackfury/voter/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/voter/v1/genesis.proto



<a name="blackfury.voter.v1.GenesisState"></a>

### GenesisState
GenesisState defines the voter module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.voter.v1.Params) |  |  |






<a name="blackfury.voter.v1.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="blackfury/voter/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/voter/v1/query.proto



<a name="blackfury.voter.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="blackfury.voter.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#blackfury.voter.v1.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.voter.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#blackfury.voter.v1.QueryParamsRequest) | [QueryParamsResponse](#blackfury.voter.v1.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/blackfury/voter/v1/params|

 <!-- end services -->



<a name="blackfury/voter/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blackfury/voter/v1/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="blackfury.voter.v1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

