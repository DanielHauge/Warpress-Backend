package BlizzardOauthAPI

import (
	"../../Redis"
	log "../../Utility/Logrus"
	"../../Utility/Monitoring"
	"./BattleNetOauth"
	"context"
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
		RedirectURL:  os.Getenv("HOSTNAME") + "/bnet/auth/callback",
	}
)

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {

	region := r.FormValue("region")
	AuthenticationCFG := *OauthCfg
	AuthenticationCFG.Endpoint = bnet.Endpoint(region)
	oauthState := setStateOauthCookie(w, region)
	AuthUrl := AuthenticationCFG.AuthCodeURL(oauthState)
	http.Redirect(w, r, AuthUrl, http.StatusTemporaryRedirect)
}

func HandleOauthCallback(w http.ResponseWriter, r *http.Request) {

	// Checks if oauthstate from blizzard is correct, in case of hacks and stuff.
	oauthState, region := getStateOauthCookie(r)
	if r.FormValue("state") != oauthState {
		log.WithLocation().Error("Invalid Oauth blizzard state")
		http.Redirect(w, r, "/hov", http.StatusTemporaryRedirect)
		return
	}

	AuthenticationCFG := *OauthCfg
	AuthenticationCFG.Endpoint = bnet.Endpoint(region)
	token, e := AuthenticationCFG.Exchange(context.Background(), r.FormValue("code"))
	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
		return
	}

	// Creates client from token and fetches the users accountId
	authClient := AuthenticationCFG.Client(oauth2.NoContext, token)
	client := bnet.NewClient(region, authClient)
	user, _, e := client.Account().User()
	log.WithField("User", user.ID).WithField("Token", token).Debug("Token")

	Monitoring.LoginInc()

	// Caches the AccessToken in redis for later validation.
	Redis.CacheAccesToken("AT:"+strconv.Itoa(user.ID), token)

	setAccessTokenCookieOnClient(user.ID, region, token, w)

	// If user.id exists in database, fetch data and redirect to login with that pass and accesstoken.
	isRegistered := Redis.DoesKeyExist("MAIN:" + strconv.Itoa(user.ID))
	if isRegistered {
		http.Redirect(w, r, "https://wowhub.io/#/", http.StatusPermanentRedirect)
	} else { // Redirect to register
		http.Redirect(w, r, "https://wowhub.io/#/register", http.StatusPermanentRedirect)
	}

	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
	}
}

func areAccessTokensSame(a oauth2.Token, b oauth2.Token) bool {
	at := a.AccessToken == b.AccessToken
	rt := a.RefreshToken == b.RefreshToken
	tt := a.TokenType == b.TokenType
	return at && rt && tt
}

func DoesUserHaveAccess(w http.ResponseWriter, r *http.Request) (bool, int, string) {
	acesToken, accountId, region, e := getAccessTokenCookieFromClient(r)
	if e != nil {
		log.WithLocation().WithError(e).Warn("User does not have an acceptable cookie")
		return false, 0, ""
	}
	cachedAccessToken, e := Redis.GetAccessToken("AT:" + strconv.Itoa(accountId))
	return areAccessTokensSame(acesToken, cachedAccessToken), accountId, region
}
