package PreparedProcedures

import (
	. "../../Postgres"
	. "../DataModel"
	"time"
)

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
