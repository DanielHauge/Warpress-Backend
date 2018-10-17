package Personal

import (
	"../Integrations/BlizzardOpenAPI"
	"../Integrations/Raider.io"
	"../Integrations/WarcraftLogs"
	"../Integrations/Wowprogress"
)

type PersonalProfile struct {
	Character         BlizzardOpenAPI.FullCharInfo `json:"character"`
	WarcraftLogsRanks []WarcraftLogs.Encounter     `json:"warcraft_logs"`
	RaiderIOProfile   Raider_io.CharacterProfile   `json:"raider_io_profile"`
	GuildRank         Wowprogress.GuildRank        `json:"guild_rank"`
}
