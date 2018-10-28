package PreparedProcedures

import (
	. "../../Postgres"
	. "../DataModel"
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
