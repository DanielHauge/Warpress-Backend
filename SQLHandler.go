package main

import (
	"database/sql"
	"os"
	"fmt"
)

func SetupDB(){
	db, err := sql.Open("mysql", os.Args[1]+":"+os.Args[2]+"@tcp("+os.Args[3]+":3306)/HackerNewsDB?parseTime=True")
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
