package PreparedProcedures

import (
	. "../../Postgres"
	. "../DataModel"
	"database/sql"
	"github.com/pkg/errors"
)

func AddAddon(name string, twitchlink string, guildid int) error{
	_, e := Execute("INSERT INTO addons "+
		"(name, twitchlink, guildid) "+
		"VALUES ($1, $2, $3);", name, twitchlink, guildid)
	return e
}

func EditAddon(name string, twitchlink string, guildid int, id int) error{
	_, e := Execute("UPDATE addons "+
		"SET name = $1, twitchlink = $2 "+
		"WHERE id = $3 AND guildid = $4;", name, twitchlink, id, guildid)
	return e
}

func DeleteAddon(guildid int, id int)error{
	_, e := Execute("DELETE FROM addons WHERE id = $1 AND guildid = $2", id, guildid)
	return e
}

func GetAddon(guildid int) ([]Addon, error){
	var res []Addon
	e := QueryMultiple("SELECT * FROM addons WHERE guildid=$1;",
		[]interface{}{guildid},
		&res,
	)

	if e == sql.ErrNoRows{
		e = errors.New("No Addons have been added for this guild")
	}

	return res, e
}