package main

import (
	"./Redis"
	log "./Utility/Logrus"
	"crypto/tls"
	"github.com/json-iterator/go"
	"github.com/rs/cors"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest

// LatentTODO: Setup service for API and make that running

// TODO: Fanger ikke warcraft reports? Done?


func main() {

	router := NewRouter()
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://wowhub.io"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler(router)
	IndexPage = SetupIndexPage()

	if e := Redis.CanIConnect(); e != nil {
		log.WithLocation().WithError(e).Fatal("Cannot connect to database. Make sure redis is running.")
	}

	if os.Getenv("DEBUG") == "true" {
		StartLocalhost(handler)
	} else {
		StartProduction(handler)
	}

}

func StartLocalhost(handler http.Handler){
	cfg := &tls.Config{
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
	//Start Server
	log.Info("Starting Developer server")
	log.Fatal(srv.ListenAndServeTLS(os.Getenv("CERT_PUBLIC"), os.Getenv("CERT_PRIVATE")))
}

func StartProduction(handler http.Handler){
	log.Info("Starting production server")
	log.Fatal(http.ListenAndServe(":http", handler))
}
