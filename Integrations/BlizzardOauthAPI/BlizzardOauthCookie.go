package BlizzardOauthAPI

import (
	log "../../Utility/Logrus"
	"crypto/rand"
	"encoding/base64"
	"github.com/gorilla/securecookie"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"time"
)

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

func setAccessTokenCookieOnClient(accountId int, region string, token *oauth2.Token, w http.ResponseWriter) {
	tokenAsMap := map[string]string{
		"region":       region,
		"accountId":    strconv.Itoa(accountId),
		"expire":       token.Expiry.Format(time.RFC3339),
		"tokentype":    token.TokenType,
		"refreshtoken": token.RefreshToken,
		"accesstoken":  token.AccessToken,
	}
	encoded, err := s.Encode("WowHubAccessToken", tokenAsMap)
	if err == nil {
		cookie := &http.Cookie{
			Name:     "WowHubAccessToken",
			Value:    encoded,
			Expires:  token.Expiry,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		}
		log.WithField("User", accountId).Debug("Setting AccessToken")
		http.SetCookie(w, cookie)
	} else {
		log.WithLocation().WithError(err).Error("Hov!")
	}
}

func getAccessTokenCookieFromClient(r *http.Request) (oauth2.Token, int, string, error) {
	cookie, err := r.Cookie("WowHubAccessToken")
	if err == nil {
		value := make(map[string]string)
		if err = s.Decode("WowHubAccessToken", cookie.Value, &value); err == nil {
			time, err := time.Parse(time.RFC3339, value["expire"])
			token := oauth2.Token{
				Expiry:       time,
				TokenType:    value["tokentype"],
				RefreshToken: value["refreshtoken"],
				AccessToken:  value["accesstoken"],
			}
			log.WithField("User", value["accountId"]).Debug("Getting AccessToken")
			aid, err := strconv.Atoi(value["accountId"])

			return token, aid, value["region"], err
		}
	}
	return oauth2.Token{}, 0, "", err
}

func DeleteCookies(w http.ResponseWriter){
	at := &http.Cookie{
		Name:     "WowHubAccessToken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	oauth := &http.Cookie{
		Name:     "oauthstate",
		Value:    "",
		Expires:  time.Unix(0,0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/bnet/auth/callback",
	}

	http.SetCookie(w, at)
	http.SetCookie(w, oauth)
}

func setStateOauthCookie(w http.ResponseWriter, region string) string {
	var expiration = time.Now().Add(30 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	StateOauthCookieMap := map[string]string{
		"oauthstate": state,
		"region":     region,
	}
	encoded, err := s.Encode("oauthstate", StateOauthCookieMap)
	if err == nil {
		cookie := http.Cookie{
			Name:     "oauthstate",
			Value:    encoded,
			Expires:  expiration,
			HttpOnly: true,
			Secure:   true,
			Path:     "/bnet/auth/callback",
		}
		http.SetCookie(w, &cookie)
	}
	return state
}

func getStateOauthCookie(r *http.Request) (string, string) {
	cookie, err := r.Cookie("oauthstate")
	if err == nil {
		value := make(map[string]string)
		if err = s.Decode("oauthstate", cookie.Value, &value); err == nil {
			log.WithField("Map", value).Debug("Got OauthState cookie")
			return value["oauthstate"], value["region"]
		} else {
			log.WithLocation().WithError(err).Error("Hov! Cannot decode Oauthstate cookie")
		}
	} else {
		log.WithLocation().WithError(err).Error("Hov! Cannot get Oauthstate cookie")
		return "", ""
	}
	return "", ""
}
