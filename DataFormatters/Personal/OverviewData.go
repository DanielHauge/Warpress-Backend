package Personal

type Overview struct {
	Character       character        `json:"character"`
	Guild           guild            `json:"guild"`
	BestParses      []dificultyParse      `json:"best_parses"`
	RaiderIOProfile characterProfile `json:"raider_io_profile"`
}

type StandaloneLogs struct {
	BestParses logs `json:"logs"`
}

type StandaloneRaiderio struct {
	RaiderIOProfile characterProfile `json:"raider_io_profile"`
}

type StandaloneGuild struct {
	GuildRank guildRank `json:"guild_rank"`
}

type characterProfile struct {
	ProfileUrl                       string       `json:"profile_url"`
	MythicPlusRanks                  mythicRanks  `json:"mythic_plus_ranks"`
	MythicPlusRecentRuns             []dungeonRun `json:"mythic_plus_recent_runs"`
	MythicPlusHighestLevelRuns       []dungeonRun `json:"mythic_plus_highest_level_runs"`
	MythicPlusWeeklyHighestLevelRuns []dungeonRun `json:"mythic_plus_weekly_highest_level_runs"`
}

type mythicRanks struct {
	Overall     rank `json:"overall"`
	Dps         rank `json:"dps"`
	Healer      rank `json:"healer"`
	Tank        rank `json:"tank"`
	Class       rank `json:"class"`
	ClassDps    rank `json:"class_dps"`
	ClassHealer rank `json:"class_healer"`
	ClassTank   rank `json:"class_tank"`
}

type rank struct {
	World  int `json:"world"`
	Region int `json:"region"`
	Realm  int `json:"realm"`
}

type dungeonRun struct {
	Dungeon            string  `json:"dungeon"`
	ShortName          string  `json:"short_name"`
	MythicLevel        int     `json:"mythic_level"`
	CompletedAt        string  `json:"completed_at"`
	ClearTimeMs        int     `json:"clear_time_ms"`
	NumKeystoneUpgrade int     `json:"num_keystone_upgrade"`
	Score              float32 `json:"score"`
	Affixes            []affix `json:"affixes"`
	Url                string  `json:"url"`
}

type affix struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WowheadUrl  string `json:"wowhead_url"`
}

type logs struct {
	BestParses []dificultyParse `json:"best_parses"`
}
/*
type Parses struct {
	Mythic dificultyParse `json:"mythic"`
	Heroic dificultyParse    `json:"heroic"`
	Normal dificultyParse    `json:"normal"`
}
*/

type dificultyParse struct {
	Logs []encounter `json:"logs"`
	Specs []string `json:"specs"`
	Difficulty string `json:"difficulty"`
}

type encounter struct {
	EncounterID        int     `json:"encounterID"`
	EncounterName      string  `json:"encounterName"`
	Spec               string  `json:"spec"`
	Rank               int     `json:"rank"`
	OutOf              int     `json:"outOf"`
	Duration           int     `json:"duration"`
	ReportID           string  `json:"reportID"`
	CharacterName      string  `json:"characterName"`
	Percentile         int     `json:"percentile"`
	ItemLevelKeyOrPath int     `json:"ilvlKeyOrPatch"`
	Total              float64 `json:"total"`
}

type character struct {
	Name    string         `json:"name"`
	Realm   string         `json:"realm"`
	SluggedRealm string `json:"slugged_realm"`
	Class   int            `json:"class"`
	Race    int            `json:"race"`
	Stats   Stats         `json:"stats"`
	Gender  int            `json:"gender"`
	Level   int            `json:"level"`
	Avatar  string         `json:"avatar"`
	Main    string         `json:"main"`
	Faction int            `json:"faction"`
	Spec    specialization `json:"spec"`
	Items   items          `json:"items"`
}

type specialization struct {
	Name            string       `json:"name"`
	Role            string       `json:"role"`
	BackgroundImage string       `json:"backgroundImage"`
	Icon            string       `json:"icon"`
	Description     string       `json:"description"`
	MasterySpellID 	int 			`json:"mastery_spell_id"`
	Order           int          `json:"order"`
	Talents         []talentTier `json:"talents"`
}

