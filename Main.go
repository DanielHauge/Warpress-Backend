package main

import (
	"./Redis"
	"crypto/tls"
	"database/sql"
	"github.com/json-iterator/go"
	"github.com/kz/discordrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// go get github.com/go-sql-driver/mysql
// go get github.com/gorilla/mux
// go get github.com/rs/cors
// go get gopkg.in/russross/blackfriday.v2
// go get -u github.com/go-redis/redis
// go get github.com/gorilla/securecookie
// go get golang.org/x/oauth2
// go get github.com/avelino/slugify
// go get github.com/brianvoe/gofakeit
// go get github.com/json-iterator/go
// go get -u github.com/go-redis/cache
// go get github.com/vmihailenco/msgpack
// go get github.com/prometheus/client_golang/prometheus
// go get github.com/sirupsen/logrus
// go get -u github.com/kz/discordrus
// go get github.com/jinzhu/copier


// Unsure but might need:
// go get golang.org/x/sys/windows/svc/eventlog
// go get gopkg.in/alecthomas/kingpin.v2


var DB *sql.DB
var json = jsoniter.ConfigFastest

// 1. BNET_CLIENTID
// 2. BNET_SECRET
// 3. CONNECTION_STRING
// 4. APIKEY
// 5. Public warcraftlogs
// 6. private warcraftlogs


func init(){
	prometheus.MustRegister(promRequests, promLogins)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		TimestampFormat: time.RFC822,
	})

	log.SetOutput(os.Stderr)



	log.AddHook(discordrus.NewHook(

		os.Getenv("DISCORDRUS_WEBHOOK_URL"),
		log.WarnLevel,
		&discordrus.Opts{
			Username: "Logrus",
			EnableCustomColors: true,
			CustomLevelColors: &discordrus.LevelColors{
				Debug: 10170623,
				Info:  3581519,
				Warn:  14327864,
				Error: 13631488,
				Panic: 13631488,
				Fatal: 13631488,
			},
			DisableInlineFields: false,
		},
		))
}


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
		log.Warn("Cannot connect to database. Make sure redis is running.")
		log.Fatal(e)
	}
	//Start Server
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))


}

