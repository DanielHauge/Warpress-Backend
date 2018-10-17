package BlizzardOauthAPI

import (
	"../../Prometheus"
	"../../Redis"
	"./BattleNetOauth"
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strconv"
)

var (
	OauthCfg = &oauth2.Config{
		ClientID:     os.Getenv("BNET_CLIENTID"),
		ClientSecret: os.Getenv("BNET_SECRET"),
		Scopes:       []string{"wow.profile"},
		RedirectURL:  "https://localhost:443/bnet/auth/callback",
	}
)

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {

	region := r.FormValue("region")
	AuthenticationCFG := *OauthCfg
	AuthenticationCFG.Endpoint = bnet.Endpoint(region)
	oauthState := SetStateOauthCookie(w, region)
	AuthUrl := AuthenticationCFG.AuthCodeURL(oauthState)
	http.Redirect(w, r, AuthUrl, http.StatusTemporaryRedirect)
}

func HandleOauthCallback(w http.ResponseWriter, r *http.Request) {

	// Checks if oauthstate from blizzard is correct, in case of hacks and stuff.
	oauthState, region := GetStateOauthCookie(r)
	if r.FormValue("state") != oauthState {
		log.Error("invalid oauth blizzard state")
		http.Redirect(w, r, "/hov", http.StatusTemporaryRedirect)
		return
	}

	AuthenticationCFG := *OauthCfg
	AuthenticationCFG.Endpoint = bnet.Endpoint(region)
	token, e := AuthenticationCFG.Exchange(context.Background(), r.FormValue("code"))
	if e != nil {
		log.Error(e)
		return
	}

	// Creates client from token and fetches the users accountId
	authClient := AuthenticationCFG.Client(oauth2.NoContext, token)
	client := bnet.NewClient(region, authClient)
	user, _, e := client.Account().User()
	log.Debug("TOKEN: "+strconv.Itoa(user.ID), token)

	Prometheus.LoginInc()

	// Caches the AccessToken in redis for later validation.
	Redis.CacheAccesToken("AT:"+strconv.Itoa(user.ID), token)

	SetAccessTokenCookieOnClient(user.ID, region, token, w)

	// If user.id exists in database, fetch data and redirect to login with that pass and accesstoken.
	isRegistered := Redis.DoesKeyExist("MAIN:" + strconv.Itoa(user.ID))
	if isRegistered {
		http.Redirect(w, r, "http://localhost:8080/#/Login", http.StatusPermanentRedirect)
	} else { // Redirect to register
		http.Redirect(w, r, "http://localhost:8080/#/Register", http.StatusPermanentRedirect)
	}

	if e != nil {
		log.Error(e)
	}
}

func AreAccessTokensSame(a oauth2.Token, b oauth2.Token) bool {
	at := a.AccessToken == b.AccessToken
	rt := a.RefreshToken == b.RefreshToken
	tt := a.TokenType == b.TokenType
	return at && rt && tt
}

func DoesUserHaveAccess(w http.ResponseWriter, r *http.Request) (bool, int, string) {
	acesToken, accountId, region, e := GetAccessTokenCookieFromClient(r)
	if e != nil {
		log.Error(e)
		return false, 0, ""
	}
	cachedAccessToken, e := Redis.GetAccessToken("AT:" + strconv.Itoa(accountId))
	return AreAccessTokensSame(acesToken, cachedAccessToken), accountId, region
}
