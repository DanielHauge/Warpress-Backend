package main

import (
	"./Blizzard"
	"./GoBnet"
	"encoding/json"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

var ApiURL = "https://eu.api.battle.net/wow/"

type charRequest struct {
	Name string `json:"name"`
	Realm string `json:"realm"`
	Locale string `json:"locale"`
}

func GetCharactersForRegistration(w http.ResponseWriter, r *http.Request){

	accesToken, accountid, e := GetAccessTokenCookieFromClient(r)
	if e != nil{
		w.Write([]byte("Something went wrong:"+e.Error()))
		return
	}
	cachedAccessToken, e := GetAccessToken(accountid)

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

		var char charRequest
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
		w.WriteHeader(200)
		SetMainChar(id, char)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetMainCharacter(w http.ResponseWriter, r *http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {
		d, e := GetMainChar(id)
		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Println(e.Error())
		} else {
			msg, err := json.Marshal(d); if err != nil{ log.Println(err); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DoesUserHaveAccess(w http.ResponseWriter, r *http.Request) (bool, int) {
	accesToken, accountid, e := GetAccessTokenCookieFromClient(r)
	if e != nil{
		w.Write([]byte("Something went wrong: "+e.Error()))
		w.WriteHeader(500)
		return false, 0
	}
	cachedAccessToken, e := GetAccessToken(accountid)
	return AreAccessTokensSame(accesToken, cachedAccessToken), accountid
}

func GetFullCharInfo(w http.ResponseWriter, r *http.Request){
	acces, _ := DoesUserHaveAccess(w, r)
	if acces {

		var char charRequest
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


		fullChar, e := getFullCharInfo(char)

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Println(e.Error())
		} else {
			msg, err := json.Marshal(fullChar); if err != nil{ log.Println(err); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}



	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func getFullCharInfo(char charRequest) (Blizzard.FullCharInfo, error){

	url := ApiURL+"/character/"+char.Realm+"/"+char.Name+"?locale="+char.Locale+"apikey="+os.Args[4]
	resp, e := http.Get(url)

	var fullChar Blizzard.FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(fullChar)
	if e != nil { log.Println(e.Error()) }

	return fullChar, e
}