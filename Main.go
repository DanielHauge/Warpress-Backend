package main

import (
	"./Redis"
	"crypto/tls"
	"database/sql"
	"github.com/json-iterator/go"
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
// go get github.com/avelino/slugify
// go get github.com/brianvoe/gofakeit
// go get github.com/json-iterator/go
// go get -u github.com/go-redis/cache
// go get github.com/vmihailenco/msgpack
// go get github.com/prometheus/client_golang/prometheus


var DB *sql.DB
var json = jsoniter.ConfigFastest

// 1. BNET_CLIENTID
// 2. BNET_SECRET
// 3. CONNECTION_STRING
// 4. APIKEY
// 5. Public warcraftlogs
// 6. private warcraftlogs




func main() {

	router := NewRouter()
	handler := cors.Default().Handler(router)
	IndexPage = SetupIndexPage()

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

	if e := Redis.CanIConnect(); e != nil{
		log.Println("Cannot connect to database. Make sure redis is running.")
		log.Fatal(e)
	}
	//Start Server
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))


}

