package Postgres

import (
	log "../Utility/Logrus"
	"database/sql"
)

var db *sql.DB

func init() {
	connectionPool, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.WithLocation().WithError(err).Fatal("Could not establish connection pool for postgres")
	}
	e := connectionPool.Ping()
	if e != nil {
		log.WithLocation().WithError(e).Fatal("Could not establish connection to postgres")
	}

	connectionPool.SetMaxIdleConns(1)
	connectionPool.SetMaxOpenConns(10)
	db = connectionPool

}
