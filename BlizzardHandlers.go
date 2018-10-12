package main

import (
	"./Blizzard"
	"./GoBnet"
	"./Redis"
	"github.com/avelino/slugify"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)






func GetCharactersForRegistration(w http.ResponseWriter, r *http.Request){

	accesToken, accountid, e := GetAccessTokenCookieFromClient(r)
	if e != nil{
		w.Write([]byte("Something went wrong:"+e.Error()))
		return
	}
	cachedAccessToken, e := Redis.GetAccessToken("AT:"+strconv.Itoa(accountid))

	if AreAccessTokensSame(accesToken, cachedAccessToken){
		authClient := oauthCfg.Client(oauth2.NoContext, &accesToken)
		client := bnet.NewClient("eu", authClient)
		WowProfile, _, e := client.Profile().WOW()
		if e != nil { log.Println(e.Error()) }
		chars := WowProfile.Characters
		sort.Sort(bnet.ByLevel(chars))
		e = json.NewEncoder(w).Encode(chars[0:4])
		if e != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to parse to json"))
		}
	} else {
		w.WriteHeader(401)
		w.Write([]byte("It seems like the credentials are not matching."))
	}
}

func SetMainCharacter(w http.ResponseWriter, r*http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {

		var char Blizzard.CharacterMinimal
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil{
			log.Println(err)
			w.WriteHeader(400)
			w.Write([]byte("Could not read body"))
			return
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &char); err != nil{
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err := json.NewEncoder(w).Encode(err); err != nil{
				panic(err)
			}
		}
		char.Realm = slugify.Slugify(char.Realm)
		w.WriteHeader(201)
		Redis.SetStruct("MAIN:"+strconv.Itoa(id), char.ToMap())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetMainCharacter(w http.ResponseWriter, r *http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {
		d, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
		char := Blizzard.CharacterMinimalFromMap(d)
		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Println(e.Error())
		} else {
			msg, err := json.Marshal(char); if err != nil{ log.Println(err); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func FromLocaleToRegion(locale string) string{
	switch locale {
	case "en_GB":
		return "EU"
	case "de_DE":
		return "EU"
	case "es_ES":
		return "EU"
	case "fr_FR":
		return "EU"
	case "it_IT":
		return "EU"
	case "pl_PL":
		return "EU"
	case "pt_PT":
		return "EU"
	case "ru_RU":
		return "EU"
	case "en_US":
		return "US"
	case "pt_BR":
		return "US"
	case "es_MX":
		return "US"
	case "zh_TW":
		return "TW"
	case "ko_KR":
		return "KR"
	default:
		return "EU"
	}
}