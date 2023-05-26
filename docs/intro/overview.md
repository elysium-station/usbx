---
order: 1
---

[//]: # (![Welcome to Blackfury]&#40;../images/blackfury-banner.png&#41;)

# High-level Overview

## What is Blackfury

The Blackfury project is the most decentralized, scalable, and high-throughput blockchain for fractional-algorithmic
stablecoin and various vanward DeFi-specific innovations. It is built
using [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) and [Tendermint Core](https://github.com/tendermint/tendermint)
. It also enables fully compatibility and interchangeability with Ethereum/EVM based DApps, by embedding the evm module
from [Ethermint](https://github.com/tharsis/ethermint).

Absorbed the creations and lessons of some pioneering protocols/projects, the architecture and mechanism of Blackfury are
designed to meet the demands of an increasingly diversified multi/cross-chain world. Blackfury's tokenomics is elaborately
formulated to not only incentivize early adopters, but also encourage long-term cooperation to grow bigger and stronger.

Blackfury's vision is to become the most decentralized, permissionless, high-throughput, low-cost, easy-to-use cross-chain
assets settlement center and experimental DeFi innovation zone.

### Features

Here’s a glance at some key features of Blackfury:

- System native stablecoin **BlackUSD**, or **USBX**
- Pegging through partial backing/collateral and partial algorithmic
- ve(3,3) time-locked of native **Fury** and proportional incentive
- Web3 and EVM compatibility
- High throughput and fast transaction finality via [Tendermint Core](https://github.com/tendermint/tendermint)
- Horizontal scalability via [IBC](https://cosmos.network/ibc) and [Gravity Bridge](https://www.gravitybridge.net)
- Full decentralized on-chain governance and flexible dynamical policy regulation

## Black and Fury

The blackfury system consists of two main native coins, Black and Fury.

- **Black**: Stablecoins that track the price of fiat currencies, and they are named for their fiat counterparts. In the
  early stage of the mainnet launch, it will mainly issue **BlackUSD**, or **USBX**, which tracks/pegs the price of $USD.
- **Fury**: Blackfury blockchain's native staking coin that partially absorbs the price volatility of Black. Users stake
  Fury to validators to add blocks of transactions to the blockchain, and earn various fees and rewards. Holders of Fury
  also can vote on proposals and participate in on-chain governance.

## The Stablecoin Protocol

### Stablecoin

Like [Terra](https://terra.money), Blackfury is also a DeFi-specific blockchain with built-in algorithmic stablecoin
protocol. But the difference is that Blackfury does not rely on any purely algorithmic design which is difficult to grow
and exhibits extreme periods of volatility. Blackfury stablecoin will be minted in two ways:
fractional-backing-algorithmic and over-collateralized-catalytic.

The **fractional-backing-algorithmic**, or **FBA**, is with parts of its backing assets and parts of the algorithmic
supply. The ratio of backing and algorithmic depends on the market's pricing of the Black stablecoin. We named the
ratio **BR (Backing Ratio)**. If BlackUSD is trading at above $1, the system decreases the ratio. If Black is trading at
under $1, the system increases the ratio. At any point in time, BR is determined. If users want to mint Black, they
must **spend** a certain amount of backing assets and a certain amount of Fury coins, which will enter the unified swap
pool. Conversely, when a user wants to acquire backing assets and Fury coins in the swap pool, he must spend a certain
amount of Black coins in exchange, and can only get the proportional coins that follow the BR ratio at the market price.

The **over-collateralized-catalytic**, or **OCC**, is over collateralized for interest-bearing lending, and
loan-to-value maximized by catalytic Fury. Each kind of supported over-collateralized asset forms a separate pool. Users
must pre-deposit over-collateralized assets and then lend Black (
actually minted directly by the system) when needed. The maximum ratio that can be lent (called loan-to-value, or LTV)
depends on the parameters set by the system for this collateral pool, and the additional Fury-boosting minting (
called catalytic) added by the user. If users want to redeem their collateral, they are obliged to repay the principal
and interest of the lent assets, and they must always pay attention to the price fluctuation of the collateral assets,
to avoid triggering the liquidation mechanism.

Why two mintage ways? Well, some users have a lot of real demand for stablecoins and don't want to hold volatile
assets (like Fury, ETH, etc.). Other users want to hold volatile assets for a long time in order to obtain their
appreciation, but have short-term liquidity needs, like to borrow stablecoins, and are willing to pay interest and bear
possible liquidation risks. Exactly, FBA is suitable for the former, and OCC is suitable for the latter.

### Oracle, Maker and Arbitrage

The built-in oracle module is responsible for providing fair and true on-chain prices for specified assets. It accepts
near real-time quotes from active validators, removes deviations, and takes the median as the final standard on-chain
price. Since the price is the most important indicator for the whole system to reflect the real market situation, and it
is also an important reliance on whether the stablecoin can remain pegged, the oracle module will periodically reward
validators who consistently quote correctly.

The most important Black minting/burning activities are handled by the maker module. Whether it is FBA-style spending
mintage or OCC-style collateralized mintage, the maker module will accept the assets deposited by users and securely
host them. The maker module defines a series of transaction types to facilitate the user to mint any desired amount of
Black coins according to the deterministic parameters of the current system and the token price data from the oracle
module. The maker module will charge a certain seigniorage-like fee for the oracle module to incentivize validators who
provide quotes. For OCC-style collateralized mintage, the maker module is also responsible for settling interest and
providing possible liquidation channels.

When possible market fluctuations or black swan events cause Black prices to decouple, the arbitrage behavior of
arbitrageurs helps Black return to their target prices. For FBA-style mintage, when BlackUSD is lower than $1, the system
target BR value automatically increases accordingly. At this time, the system shows that there is a lack of backing
assets, and a surplus of Fury coins. Arbitrageurs can use a certain amount of backing assets to buy Fury coins from the
system at a discounted price, so that the actual BR value tends to the target BR value. An increase in the reserve of
backing assets will increase people's confidence and bring BlackUSD back to the price of $1. For OCC-style mintage, when
BlackUSD is below $1, users tend to buy low-priced BlackUSD from the market and repay the system with BlackUSD's nominal value
of $1 to unlock the collateralized assets. This will reduce the circulation of BlackUSD in the market, and will also bring
BlackUSD back to the price of $1. And for BlackUSD above $1, it is obvious and easy to get it back to $1, so we won't go
into details here.

## Proof-of-Stake and Validators

Inherited from [Tendermint Core](https://github.com/tendermint/tendermint)
and [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), Blackfury uses BFT (Byzantine Fault Tolerance)
consensus protocol to securely and consistently replicate states/blocks/transactions on many machines (or validators)
all over the world. Validators run Blackfury programs called full nodes, and take turns proposing blocks of transactions
and voting on them. Blocks are committed in a chain, with one block at each height. A block may fail to be committed, in
which case the protocol moves to the next round, and a new validator gets to propose a block for that height. Two stages
of voting are required to successfully commit a block; Tendermint call them **pre-vote** and **pre-commit**. A block is
committed when more than 2/3 of validators pre-commit for the same block in the same round. Blackfury can tolerate up to a
1/3 of validators failures, and those failures can include arbitrary behaviour, e.g., hacking and malicious attacks.

Not all validators will have the same "weight" in the consensus protocol. Every validator will have some voting power,
which may not be uniformly distributed across individual validators. Blackfury denominates the weight or stake in our
native Fury coin, and hence the system is often referred to as **Proof-of-Stake**. Validators will be forced to "bond"
or stake their Fury holdings in the security deposit that can be slashed if they're found to misbehave in the consensus
protocol. This adds an economic element to the security of the protocol, allowing one to quantify the cost of violating
the assumption that less than one-third of voting power is Byzantine.

Blackfury allows up to the top 120 validators to participate in consensus. A validator’s rank is determined by their stake
or the total amount of Fury bonded to them. Although validators can bond Fury to themselves, they mainly amass larger
stakes from delegators. Validators with larger stakes get chosen more often to propose new blocks and earn
proportionally more rewards.

Delegators are users who want to receive rewards from consensus without running a full node. Any user that stakes Fury
is a delegator. Delegators stake their Fury to a validator, adding to a validator's weight, or total stake. In return,
delegators receive a portion of system fees as staking rewards.

## The ve(3,3) mechanism and Time-locked Voting Escrow

In addition to the normal Fury staking through the Proof-of-Stake consensus protocol, Blackfury brings in another enhanced
time-locked voting escrow mechanism, called **ve(3,3)**. We use ve(3,3) to incentivize various innovative DeFi DApps in
the Blackfury system as well as getting as many users/investors involved as possible in the governance of the network. We
define the NFT token **veFury** as the vote-escrowed Fury, which is simply Fury locked for a period of time, from 1 week
to 4 years. veFury token holders will receive a multiplied amount of voting power, compared to normal PoS staking. Along
with that, they will gain more staking rewards and voting power on governance proposals. In essence, what is more
important and further is that they will have certain voting rights to incentivize various innovative DeFi DApps in the
system. Of course, they will also be rewarded with extra Fury coins, which come from the reserve in the treasury
according to Blackfury's tokenomics.

The rewarded Fury coins will be distributed weekly. The emission amount are adjusted as a percentage of circulating
supply (`circulating_supply = total_supply - (ve_locked + normal_staked)`). Meaning, assuming the maximum weekly
emission is 500,000, if 0% of the coin is staked or ve-locked, the weekly emission would be 500,000. If 50% of the coin
is staked or ve-locked, the weekly emission would be 250,000. If 100% of the coin is staked or ve-locked, the weekly
emission would be 0.

To ensure that veFury holders are never diluted, their holdings will be increased proportional to the weekly emission.
And since the lock position veFury is tokenized as NFT, it allows veFury to be traded on future secondary markets, as
well as to allow participants to borrow against their veFury in future lending marketplaces. This addresses the capital
inefficiency problem of ve assets, as well as addresses concerns over future liquidity (should it ever be required).

## Smart Contracting and Virtual Machine

As a general-purpose DeFi-specific blockchain platform, Blackfury must integrate certain smart contract virtual machines
to facilitate the deployment of innovative DeFi protocols or DApps by various third-party developers/teams. Fortunately,
nowadays we have many options for mature virtual machine protocols/modules. First and foremost, Blackfury will integrate
the **evm** module from [Ethermint](https://github.com/tharsis/ethermint), to provide native Web3/EVM capabilities and
be compatible with the huge Ethereum ecosystem.

The growth of EVM-based chains (e.g., Ethereum), however, has uncovered several scalability challenges that are often
referred to as
the [Trilemma of decentralization, security, and scalability](https://vitalik.ca/general/2021/04/07/sharding.html).
Developers are frustrated by high gas fees, slow transaction speed & throughput, and chain-specific governance that can
only undergo slow change because of its wide range of deployed applications. A solution is required that eliminates
these concerns for developers, who build applications within a familiar EVM environment.

The evm module provides the EVM familiarity on a scalable, high-throughput Proof-of-Stake blockchain. It allows for the
deployment of smart contracts, interaction with the EVM state machine (state transitions), and the use of EVM tooling.
It alleviates the aforementioned concerns through high transaction throughput
via [Tendermint Core](https://github.com/tendermint/tendermint), fast transaction finality, and horizontal scalability.

## Cross-chain Interoperability

Cross-chain interoperability is increasingly becoming an important infrastructure for all blockchains, and every user
roaming in the crypto world has a strong demand to use cross-chain facilities to meet their growing asset flowing and
interoperability needs.

In the past, many blockchains relied on third-party cross-chain bridges or smart contracts. However, the frequent bugs
and hacking behaviors have caused many users to suffer huge asset losses. At this time, the blockchain infra layer is
often indifferent or powerless. Now we believe that cross-chain is the basic and core component of a blockchain, and
security must be given the utmost importance and guarantee. Therefore, we decide that Blackfury should have built-in
cross-chain protocols and facilities, and rely on Blackfury's PoS consensus and staking validators to strongly ensure
asset security.

For the Cosmos ecosystem, we will introduce [IBC modules and protocols](https://ibcprotocol.org) to natively support
cross-chain communication with other Cosmos SDK-based blockchains. The protocol consists of two distinct layers: the
transport layer which provides the necessary infrastructure to establish secure connections and authenticate data
packets between chains, and the application layer, which defines exactly how these data packets should be packaged and
interpreted by the sending and receiving chains.

For the Ethereum ecosystem, we will introduce [Gravity Bridge](https://github.com/Gravity-Bridge/Gravity-Bridge) as a
built-in module, making it possible to cross-chain assets from Evm-compatible chains to Blackfury or vice versa. Control
of the bridge mirrors the active validator set on Blackfury, and any validator stake on Blackfury can be slashed for
misbehavior.

Not limited to the above, we will also introduce a cross-chain insurance fund in the near future to provide claim
settlement services for users who have suffered losses from cross-chain hacking.

## Governance

To fully embody the decentralized, permissionless nature of Blackfury, we will introduce a full on-chain governance
protocol. Holders of native Fury coin can vote on proposals on a 1 staked token 1 vote basis, and also veFury holders
will gain certain boosting voting power according to length of their locked period.

Any Fury holder can submit proposals by sending TxGovProposal transaction. Once a proposal is submitted, it is
identified by its unique proposalID. Types of proposals include, but are not limited to:

- text proposal that do not involve a modification of the source code, e.g., an opinion poll.
- software upgrade proposal that validators are expected to update their software in accordance with the proposal.
- community pool spend proposal details a proposal for use of community funds, together with how many coins are proposed
  to be spent, and to which recipient account.
- parameter change proposal defines a proposal to change one or more parameters. If accepted, the requested parameter
  change is updated automatically by the proposal handler upon conclusion of the voting period.

## Tokenomics

Blackfury will issue a total supply of **1 bilfury (1,000,000,000) $FURY** at genesis, and there are no inflation after
mainnet launch. Tokens will be released to the following distribution:

- TBD
