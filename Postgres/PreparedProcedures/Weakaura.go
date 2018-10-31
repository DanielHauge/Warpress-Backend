package PreparedProcedures

import (
	. "../../Postgres"
	. "../DataModel"
	"database/sql"
	"github.com/pkg/errors"
)

func AddWeakaura(guildid int, name string, link string, imp string) error {
	_, e := Execute("INSERT INTO weakauras "+
		"(guildid, name, link, import) "+
		"VALUES ($1, $2, $3, $4);", guildid, name, link, []byte(imp))
	return e
}

func EditWeakaura(guildid int, name string, link string, imp string, id int) error {
	_, e := Execute("UPDATE weakauras "+
		"SET name = $1, link = $2, imp = $3 "+
		"WHERE id = $5 AND guildid = $4;", name, link, imp, guildid, id)
	return e
}

func DeleteWeakaura(guildid int, id int) error {
	_, e := Execute("DELETE FROM weakauras WHERE id = $1 AND guildid = $2", id, guildid)
	return e
}

func GetWeakaura(guildid int) ([]Weakaura, error) {
	var res []Weakaura
	e := QueryMultiple("SELECT * FROM weakauras WHERE guildid=$1;",
		[]interface{}{guildid},
		&res,
	)
	if e == sql.ErrNoRows {
		e = errors.New("No Weakaura have been added for this guild")
	}
	return res, e
}
