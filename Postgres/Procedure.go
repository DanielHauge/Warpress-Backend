package Postgres

import (
	log "../Utility/Logrus"
	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"reflect"
)

func QuerySingle(query string, args []interface{}, obj ...interface{}) error {
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		log.WithLocation().WithError(err).WithField("Query", query).Error("Could not prepare statement")
	}
	err = statement.QueryRow(args...).Scan(obj...)
	if err != nil {
		if err == sql.ErrNoRows {
			log.WithField("Query", query).WithField("Struct", obj).WithField("Arguments", args).Info("Could not find any results")
		} else {
			log.WithField("Query", query).WithField("Struct", obj).WithField("Arguments", args).WithError(err).Error("Could not map struct from row")
		}
	}

	log.WithFields(logrus.Fields{
		"Postgres":  "Query Single",
		"Query":     query,
		"Arguments": args,
		"Struct":    obj,
	}).Info("Postgress query")

	return err
}

// TODO: Debug and fix this.
func QueryMultiple(query string, args []interface{}, obj interface{}) error {
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		log.WithLocation().WithError(err).WithField("Query", query).Error("Could not prepare statement")
	}

	rows, err := statement.Query(args...)
	if err == sql.ErrNoRows {
		log.WithField("Query", query).WithField("Struct", obj).WithField("Arguments", args).Info("Could not find any results")
		return err
	}
	defer rows.Close()

	ty := reflect.TypeOf(obj).Elem()

	results := reflect.MakeSlice(reflect.TypeOf(obj).Elem(), 0, 500)
	for rows.Next() {

		single := reflect.New(ty.Elem())
		var fieldPointers []interface{}

		for i := 0; i < single.Elem().NumField(); i++ {
			fieldPointers = append(fieldPointers, single.Elem().Field(i).Addr().Interface())
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			log.WithLocation().WithField("Query", query).WithField("Struct", single).WithField("Arguments", args).WithError(err).Error("Could not map struct from row")
		}
		results = reflect.Append(results, single.Elem())

	}
	log.WithFields(logrus.Fields{
		"Postgres":  "Query Multiple",
		"Query":     query,
		"Arguments": args,
		"Struct":    obj,
	}).Info("Postgress query")

	err = copier.Copy(obj, results.Interface())

	return err
}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		log.WithLocation().WithError(err).WithField("Query", query).Error("Could not prepare statement")
	}

	res, err := statement.Exec(args...)
	if err != nil {
		log.WithLocation().WithError(err).WithField("Query", query).Error("Could not Execute statement")
	}
	log.WithFields(logrus.Fields{
		"Postgres":  "Execute",
		"Query":     query,
		"Arguments": args,
		"Result":    res,
	}).Info("Postgress query")

	return res, err
}
