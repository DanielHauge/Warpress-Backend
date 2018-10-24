package Postgres

import (
	log "../Utility/Logrus"
)

func Statement(query string, args ...interface{}) {
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		log.WithLocation().WithError(err).WithField("Query", query).Fatal("Could not prepare statement")
	}

	rows, err := statement.Query(args)
}
