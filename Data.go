package main

import (
	"./Blizzard"
	"./WarcraftLogs"
	"./Raider.io"
	"./Wowprogress"
	)



type PersonalProfile struct {
	Character Blizzard.FullCharInfo `json:"character"`
	WarcraftLogsRanks []WarcraftLogs.Encounter `json:"warcraft_logs_ranks"`
	RaiderIOProfile Raider_io.CharacterProfile `json:"raider_io_profile"`
	GuildRank Wowprogress.GuildRank `json:"guild_rank"`
}


