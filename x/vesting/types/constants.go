package types

import blackfury "github.com/elysium-station/blackfury/types"

const (
	StakingRewardVestingName = "staking_reward_vesting"
	CommunityPoolVestingName = "community_pool_vesting"
	TeamVestingName          = "team_vesting"

	// Strate reserve pool controlled by governance.
	// Not used now, maybe future.
	StrategicReservePoolName = "strategic_reserve_pool"

	StakingRewardVestingTime = blackfury.SecondsPer4Years
	CommunityPoolVestingTime = blackfury.SecondsPer4Years
	TeamVestingTime          = blackfury.SecondsPer4Years

	ClaimVestedPeriod = 10
)
