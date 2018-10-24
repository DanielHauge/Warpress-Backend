package Postgres

import (
	. "./DataModel"
)
// TODO: Not the way to do it.

func SetMain(id int, name string, realm string, region string) error {
	_, e := Execute("INSERT INTO main (accountId, name, realm, region) " +
		"VALUES ($1, $2, $3, $4) " +
		"ON CONFLICT (accountId) DO UPDATE " +
		"SET name = $2, realm = $3, region = $4;", id, name, realm, region)
	return e
}

func GetMain(id int) (string, string, string, error) {
	var res Character
	e := QuerySingle("SELECT name, realm, region FROM main WHERE accountid=$1", []interface{}{id}, &res.Name, &res.Realm, &res.Region)
	return res.Name, res.Realm, res.Region, e
}
