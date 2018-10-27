package Postgres

import (
	. "./DataModel"
	"time"
)

func RegisterGuild(name string, realm string, region string, officer int, raider int, trial int) error {
	_, e := Execute("INSERT INTO guild "+
		"(name, realm, region, officerrank, raiderrank, trialrank) "+
		"VALUES ($1, $2, $3, $4, $5, $6);", name, realm, region, officer, raider, trial)
	return e
}

func GetGuildByID(guildid int) (Guild, error) {
	var res Guild
	e := QuerySingle("SELECT name, realm, region, officerrank, raiderrank, trialrank, guildid FROM guild WHERE guildid=$1;",
		[]interface{}{guildid},
		&res.Name, &res.Realm, &res.Region, &res.Officer, &res.Raider, &res.Trial, &res.Id)
	return res, e
}

func GetGuildByComposite(name string, realm string, region string) (Guild, error) {
	var res Guild
	e := QuerySingle("SELECT name, realm, region, officerrank, raiderrank, trialrank, guildid "+
		"FROM guild WHERE name LIKE $1 "+
		"AND realm LIKE $2 "+
		"AND region LIKE $3;",
		[]interface{}{name, realm, region},
		&res.Name, &res.Realm, &res.Region, &res.Officer, &res.Raider, &res.Trial, &res.Id)
	return res, e
}

func AddRaidNight(duration time.Duration, start time.Duration, dayoftheweek int, guildid int) error {
	// Make sure it doesn't overlap.
	_, e := Execute("INSERT INTO raidnight "+
		"(duration, start, dayoftheweek, guildid) "+
		"VALUES ($1, $2, $3, $4);", duration, start, dayoftheweek, guildid)
	return e
}

func EditRaidNight(duration time.Duration, start time.Time, dayoftheweek int, id int, guildid int) error {
	// Make sure it doesn't overlap
	_, e := Execute("UPDATE raidnight "+
		"SET duration = $1, start = $2, dayoftheweek = $3 "+
		"WHERE id = $4 AND guildid = $5;", duration, start, dayoftheweek, id, guildid)
	return e
}

func DeleteRaidNight(id int, guildid int) error {
	_, e := Execute("DELETE FROM raidnight WHERE id = $1 AND guildid = $2", id, guildid)
	return e
}

func GetRaidNights(guildid int) ([]RaidNight, error) {
	var res []RaidNight
	e := QueryMultiple("SELECT duration, start, dayoftheweek, guildid, id FROM raidnight WHERE guildid=$1;",
		[]interface{}{guildid},
		&res,
	)

	return res, e
}
