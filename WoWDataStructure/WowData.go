package WoWDataStructure

type FullCharInfo struct {
	LastModified int `json:"lastModified"`
	Name string `json:"name"`
	Realm string `json:"realm"`
	Battlegroup string `json:"battlegroup"`
	Class int `json:"class"`
	Race int `json:"race"`
	Gender int `json:"gender"`
	Level int `json:"level"`
	AchievementPoints string `json:"achievementPoints"`
	Thumbnail string `json:"thumbnail"`
	CalcClass string `json:"calcClass"`
	Faction int `json:"faction"`
	TotalHonorableKills int `json:"totalHonorableKills"`
	Items []Item `json:"items"`
}

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Quality int `json:"quality"`
	ItemLevel int `json:"itemLevel"`
	ToolTipParams ToolTipParams `json:"tooltipParams"`
	stats []Stat `json:"stats"`
	Armor int `json:"armor"`
	WeaponInfo
	Context string `json:"context"`
	bonusLists []int `json:"bonusLists"`
	ArtifactId int `json:"artifactId"`
	DisplayInfoId int `json:"displayInfoId"`
	ArtifactAppearanceId string `json:"artifactAppearanceId"`
	ArtifactTraits []ArtifactTrait `json:"artifactTraits"`
	Relics []Relic `json:"relics"`
	Appearance Apperence `json:"appearance"`
	AzeriteItem AzeriteItem `json:"azeriteItem"`
	AzeriteEmpoweredItem AzeriteEmpItem `json:"azeriteEmpoweredItem"`
}

type WeaponInfo struct {
	Damage Dmg `json:"damage"`
	WeaponSpeed float32 `json:"weaponSpeed"`
	Dps float32 `json:"dps"`
}

type Dmg struct {
	Min int `json:"min"`
	Max int `json:"max"`
	ExactMin int `json:"exactMin"`
	ExactMax int `json:"exactMax"`
}

type AzeriteEmpItem struct {
	AzeritePowers []AzeritePower `json:"azeritePowers"`
}

type AzeritePower struct {
	Id int `json:"id"`
	Tier int `json:"tier"`
	SpellId int `json:"spellId"`
	BonusListId int `json:"bonusListId"`
}

type AzeriteItem struct {
	AzeriteLevel int `json:"azeriteLevel"`
	AzeriteExperience int `json:"azeriteExperience"`
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
	Gem0 int `json:"gem0"`
	AzeritePower0 int `json:"azeritePower0"`
	AzeritePower1 int `json:"azeritePower1"`
	AzeritePower2 int `json:"azeritePower2"`
	AzeritePower3 int `json:"azeritePower3"`
	AzeritePowerLevel int `json:"azeritePowerLevel"`
}

type Stat struct {
	Stat int `json:"stat"`
	Amount int `json:"amount"`
}