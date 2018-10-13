package main

import (
	"./Blizzard"
	"./Raider.io"
	"./Redis"
	"./WarcraftLogs"
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
	buffer.WriteString("# Warpress API\n")
	buffer.WriteString("This is a api for the website of Warpress, this page is only available during development\n\n")
	buffer.WriteString("## Api endpoints:\n\n")
	buffer.Write(IndexPage)
	output := blackfriday.Run([]byte(buffer.Bytes()))
	w.Write(output)
}


func GetPersonalFull(w http.ResponseWriter, r *http.Request) {

	acces, id := DoesUserHaveAccess(w, r)
	if acces {
		var Profile PersonalProfile
		key := "PERSONAL:"+strconv.Itoa(id)
		var e error
		if Redis.DoesKeyExist(key){
			e = Redis.CacheGetResult(key, &Profile)
			go func() {
				var Caching PersonalProfile
				FetchFullPersonal(id, &Caching)
				Redis.CacheSetResult("PERSONAL:"+strconv.Itoa(id), Caching)
			}()
		} else {
			e = FetchFullPersonal(id, &Profile)
			if e == nil {
				go Redis.CacheSetResult("PERSONAL:"+strconv.Itoa(id), Profile)
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
		log.Info("User tried to get personal, but was not autherized")
		w.Write([]byte("Unfortunately it seemed like you didn't have access, try login with blizzard again"))
	}

}

func GetPersonalRaiderio(w http.ResponseWriter, r *http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {

		charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
		char := Blizzard.CharacterMinimalFromMap(charMap)

		raiderio, e := Raider_io.GetRaiderIORank(Raider_io.CharInput{Name:char.Name, Realm:char.Realm, Region:FromLocaleToRegion(char.Locale)})

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(raiderio); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal, but was not autherized")
	}
}

func GetPersonalWarcraftLogs(w http.ResponseWriter, r *http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {

		charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
		char := Blizzard.CharacterMinimalFromMap(charMap)

		logs, e := WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name:char.Name, Realm:char.Realm, Region:FromLocaleToRegion(char.Locale)})

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
		log.Info("User tried to get personal, but was not autherized")
	}
}

func GetPersonalBlizzardChar(w http.ResponseWriter, r *http.Request){
	acces, id := DoesUserHaveAccess(w, r)
	if acces {

		charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
		char := Blizzard.CharacterMinimalFromMap(charMap)

		blizzChar, e := Blizzard.GetBlizzardChar(char)

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Error(e)
		} else {
			msg, err := json.Marshal(blizzChar); if err != nil{ log.Error(e); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}
	} else {
		log.Info("User tried to get personal, but was not autherized")
	}
}

