package main

import "./Integrations/BlizzardOpenAPI"

type ExampleInput struct {
	ExampleString string `json:"example_string"`
	ExampleInteger int `json:"example_integer"`
}

type ExampleOutput struct {
	ExampleListOfIntergers []int `json:"example_list_of_intergers"`
}

type ExamplePleaseTryIt struct{
	AlotOfJson string `json:"alot_of_json"`
}

var ExampleFullBlizzChar = BlizzardOpenAPI.FullCharInfo{
	LastModified: 23321321,
	Name: "Rakhoal",
	Realm: "Twisting-Nether",
	Battlegroup: "BattleWorld / Yoyoyo",
	Class: 5,
	Race: 9,
	Gender: 0,
	Level: 120,
	AchievementPoints: 14090,
	Thumbnail: "SomeThumbNailString",
	CalcClass: "Some weird String -> Please go try the API, the items are long!",
	Faction: 0,
	TotalHonorableKills: 9001,
	Guild: BlizzardOpenAPI.Guild{
		Name: "Time in Motion",
		Realm: "Twisting-Nether",
		Battlegroup: "BattleWorld / Yoyoyo",
		Members: 502,
		AchievementPoint: 9020,
		Emblem: BlizzardOpenAPI.Emblem{
			Icon: 5,
			IconColor: "Blue",
			IconColorId: 5,
			Border: 10,
			BorderColor: "Red",
			BorderColorId: 13,
			BackgroundColorId: 2,
		},
	},
}
