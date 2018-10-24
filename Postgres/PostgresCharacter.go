package Postgres

import (
	log "../Utility/Logrus"
	"database/sql"
)

// TODO: Not the way to do it.

func SetMain(id int, name string, realm string, region string) {

}

func GetMain(id int) {
	return db.Exec("exec GETMAIN;", id)
}
