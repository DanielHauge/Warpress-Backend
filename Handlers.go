package main

import (
	"./Integrations/BlizzardOauthAPI"
	"./Integrations/BlizzardOpenAPI"
	"./Integrations/Raider.io"
	"./Integrations/WarcraftLogs"
	"./Personal"
	"./Redis"
	"bytes"
	log "github.com/sirupsen/logrus"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
	"reflect"
	"strconv"
)


func SetupIndexPage()[]byte{
	var buffer bytes.Buffer

	for _, v := range routes {
		buffer.WriteString("\n\n### "+v.Name+"\n")
		buffer.WriteString("##### Route: "+v.Pattern+"\n")
		buffer.WriteString("##### Method: "+v.Method+"\n")


		if v.ExpectedInput != nil {
			buffer.WriteString("##### Input: \n")

			inputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedInput))
			numOfInputFields := inputFields.Type().NumField()

			for i := 0; i < numOfInputFields; i++ {
				buffer.WriteString("- ")
				buffer.WriteString(inputFields.Type().Field(i).Name)
				buffer.WriteString(" : ")
				buffer.WriteString(inputFields.Type().Field(i).Type.Kind().String() + "\n\n")
			}
		}

		if v.ExpectedOutput != nil {
			buffer.WriteString("##### Output: \n")
			outputFields := reflect.Indirect(reflect.ValueOf(v.ExpectedOutput))
			numOfOutputFields := outputFields.Type().NumField()

			for i := 0; i < numOfOutputFields; i++ {
				buffer.WriteString("- ")
				buffer.WriteString(outputFields.Type().Field(i).Name)
				buffer.WriteString(" : ")
				buffer.WriteString(outputFields.Type().Field(i).Type.Kind().String() + "\n")
			}
		}

		buffer.WriteString("\n##### Example:\n")
		buffer.WriteString("- Input:\n")
		var b []byte
		if v.ExpectedInput != nil{
			b, _ = json.Marshal(v.ExpectedInput)
		} else {
			b = []byte("Nothing")
		}
		buffer.WriteString(string(b)+"\n")

		buffer.WriteString("\n- Output:\n")
		if v.ExpectedOutput != nil {
			b, _ = json.Marshal(v.ExpectedOutput)
		} else {
			b = []byte("Nothing")
		}
		buffer.WriteString(string(b)+"\n")
	}
	return buffer.Bytes()
}

var IndexPage []byte

func Index(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("# Wowhub API\n")
	buffer.WriteString("This is a api for the website of wowhub, this page is only available during development\n\n")
	buffer.WriteString("## Api endpoints:\n\n")
	buffer.Write(IndexPage)
	output := blackfriday.Run([]byte(buffer.Bytes()))
	w.Write(output)
}

func HandleGetPersonalFull(w http.ResponseWriter, r *http.Request) {

	acces, id := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
	if acces {
		var Profile Personal.PersonalProfile
		key := "PERSONAL:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){

			e = Redis.CacheGetResult(key, &Profile)
			go func() {
				var Caching Personal.PersonalProfile
				Personal.FetchFullPersonal(id, &Caching)
				Redis.CacheSetResult(key, Caching)
			}()

		} else {

			e = Personal.FetchFullPersonal(id, &Profile)
			if e == nil {
				go Redis.CacheSetResult(key, Profile)
			}

		}


		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
		} else {
			msg, err := json.Marshal(Profile); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get full personal, but was not autherized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}

}

func HandleGetPersonalRaiderio(w http.ResponseWriter, r *http.Request){
	acces, id := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
	if acces {

		var RaiderioProfile Raider_io.CharacterProfile
		key := "PERSONAL/RAIDERIO:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){

			e = Redis.CacheGetResult(key, &RaiderioProfile)
			go func(){
				var Caching Raider_io.CharacterProfile
				Personal.FetchRaiderioPersonal(id, &Caching)
				Redis.CacheSetResult(key, Caching)
			}()

		} else {

			e = Personal.FetchRaiderioPersonal(id, &RaiderioProfile)
			if e == nil{
				go Redis.CacheSetResult(key, RaiderioProfile)
			}

		}

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(RaiderioProfile); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal raiderio, but was not autherized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}
}

func HandleGetPersonalWarcraftLogs(w http.ResponseWriter, r *http.Request){
	acces, id := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
	if acces {




		var logs []WarcraftLogs.Encounter
		key := "PERSONAL/LOGS:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){

			e = Redis.CacheGetResult(key, &logs)
			go func(){
				var Caching []WarcraftLogs.Encounter
				Personal.FetchWarcraftlogsPersonal(id, &Caching)
				Redis.CacheSetResult(key, Caching)
			}()

		} else {

			e = Personal.FetchWarcraftlogsPersonal(id, &logs)
			if e == nil{
				go Redis.CacheSetResult(key, logs)
			}
		}

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(logs); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal warcraftlogs, but was not autherized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}
}

func HandleGetPersonalBlizzardChar(w http.ResponseWriter, r *http.Request){
	acces, id := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
	if acces {

		var blizzProfile BlizzardOpenAPI.FullCharInfo
		key := "PERSONAL/BLIZZARD:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){

			e = Redis.CacheGetResult(key, &blizzProfile)
			go func(){
				var Caching BlizzardOpenAPI.FullCharInfo
				Personal.FetchBlizzardPersonal(id, &Caching)
				Redis.CacheSetResult(key, Caching)
			}()

		} else {

			e = Personal.FetchBlizzardPersonal(id, &blizzProfile)
			if e == nil{
				go Redis.CacheSetResult(key, blizzProfile)
			}
		}

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(blizzProfile); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal blizzard, but was not autherized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}
}

func HandleGetPersonalImprovements(w http.ResponseWriter, r *http.Request){
	acces, id := BlizzardOauthAPI.DoesUserHaveAccess(w, r)
	if acces {

		var improvementsProfile Personal.PersonalImprovement
		key := "PERSONAL/IMPROVEMENT:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){

			e = Redis.CacheGetResult(key, &improvementsProfile)
			go func(){
				var Caching Personal.PersonalImprovement
				Personal.FetchPersonalImprovementsFull(id, &Caching)
				Redis.CacheSetResult(key, Caching)
			}()

		} else {

			e = Personal.FetchPersonalImprovementsFull(id, &improvementsProfile)
			if e == nil{
				go Redis.CacheSetResult(key, improvementsProfile)
			}
		}

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(improvementsProfile); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal improvements, but was not authorized")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}
}