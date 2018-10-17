package BlizzardOauthAPI

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gorilla/securecookie"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"time"
)

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

func SetAccessTokenCookieOnClient(accountId int, region string, token *oauth2.Token, w http.ResponseWriter) {
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
		log.Debug("Setting AccessToken: " + strconv.Itoa(accountId))
		http.SetCookie(w, cookie)
	} else {
		log.Error(err)
	}
}

func GetAccessTokenCookieFromClient(r *http.Request) (oauth2.Token, int, string, error) { // TODO: When application starts, new key is generated, and therefor needs to ask for new accessToken from blizzard.
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
			log.Debug("Getting AccessToken: " + value["accountId"])
			aid, err := strconv.Atoi(value["accountId"])

			return token, aid, value["region"], err
		}
	}
	return oauth2.Token{}, 0, "", err
}

func SetStateOauthCookie(w http.ResponseWriter, region string) string {
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

func GetStateOauthCookie(r *http.Request) (string, string) {
	cookie, err := r.Cookie("oauthstate")
	if err == nil {
		value := make(map[string]string)
		if err = s.Decode("oauthstate", cookie.Value, &value); err == nil {
			log.Debug("Got OauthState cookie: ", value)
			return value["oauthstate"], value["region"]
		} else {
			log.Error(err, " -> Occured in decoding stateOauth")
		}
	} else {
		log.Error(err, " -> Occured in getting cookie")
		return "", ""
	}
	return "", ""
}
