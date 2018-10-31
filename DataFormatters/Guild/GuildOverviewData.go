package Guild

import (
	"../../Integrations/Raider.io"
	"../../Integrations/WarcraftLogs"
)

type Overview struct {
	Name               string                      `json:"name"`
	Realm              string                      `json:"realm"`
	SluggedRealm       string                      `json:"slugged_realm"`
	LastRaid           string                      `json:"last_raid"`
	NextRaid           string                      `json:"next_raid"`
	TimeUntilNextRaid  string                      `json:"time_until_next_raid"`
	Roster             []Member                    `json:"roster"`
	Progress           Raider_io.RaidProgression   `json:"progress"`
	WarcraftlogReports []WarcraftLogs.GuildReports `json:"warcraftlog_reports"`
	RaidDays           []RaidNight                 `json:"raid_days"`
}

type Member struct {
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Role  string `json:"role"`
	Class int    `json:"class"`
}

type RaidNight struct {
	DayOfTheWeek int    `json:"day_of_the_week"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}
