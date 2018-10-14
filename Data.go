package main

import (
	"./Blizzard"
	"./Raider.io"
	"./WarcraftLogs"
	"./Wowprogress"
)



type PersonalProfile struct {
	Character Blizzard.FullCharInfo `json:"character"`
	WarcraftLogsRanks []WarcraftLogs.Encounter `json:"warcraft_logs"`
	RaiderIOProfile Raider_io.CharacterProfile `json:"raider_io_profile"`
	GuildRank Wowprogress.GuildRank `json:"guild_rank"`
}




