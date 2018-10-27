package BlizzardOauthAPI

import (
	"../../Postgres"
	"../../Redis"
	"../../Utility/HttpHelper"
	log "../../Utility/Logrus"
	"./BattleNetOauth"
	"github.com/avelino/slugify"
	"github.com/jinzhu/copier"
	"github.com/json-iterator/go"
	"golang.org/x/oauth2"
	"net/http"
	"sort"
	"strconv"
)

var json = jsoniter.ConfigFastest

func GetCharactersForRegistration(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("CHARS:", id, wowCharacters{}, MakeFetcherFunction(region, WowCharacters))

	result := <- channel

	if result.Error == nil{
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("was not able to marshal chars")
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.WithLocation().WithError(result.Error).Error("How!")
		w.WriteHeader(500)
		w.Write([]byte(result.Error.Error()))
	}


}

type wowCharacters struct {
	Chars []bnet.WOWCharacter `json:"chars"`
}

func WowCharacters(id int, region string) (wowCharacters, error) {

	log.WithField("ID", id).WithField("Region", region).Info("BlizzardOauth2 Fetching account characters")
	accesToken, e := Redis.GetAccessToken("AT:" + strconv.Itoa(id))

	AuthenticationCFG := *OauthCfg
	AuthenticationCFG.Endpoint = bnet.Endpoint(region)

	authClient := AuthenticationCFG.Client(oauth2.NoContext, &accesToken)

	client := bnet.NewClient(region, authClient)
	WowProfile, _, e := client.Profile().WOW()
	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
	}

	chars := WowProfile.Characters
	sort.Sort(bnet.ByLevel(chars))
	if len(chars) > 4 {
		return wowCharacters{chars[0:5]}, e
	} else {
		return wowCharacters{chars[0:]}, e
	}
	return wowCharacters{WowProfile.Characters}, e

}

func MakeFetcherFunction(region string, fetcher func(id int, region string) (wowCharacters, error)) func(id int, obj *interface{}) error {
	return func(id int, obj *interface{}) error {
		log.Info("")
		wowchars, e := fetcher(id, region)
		copier.Copy(obj, wowchars)
		return e
	}
}

func SetMainCharacter(w http.ResponseWriter, r *http.Request, id int, region string) {

	ids := strconv.Itoa(id)
	w.WriteHeader(201)
	Redis.DeleteKey("MAIN:"+ids, "PERSONAL:"+ids, "PERSONAL/RAIDERIO:"+ids, "PERSONAL/LOGS:"+ids, "PERSONAL/BLIZZARD:"+ids, "PERSONAL/IMPROVEMENT:"+ids, "GUILD/OVERVIEW:"+ids, "GUILD:"+ids)

	var char CharacterMinimal
	HttpHelper.ReadFromRequest(w, r, &char)

	char.Realm = slugify.Slugify(char.Realm)
	char.Region = region

	Postgres.SetMain(id, char.Name, char.Realm, char.Region)
}

func GetMainCharacter(w http.ResponseWriter, r *http.Request, id int, region string) {

	nam, real, regio, e := Postgres.GetMain(id)
	log.Info(nam, real, regio, "This is the values")
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		w.WriteHeader(500)
		w.Write([]byte("Hov!, there was not main detected!"))
	} else {

		char := CharacterMinimal{nam, real, regio}

		msg, err := json.Marshal(char)
		if err != nil {
			log.WithLocation().WithError(err).Error("Hov!")
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(200)
		w.Write(msg)
	}

}
