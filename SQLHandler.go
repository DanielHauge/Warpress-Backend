package main

import (
	"database/sql"
	"fmt"
	"os"
)

func SetupDB() {
	db, err := sql.Open("mysql", os.Args[3])
	if err != nil {
		fmt.Print(err.Error())

	}
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	DB = db
}
