package Raider_io

type CharInput struct {
	Name   string
	Realm  string
	Region string
}

type CharacterProfile struct {
	Name                             string       `json:"name"`
	Race                             string       `json:"race"`
	Class                            string       `json:"class"`
	Spec                             string       `json:"active_spec_name"`
	Role                             string       `json:"active_spec_role"`
	Gender                           string       `json:"gender"`
	Faction                          string       `json:"faction"`
	AchievementPoints                int          `json:"achievement_points"`
	HonorableKills                   int          `json:"honorable_kills"`
	ThumbnailUrl                     string       `json:"thumbnail_url"`
	Region                           string       `json:"region"`
	Realm                            string       `json:"realm"`
	ProfileUrl                       string       `json:"profile_url"`
	MythicPlusRanks                  MythicRanks  `json:"mythic_plus_ranks"`
	MythicPlusRecentRuns             []DungeonRun `json:"mythic_plus_recent_runs"`
	MythicPlusHighestLevelRuns       []DungeonRun `json:"mythic_plus_highest_level_runs"`
	MythicPlusWeeklyHighestLevelRuns []DungeonRun `json:"mythic_plus_weekly_highest_level_runs"`
}

type MythicRanks struct {
	Overall     Rank `json:"overall"`
	Dps         Rank `json:"dps"`
	Healer      Rank `json:"healer"`
	Tank        Rank `json:"tank"`
	Class       Rank `json:"class"`
	ClassDps    Rank `json:"class_dps"`
	ClassHealer Rank `json:"class_healer"`
	ClassTank   Rank `json:"class_tank"`
}

type Rank struct {
	World  int `json:"world"`
	Region int `json:"region"`
	Realm  int `json:"realm"`
}

type DungeonRun struct {
	Dungeon            string  `json:"dungeon"`
	ShortName          string  `json:"short_name"`
	MythicLevel        int     `json:"mythic_level"`
	CompletedAt        string  `json:"completed_at"`
	ClearTimeMs        int     `json:"clear_time_ms"`
	NumKeystoneUpgrade int     `json:"num_keystone_upgrade"`
	MapChallengeModeId int     `json:"map_challenge_mode_id"`
	Score              float32 `json:"score"`
	Affixes            []Affix `json:"affixes"`
	Url                string  `json:"url"`
}

type Affix struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WowheadUrl  string `json:"wowhead_url"`
}
