package main

import (
	"./GoBnet"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"sort"
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
	user, _, e := client.Account().User()
	CacheAccesToken(user.ID, token)

	// If user.id exists in database, fetch data and redirect to login with that pass and accesstoken.
	isRegistered := IsUserRegistered(user.ID)
	if isRegistered {
		// Gad vide om man kan skrive til responsewriteren og fange det i fronteden som json?
		http.Redirect(w,r, "http://localhost:8080/#/Login/"+string(user.ID), http.StatusPermanentRedirect)
	} else { // Redirect to register
		WowProfile, _, e := client.Profile().WOW()
		if e != nil { log.Println(e.Error()) }
		chars := WowProfile.Characters
		sort.Sort(bnet.ByLevel(chars))
		//bytes, e := json.Marshal(chars[0:4])
		coockie := http.Cookie{Name: "WarpressAccessToken", Value: string(user.ID), Expires: time.Now().Add(time.Hour)}

		http.SetCookie(w, &coockie)
		http.Redirect(w,r, "http://localhost:8080/#/Register", http.StatusFound)
	}



	if e != nil {
		log.Println(e.Error())
		fmt.Fprint(w, "Something went wrong!", e.Error())
	}


}



