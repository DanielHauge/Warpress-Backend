package main

import (
	"./Blizzard"
	"./WarcraftLogs"
	"./Raider.io"
	"./Wowprogress"
	)

type ExampleInput struct {
	ExampleString string `json:"example_string"`
	ExampleInteger int `json:"example_integer"`
}

type ExampleOutput struct {
	ExampleListOfIntergers []int `json:"example_list_of_intergers"`
}

type PersonalProfile struct {
	Character Blizzard.FullCharInfo `json:"character"`
	WarcraftLogsRanks []WarcraftLogs.Encounter `json:"warcraft_logs_ranks"`
	RaiderIOProfile Raider_io.CharacterProfile `json:"raider_io_profile"`
	GuildRank Wowprogress.GuildRank `json:"guild_rank"`
}


