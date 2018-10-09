package Blizzard

type Guild struct {
	Name string `json:"name"`
	Realm string `json:"realm"`
	Battlegroup string `json:"battlegroup"`
	Members int `json:"members"`
	AchievementPoint int `json:"achievementPoint"`
	Emblem Emblem `json:"emblem"`
}

type Emblem struct {
	Icon int `json:"icon"`
	IconColor string `json:"iconColor"`
	IconColorId int `json:"iconColorId"`
	Border int `json:"border"`
	BorderColor string `json:"borderColor"`
	BorderColorId int `json:"borderColorId"`
	backgroundColor string `json:"backgroundColor"`
	BackgroundColorId int `json:"backgroundColorId"`
}