type Stats struct {
	Health int `json:"health"`
	PowerType string `json:"powerType"`
	Power int `json:"power"`
	Str int `json:"str"`
	Agi int `json:"agi"`
	Int int `json:"int"`
	Sta int `json:"sta"`
	SpeedRating float32 `json:"speedRating"`
	SpeedRatingBonus float32 `json:"speedRatingBonus"`
	Crit float32 `json:"crit"`
	CritRating int `json:"critRating"`
	Haste float32 `json:"haste"`
	HasteRating int `json:"hasteRating"`
	HasteRatingPercent float32 `json:"hasteRatingPercent"`
	Mastery float32 `json:"mastery"`
	MasteryRating int `json:"masteryRating"`
	Leech float32 `json:"leech"`
	LeechRating float32 `json:"leechRating"`
	LeechRatingBonus float32 `json:"leechRatingBonus"`
	Versatility int `json:"versatility"`
	VersatilityDamageDoneBonus float32 `json:"versatilityDamageDoneBonus"`
	VersatilityHealingDoneBonus float32 `json:"versatilityHealingDoneBonus"`
	VersatilityDamageTakenBonus float32 `json:"versatilityDamageTakenBonus"`
	AvoidanceRating float32 `json:"avoidanceRating"`
	AvoidanceRatingBonus float32 `json:"avoidanceRatingBonus"`
	SpellPen int `json:"spellPen"`
	SpellCrit float32 `json:"spellCrit"`
	SpellCritRating int `json:"spellCritRating"`
	Armor int `json:"armor"`
	Dodge float32 `json:"dodge"`
	DodgeRating int `json:"dodgeRating"`
	MainHandDmgMin float32 `json:"mainHandDmgMin"`
	MainHandDmgMax float32 `json:"mainHandDmgMax"`
	MainHandSpeed float32 `json:"mainHandSpeed"`
	MainHandDps float32 `json:"mainHandDps"`
	OffHandDmgMin float32 `json:"offHandDmgMin"`
	OffHandDmgMax float32 `json:"offHandDmgMax"`
	OffHandSpeed float32 `json:"offHandSpeed"`
	OffHandDps float32 `json:"offHandDps"`
	RangedDmgMin float32 `json:"offHandDmgMin"`
	RangedDmgMax float32 `json:"offHandDmgMax"`
	RangedSpeed float32 `json:"offHandSpeed"`
	RangedDps float32 `json:"offHandDps"`
}

type talentTier struct {
	Tier   int   `json:"tier"`
	Column int   `json:"column"`
	Spell  spell `json:"spell"`
}

type spell struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
	Cooldown    string `json:"cooldown"`
}

type items struct {
	AverageItemLevel      int  `json:"averageItemLevel"`
	AverItemLevelEquipped int  `json:"averItemLevelEquipped"`
	Head                  item `json:"head"`
	Neck                  item `json:"neck"`
	Shoulder              item `json:"shoulder"`
	Back                  item `json:"back"`
	Chest                 item `json:"chest"`
	Wrist                 item `json:"wrist"`
	Hands                 item `json:"hands"`
	Waist                 item `json:"waist"`
	Legs                  item `json:"legs"`
	Feet                  item `json:"feet"`
	Finger1               item `json:"finger1"`
	Finger2               item `json:"finger2"`
	Trinket1              item `json:"trinket1"`
	Trinket2              item `json:"trinket2"`
	MainHand              item `json:"mainHand"`
	OffHand               item `json:"offHand"`
}

type item struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	Icon                 string         `json:"icon"`
	Quality              int            `json:"quality"`
	ItemLevel            int            `json:"itemLevel"`
	BonusLists           []int          `json:"bonusLists"`
	Appearance           int            `json:"appearance"`
	AzeriteItem          azeriteItem    `json:"azeriteItem"`
	AzeriteEmpoweredItem []azeritePower `json:"azeritePowers"`
	Gem0                 int            `json:"gem_0"`
	Enchant              int            `json:"enchant"`
}

type azeritePower struct {
	Id          int `json:"id"`
	Tier        int `json:"tier"`
	SpellId     int `json:"spellId"`
	BonusListId int `json:"bonusListId"`
}

type azeriteItem struct {
	AzeriteLevel               int `json:"azeriteLevel"`
	AzeriteExperience          int `json:"azeriteExperience"`
	AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
}

type stat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

type guild struct {
	Name      string    `json:"name"`
	Realm     string    `json:"realm"`
	Members   int       `json:"members"`
	Emblem    emblem    `json:"emblem"`
	GuildRank guildRank `json:"guild_rank"`
}

type guildRank struct {
	Score     int `json:"score"`
	WorldRank int `json:"world_rank"`
	AreaRank  int `json:"area_rank"`
	RealmRank int `json:"realm_rank"`
}

type emblem struct {
	Icon              int    `json:"icon"`
	IconColor         string `json:"iconColor"`
	IconColorId       int    `json:"iconColorId"`
	Border            int    `json:"border"`
	BorderColor       string `json:"borderColor"`
	BorderColorId     int    `json:"borderColorId"`
	BackgroundColor   string `json:"backgroundColor"`
	BackgroundColorId int    `json:"backgroundColorId"`
}
