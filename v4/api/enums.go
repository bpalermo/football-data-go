package api

type CardType string
type CompetitionType string
type GoalType string
type MatchGroup string
type MatchStatus string
type PenaltyType string
type Plan string
type ScoreDuration string
type TeamType string

const (
	CARD_TYPE_YEALLOW     CardType = "YEALLOW"
	CARD_TYPE_YEALLOW_RED CardType = "YEALLOW_RED"
	CARD_TYPE_RED         CardType = "RED"

	COMPETITION_TYPE_LEAGUE     CompetitionType = "LEAGUE"
	COMPETITION_TYPE_LEAGUE_CUP CompetitionType = "LEAGUE_CUP"
	COMPETITION_TYPE_CUP        CompetitionType = "CUP"
	COMPETITION_TYPE_PLAYOFFS   CompetitionType = "PLAYOFFS"

	GOAL_TYPE_REGULAR GoalType = "REGULAR"
	GOAL_TYPE_OWN     GoalType = "OWN"
	GOAL_TYPE_PENALTY GoalType = "PENALTY"

	MATCH_GROUP_A MatchGroup = "GROUP_A"
	MATCH_GROUP_B MatchGroup = "GROUP_B"
	MATCH_GROUP_C MatchGroup = "GROUP_C"
	MATCH_GROUP_D MatchGroup = "GROUP_D"
	MATCH_GROUP_E MatchGroup = "GROUP_E"
	MATCH_GROUP_F MatchGroup = "GROUP_F"
	MATCH_GROUP_G MatchGroup = "GROUP_G"
	MATCH_GROUP_H MatchGroup = "GROUP_G"
	MATCH_GROUP_I MatchGroup = "GROUP_I"
	MATCH_GROUP_J MatchGroup = "GROUP_J"
	MATCH_GROUP_K MatchGroup = "GROUP_K"
	MATCH_GROUP_L MatchGroup = "GROUP_L"

	MATCH_STATUS_SCHEDULED        MatchStatus = "SCHEDULED"
	MATCH_STATUS_TIMED            MatchStatus = "TIMED"
	MATCH_STATUS_IN_PLAY          MatchStatus = "IN_PLAY"
	MATCH_STATUS_PAUSED           MatchStatus = "PAUSED"
	MATCH_STATUS_EXTRA_TIME       MatchStatus = "EXTRA_TIME"
	MATCH_STATUS_PENALTY_SHOOTOUT MatchStatus = "PENALTY_SHOOTOUT"
	MATCH_STATUS_FINISHED         MatchStatus = "FINISHED"
	MATCH_STATUS_SUSPENDED        MatchStatus = "SUSPENDED"
	MATCH_STATUS_POSTPONED        MatchStatus = "POSTPONED"
	MATCH_STATUS_CANCELLED        MatchStatus = "CANCELLED"
	MATCH_STATUS_AWARDED          MatchStatus = "AWARDED"

	PENALTY_TYPE_MATCH    PenaltyType = "MATCH"
	PENALTY_TYPE_SHOOTOUT PenaltyType = "SHOOTOUT"

	PLAN_TIER_ONE   Plan = "TIER_ONE"
	PLAN_TIER_TWO   Plan = "TIER_TWO"
	PLAN_TIER_THREE Plan = "TIER_THREE"
	PLAN_TIER_FOUR  Plan = "TIER_FOUR"

	SCORE_DURATION_REGULAR          ScoreDuration = "REGULAR"
	SCORE_DURATION_EXTRA_TIME       ScoreDuration = "EXTRA_TIME"
	SCORE_DURATION_PENALTY_SHOOTOUT ScoreDuration = "PENALTY_SHOOTOUT"

	TEAM_TYPE_CLUB     TeamType = "CLUB"
	TEAM_TYPE_NATIONAL TeamType = "NATIONAL"
)
