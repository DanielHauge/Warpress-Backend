package main

import (
	"./DataFormatters/Guild"
	"./DataFormatters/Personal"
	"./Integrations/BlizzardOpenAPI"
	"./Integrations/Raider.io"
	"./Integrations/WarcraftLogs"
	"github.com/bxcodec/faker"
	"log"
)

type ExampleInput struct {
	ExampleString  string `json:"example_string"`
	ExampleInteger int    `json:"example_integer"`
}

type ExampleOutput struct {
	ExampleListOfIntergers []int `json:"example_list_of_intergers"`
}

type ExamplePleaseTryIt struct {
	AlotOfJson string `json:"alot_of_json"`
}

type ExampleChar struct {
	Name  string `json:"name"`
	Realm string `json:"realm"`
}

var ExampleFullBlizzChar = new(BlizzardOpenAPI.FullCharInfo)
var ExamplePersonalProfile = new(Personal.PersonalProfile)
var ExamplePersonalImprovement = new(Personal.PersonalImprovement)
var ExampleRaiderioProfile = new(Raider_io.CharacterProfile)
var ExampleWarcraftlogs = new([]WarcraftLogs.Encounter)
var ExampleGuildOverviewInfo = new(Guild.FullGuildOverviewInfo)
var ExampleCharVar = new(ExampleChar)

func init() {
	/*
		e := faker.FakeData(ExampleFullBlizzChar)

		e = faker.FakeData(ExamplePersonalProfile)
		e = faker.FakeData(ExamplePersonalImprovement)
		e = faker.FakeData(ExampleRaiderioProfile)
		e = faker.FakeData(ExampleWarcraftlogs)
		e = faker.FakeData(ExampleGuildOverviewInfo)
	*/
	e := faker.FakeData(ExampleCharVar)
	if e != nil {
		log.Println(e.Error())
	}
}
