package BlizzardOauthAPI

import (
	"../../Prometheus"
	"../../Redis"
	"./BattleNetOauth"
	"context"
	"crypto/rand"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (

 EuOauthCfg = &oauth2.Config{
	ClientID:     	os.Getenv("BNET_CLIENTID"),
	ClientSecret: 	os.Getenv("BNET_SECRET"),
	Scopes:   		[]string{"wow.profile"},
	Endpoint: 		bnet.Endpoint("eu"),
	RedirectURL: 	"https://localhost:443/bnet/auth/callback",
}

 UsOauthCfg = &oauth2.Config{
	 ClientID:     	os.Getenv("BNET_CLIENTID"),
	 ClientSecret: 	os.Getenv("BNET_SECRET"),
	 Scopes:   		[]string{"wow.profile"},
	 Endpoint: 		bnet.Endpoint("us"),
	 RedirectURL: 	"https://localhost:443/bnet/auth/callback",
 }

)





func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func generateRegionCookie(w http.ResponseWriter, region string){
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "wowhub_recent_region", Value:region, Expires:expiration, Path: "/" }
	http.SetCookie(w, &cookie)
}

func HandleAuthenticate(w http.ResponseWriter, r *http.Request){

	oauthState := generateStateOauthCookie(w)
	var AuthUrl string
	if r.FormValue("region") == "eu"{
		AuthUrl = EuOauthCfg.AuthCodeURL(oauthState)
		generateRegionCookie(w, "eu")
	} else if r.FormValue("region") == "us" {
		AuthUrl = UsOauthCfg.AuthCodeURL(oauthState)
		generateRegionCookie(w, "us")
	}

	http.Redirect(w,r, AuthUrl, http.StatusTemporaryRedirect)
}

func HandleOauthCallback(w http.ResponseWriter, r *http.Request){

	// Checks if oauthstate from blizzard is correct, in case of hacks and stuff.
	oauthState, e := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthState.Value {
		log.Error("invalid oauth blizzard state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) // TODO: redirect til error side eller sådan noget.
		return
	}

	// Find out which region it should handle from:
	region := "eu"
	regionCookie, _ := r.Cookie("wowhub_recent_region")
	region = regionCookie.Value


	// Gets accessToken
	var token *oauth2.Token
	if region == "eu"{
		token, e = EuOauthCfg.Exchange(context.Background(), r.FormValue("code"))
	} else if region == "us"{
		token, e = UsOauthCfg.Exchange(context.Background(), r.FormValue("code"))
	} else {
		token, e = EuOauthCfg.Exchange(context.Background(), r.FormValue("code"))
	}
	if e != nil {
		log.Error(e)
		return
	}

	// Creates client from token and fetches the users accountId
	authClient := EuOauthCfg.Client(oauth2.NoContext, token)
	client := bnet.NewClient(region, authClient)
	user, _, e := client.Account().User()
	log.Debug("TOKEN: " + strconv.Itoa(user.ID) ,token)


	Prometheus.LoginInc()

	// Caches the AccessToken in redis for later validation.
	Redis.CacheAccesToken("AT:"+strconv.Itoa(user.ID), token)

	SetAccessTokenCookieOnClient(user.ID, token, w)


	// If user.id exists in database, fetch data and redirect to login with that pass and accesstoken.
	isRegistered := Redis.DoesKeyExist("MAIN:"+strconv.Itoa(user.ID))
	if isRegistered {
		http.Redirect(w,r, "http://localhost:8080/#/Login", http.StatusPermanentRedirect)
	} else { // Redirect to register
		http.Redirect(w,r, "http://localhost:8080/#/Register", http.StatusPermanentRedirect)
	}

	if e != nil {
		log.Error(e)
	}
}

func AreAccessTokensSame(a oauth2.Token, b oauth2.Token)bool{
	at := a.AccessToken == b.AccessToken
	rt := a.RefreshToken == b.RefreshToken
	tt := a.TokenType == b.TokenType
	return at && rt && tt
}

func DoesUserHaveAccess(w http.ResponseWriter, r *http.Request) (bool, int) {
	accesToken, accountid, e := GetAccessTokenCookieFromClient(r)
	if e != nil{
		log.Error(e)
		w.WriteHeader(500)
		return false, 0
	}
	cachedAccessToken, e := Redis.GetAccessToken("AT:"+strconv.Itoa(accountid))
	return AreAccessTokensSame(accesToken, cachedAccessToken), accountid
}






