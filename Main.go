package main

import (
	log "./Logrus"
	"./Redis"
	"crypto/tls"
	"github.com/json-iterator/go"
	"github.com/rs/cors"
	"net/http"
	"os"
)

// IMPORTS!
/*
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
// go get github.com/wawandco/fako
// go get golang.org/x/crypto/acme/autocert

// go get github.com/swaggo/swag/cmd/swag
// go get github.com/swaggo/http-swagger

// Unsure but might need:
// go get golang.org/x/sys/windows/svc/eventlog
// go get gopkg.in/alecthomas/kingpin.v2
*/

var json = jsoniter.ConfigFastest

// TODO: Setup nginx reverse proxy
// TODO: Setup service for API and make that running

// TODO: Logout

// TODO: Make code clean and sleak
// TODO: Use real structs in examples, but randomize it. Fake it, mock it -> pretty fucked up with to much data.
// TODO: Make documentation for API better.

func main() {

	router := NewRouter()
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler(router)
	IndexPage = SetupIndexPage()

	if e := Redis.CanIConnect(); e != nil {
		log.WithLocation().WithError(e).Fatal("Cannot connect to database. Make sure redis is running.")
	}

	if os.Getenv("DEBUG") == "true" {
		/*
			certManager := autocert.Manager{
				Prompt:autocert.AcceptTOS,
				HostPolicy:autocert.HostWhitelist(os.Getenv("HOSTNAME")),
				Cache:autocert.DirCache("certs"),
			}
		*/

		cfg := &tls.Config{
			//GetCertificate:certManager.GetCertificate,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			},
		}

		srv := &http.Server{
			Addr:         ":https",
			Handler:      handler,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

		//go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

		//Start Server
		log.Fatal(srv.ListenAndServeTLS(os.Getenv("CERT_PUBLIC"), os.Getenv("CERT_PRIVATE")))
	} else {
		log.Fatal(http.ListenAndServe(":http", handler))
	}

}
