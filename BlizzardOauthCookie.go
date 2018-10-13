package main

import (
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


func SetAccessTokenCookieOnClient(accountId int, token *oauth2.Token, w http.ResponseWriter) {
	tokenAsMap := map[string]string{
		"accountId": strconv.Itoa(accountId),
		"expire":token.Expiry.Format(time.RFC3339),
		"tokentype":token.TokenType,
		"refreshtoken":token.RefreshToken,
		"accesstoken":token.AccessToken,
	}
	encoded, err := s.Encode("WarpressAccessToken", tokenAsMap)
	if err == nil{
		cookie := &http.Cookie{
			Name: "WarpressAccessToken",
			Value: encoded,
			Expires: token.Expiry,
			Path: "/",
			HttpOnly:true,
		}
		log.Debug("Setting AccessToken: "+strconv.Itoa(accountId))
		http.SetCookie(w, cookie)
	} else {
		log.Error(err)
	}
}

func GetAccessTokenCookieFromClient(r *http.Request) (oauth2.Token, int,error) { // TODO: When application starts, new key is generated, and therefor needs to ask for new accessToken from blizzard.
	cookie, err := r.Cookie("WarpressAccessToken")
	if err == nil{
		value := make(map[string]string)
		if err = s.Decode("WarpressAccessToken", cookie.Value, &value); err == nil{
			time, err := time.Parse(time.RFC3339, value["expire"])
			token := oauth2.Token{
				Expiry: time,
				TokenType: value["tokentype"],
				RefreshToken: value["refreshtoken"],
				AccessToken: value["accesstoken"],
			}
			log.Debug("Getting AccessToken: "+value["accountId"])
			aid, err := strconv.Atoi(value["accountId"])

			return token, aid, err
		}
	}
	return oauth2.Token{}, 0, err
}
