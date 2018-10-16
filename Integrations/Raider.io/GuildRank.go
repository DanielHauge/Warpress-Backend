package Raider_io

type GuildInfo struct {
	Name string `json:"name"`
	Realm string `json:"realm"`
	Region string `json:"region"`
	RaidRankings RaidRankings `json:"raid_rankings"`
	RaidProgression RaidProgression `json:"raid_progression"`
}

type RaidProgression struct {
	Uldir Progression `json:"uldir"`
}

type Progression struct {
	Summary string `json:"summary"`
	TotalBosses int `json:"total_bosses"`
	NormalBossesKilled int `json:"normal_bosses_killed"`
	HeroicBossesKilled int `json:"heroic_bosses_killed"`
	MythicBossesKilled int `json:"mythic_bosses_killed"`
}

type RaidRankings struct {

	Uldir Raid `json:"uldir"`

}

type Raid struct {
	Normal RaidRank `json:"normal"`
	Heroic RaidRank `json:"heroic"`
	Mythic RaidRank `json:"mythic"`
}

type RaidRank struct {
	World int `json:"world"`
	Region int `json:"region"`
	Realm int `json:"realm"`
}
