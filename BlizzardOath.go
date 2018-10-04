package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/mitchellh/go-bnet"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"time"
)

var oauthCfg = &oauth2.Config{
	// Get from dev.battle.net
	ClientID:     os.Args[1],
	ClientSecret: os.Args[2],

	// Endpoint from this library
	Endpoint: bnet.Endpoint("eu"),
	RedirectURL: "https://localhost/bnet/auth/callback",
}


func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}


func HandleAuthenticate(w http.ResponseWriter, r *http.Request){
	oauthState := generateStateOauthCookie(w)
	AuthUrl := oauthCfg.AuthCodeURL(oauthState)
	http.Redirect(w,r, AuthUrl, http.StatusTemporaryRedirect)
}

func HandleOauthCallback(w http.ResponseWriter, r *http.Request){
	oauthState, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth blizzard state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := oauthCfg.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Errorf("code exchange wrong: %s", err.Error())
		return
	}

	authClient := oauthCfg.Client(oauth2.NoContext, token)

	client := bnet.NewClient("eu", authClient)
	user, _, _ := client.Account().User()
	fmt.Fprint(w, "BattleTag: %s", user.BattleTag)

}
