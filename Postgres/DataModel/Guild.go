package DataModel

import "time"

type Guild struct {
	Name    string
	Realm   string
	Region  string
	Officer int
	Raider  int
	Trial   int
	Id      int
}

type RaidNight struct {
	Duration time.Duration
	Start    time.Duration
	Day      time.Weekday
	GuildId  int
	Id       int
}
