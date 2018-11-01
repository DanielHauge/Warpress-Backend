package Postgres

import (
	log "../Utility/Logrus"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"time"
)

var db *sql.DB

func init() {
	connectionPool, err := sql.Open("postgres", "postgres://"+os.Getenv("DBPASS")+"@"+os.Getenv("CONNECTION_POSTGRES")+"/admin?sslmode=disable")
	if err != nil {
		log.WithLocation().WithError(err).Fatal("Could not establish connection pool for postgres")
	}
	e := connectionPool.Ping()
	if e != nil {
		log.WithLocation().WithError(e).Fatal("Could not establish connection to postgres")
	}

	connectionPool.SetMaxIdleConns(1)
	connectionPool.SetMaxOpenConns(5)
	connectionPool.SetConnMaxLifetime(time.Hour)
	db = connectionPool

}
