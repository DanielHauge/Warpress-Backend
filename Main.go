package main

import (
	"crypto/tls"
	"database/sql"
	"github.com/rs/cors"
	"log"
	"net/http"
)

// go get github.com/go-sql-driver/mysql
// go get github.com/gorilla/mux
// go get github.com/rs/cors
// go get gopkg.in/russross/blackfriday.v2
// go get -u github.com/go-redis/redis
// go get github.com/gorilla/securecookie


var DB *sql.DB

// 1. BNET_CLIENTID
// 2. BNET_SECRET
// 3. CONNECTION_STRING
// 4. APIKEY

func main() {

	// Router
	router := NewRouter()
	handler := cors.Default().Handler(router)
	IndexPage = SetupIndexPage()

	// DB Stuff
	//SetupDB()
	//defer DB.Close()

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         ":443",
		Handler:      handler,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}


	//Start Server
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))


}

