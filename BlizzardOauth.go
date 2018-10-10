package main

import (
	"./GoBnet"
	"./Redis"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var oauthCfg = &oauth2.Config{
	ClientID:     	os.Args[1],
	ClientSecret: 	os.Args[2],
	Scopes:   		[]string{"wow.profile"},
	Endpoint: 		bnet.Endpoint("eu"),
	RedirectURL: 	"https://localhost:443/bnet/auth/callback",
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
	// Checks if oauthstate from blizzard is correct, in case of hacks and stuff.
	oauthState, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth blizzard state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) // TODO: redirect til error side eller sådan noget.
		return
	}

	// Gets accessToken
	token, err := oauthCfg.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Errorf("code exchange wrong: %s", err.Error())
		return
	}

	// Creates client from token and fetches the users accountId
	authClient := oauthCfg.Client(oauth2.NoContext, token)
	client := bnet.NewClient("eu", authClient)
	user, _, e := client.Account().User()
	log.Println(user.ID)
	log.Print("TOKEN: ")
	log.Println(token)


	// Caches the AccessToken in redis for later validation.
	Redis.CacheAccesToken("AT:"+strconv.Itoa(user.ID), token)

	SetAccessTokenCookieOnClient(user.ID, token, w)


	// If user.id exists in database, fetch data and redirect to login with that pass and accesstoken.
	isRegistered := Redis.DoesKeyExist("AT:"+strconv.Itoa(user.ID))
	if isRegistered {
		http.Redirect(w,r, "http://localhost:8080/#/Login", http.StatusPermanentRedirect)
	} else { // Redirect to register
		http.Redirect(w,r, "http://localhost:8080/#/Register", http.StatusPermanentRedirect)
	}

	if e != nil {
		log.Println(e.Error())
		fmt.Fprint(w, "Something went wrong!", e.Error())
	}
}

func AreAccessTokensSame(a oauth2.Token, b oauth2.Token)bool{
	at := a.AccessToken == b.AccessToken
	rt := a.RefreshToken == b.RefreshToken
	tt := a.TokenType == b.TokenType
	return at && rt && tt
}







