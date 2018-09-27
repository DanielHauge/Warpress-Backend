package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"database/sql"
)

// go get "github.com/go-sql-driver/mysql"
// go get “github.com/gorilla/mux”
// go get github.com/rs/cors


var DB *sql.DB



func main() {

	router := NewRouter()
	handler := cors.Default().Handler(router)
	IndexPage = SetupIndexPage()
	//SetupDB()
	//defer DB.Close()
	log.Fatal(http.ListenAndServe(":8080", handler))

}

