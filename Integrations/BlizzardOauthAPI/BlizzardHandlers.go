package BlizzardOauthAPI

import (
	"../../Redis"
	"./BattleNetOauth"
	"github.com/avelino/slugify"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

var json = jsoniter.ConfigFastest






func GetCharactersForRegistration(w http.ResponseWriter, r *http.Request, id int, region string){

	accesToken, _, _, e := GetAccessTokenCookieFromClient(r)
	if e != nil{
		log.Error(e)
		w.Write([]byte("Something went wrong:"+e.Error()))
		return
	}


	cachedAccessToken, e := Redis.GetAccessToken("AT:"+strconv.Itoa(id))

	if AreAccessTokensSame(accesToken, cachedAccessToken){

		AuthenticationCFG := *OauthCfg
		AuthenticationCFG.Endpoint = bnet.Endpoint(region)

		authClient := AuthenticationCFG.Client(oauth2.NoContext, &accesToken)


		client := bnet.NewClient(region, authClient)
		WowProfile, _, e := client.Profile().WOW()
		if e != nil { log.Error(e) }
		chars := WowProfile.Characters
		sort.Sort(bnet.ByLevel(chars))
		if len(chars) > 4 {
			e = json.NewEncoder(w).Encode(chars[0:5])
		} else {
			e = json.NewEncoder(w).Encode(chars[0:])
		}

		if e != nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(e)
			w.Write([]byte("Unable to parse to json"))
		}
	} else {
		log.Info("User tried to get characters for reg, but was not autherized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("It seems like the credentials are not matching."))
	}
}

func SetMainCharacter(w http.ResponseWriter, r*http.Request, id int, region string){

	//TODO: FY FOR DEN LEDE!. lav noget ordenligt til at læse fra bodien.

		var char CharacterMinimal
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil{
			log.Error(err)
			w.WriteHeader(400)
			w.Write([]byte("Could not read body"))
			return
		}
		if err := r.Body.Close(); err != nil {
			log.Error(err)
		}
		if err := json.Unmarshal(body, &char); err != nil{
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err := json.NewEncoder(w).Encode(err); err != nil{
				log.Error(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		char.Realm = slugify.Slugify(char.Realm)
		char.Region = region
		w.WriteHeader(201)
		Redis.SetStruct("MAIN:"+strconv.Itoa(id), char.ToMap())

}

func GetMainCharacter(w http.ResponseWriter, r *http.Request, id int, region string){

		d, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
		char := CharacterMinimalFromMap(d)
		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(char); if err != nil{ log.Error(err); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}

}