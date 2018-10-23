package BlizzardOpenAPI

type FullCharInfo struct {
	LastModified        int    `json:"lastModified"`
	Name                string `json:"name"`
	Realm               string `json:"realm"`
	Battlegroup         string `json:"battlegroup"`
	Class               int    `json:"class"`
	Race                int    `json:"race"`
	Gender              int    `json:"gender"`
	Level               int    `json:"level"`
	AchievementPoints   int    `json:"achievementPoints"`
	Thumbnail           string `json:"thumbnail"`
	CalcClass           string `json:"calcClass"`
	Faction             int    `json:"faction"`
	TotalHonorableKills int    `json:"totalHonorableKills"`
	Guild               Guild  `json:"guild"`
	Talents 			[]Specialization `json:"talents"`
	Items               Items  `json:"items"`
}

type Specialization struct{
	Selected bool `json:"selected"`
	Talents []TalentTier `json:"talents"`
	Spec SpecInfo `json:"spec"`
	CalcTalent string `json:"calcTalent"`
	CalcSpec string `json:"calcSpec"`
}

type TalentTier struct {
	tier int `json:"tier"`
	column int `json:"column"`
	Spec SpecInfo `json:"spec"`
	Spell Spell `json:"spell"`
}

type Spell struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Description string `json:"description"`
	CastTime string `json:"castTime"`
	Cooldown string `json:"cooldown"`
}

type SpecInfo struct {
	Name string `json:"name"`
	Role string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon string `json:"icon"`
	Description string `json:"description"`
	Order int `json:"order"`

}



type Items struct {
	AverageItemLevel      int  `json:"averageItemLevel"`
	AverItemLevelEquipped int  `json:"averItemLevelEquipped"`
	Head                  Item `json:"head"`
	Neck                  Item `json:"neck"`
	Shoulder              Item `json:"shoulder"`
	Back                  Item `json:"back"`
	Chest                 Item `json:"chest"`
	Wrist                 Item `json:"wrist"`
	Hands                 Item `json:"hands"`
	Waist                 Item `json:"waist"`
	Legs                  Item `json:"legs"`
	Feet                  Item `json:"feet"`
	Finger1               Item `json:"finger1"`
	Finger2               Item `json:"finger2"`
	Trinket1              Item `json:"trinket1"`
	Trinket2              Item `json:"trinket2"`
	MainHand              Item `json:"mainHand"`
	OffHand               Item `json:"offHand"`
}

type Item struct {
	Id                   int             `json:"id"`
	Name                 string          `json:"name"`
	Icon                 string          `json:"icon"`
	Quality              int             `json:"quality"`
	ItemLevel            int             `json:"itemLevel"`
	ToolTipParams        ToolTipParams   `json:"tooltipParams"`
	Stats                []Stat          `json:"stats"`
	Armor                int             `json:"armor"`
	WeaponInfo           WeaponInfo      `json:"weaponInfo"`
	Context              string          `json:"context"`
	BonusLists           []int           `json:"bonusLists"`
	ArtifactId           int             `json:"artifactId"`
	DisplayInfoId        int             `json:"displayInfoId"`
	ArtifactAppearanceId int             `json:"artifactAppearanceId"`
	ArtifactTraits       []ArtifactTrait `json:"artifactTraits"`
	Relics               []Relic         `json:"relics"`
	Appearance           Apperence       `json:"appearance"`
	AzeriteItem          AzeriteItem     `json:"azeriteItem"`
	AzeriteEmpoweredItem AzeriteEmpItem  `json:"azeriteEmpoweredItem"`
}

type WeaponInfo struct {
	Damage      Dmg     `json:"damage"`
	WeaponSpeed float32 `json:"weaponSpeed"`
	Dps         float32 `json:"dps"`
}

type Dmg struct {
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
	ExactMin float32 `json:"exactMin"`
	ExactMax float32 `json:"exactMax"`
}

type AzeriteEmpItem struct {
	AzeritePowers []AzeritePower `json:"azeritePowers"`
}

type AzeritePower struct {
	Id          int `json:"id"`
	Tier        int `json:"tier"`
	SpellId     int `json:"spellId"`
	BonusListId int `json:"bonusListId"`
}

type AzeriteItem struct {
	AzeriteLevel               int `json:"azeriteLevel"`
	AzeriteExperience          int `json:"azeriteExperience"`
	AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
}

type Apperence struct {
	ItemAppearanceModId int `json:"itemAppearanceModId"`
}

type ArtifactTrait struct {
}

type Relic struct {
}

type ToolTipParams struct {
	Gem0              int `json:"gem0"`
	AzeritePower0     int `json:"azeritePower0"`
	AzeritePower1     int `json:"azeritePower1"`
	AzeritePower2     int `json:"azeritePower2"`
	AzeritePower3     int `json:"azeritePower3"`
	AzeritePowerLevel int `json:"azeritePowerLevel"`
}

type Stat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

