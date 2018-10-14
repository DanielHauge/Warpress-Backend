package WarcraftLogs

type CharInput struct {
	Name string
	Realm string
	Region string
}


type Encounter struct {
	EncounterID int `json:"encounterID"`
	EncounterName string `json:"encounterName"`
	Class string `json:"class"`
	Spec string `json:"spec"`
	Rank int `json:"rank"`
	outOf int `json:"outOf"`
	Duration int `json:"duration"`
	ReportID string `json:"reportID"`
	FightID int `json:"fightID"`
	Difficulty int `json:"difficulty"`
	CharacterName string `json:"characterName"`
	Server string `json:"server"`
	Percentile int `json:"percentile"`
	ItemLevelKeyOrPath int `json:"ilvlKeyOrPatch"`
	Talents []Talent `json:"talents"`
	Gear []GearPeace `json:"gear"`
	Total float64 `json:"total"`
	Estimated bool `json:"estimated"`
}

type Talent struct {
	Name string `json:"name"`
	Id int `json:"id"`
	Icon string `json:"icon"`
}

type GearPeace struct {
	Name string `json:"name"`
	Id int `json:"id"`
	Icon string `json:"icon"`
	Quality string `json:"quality"`
}

type friendly struct {
	Name string `json:"name"`
	Id int `json:"id"`
}

type Report struct {
	Friendlies []friendly `json:"friendlies"`
}