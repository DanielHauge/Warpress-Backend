package main

import (
	"./DataFormatters/Guild"
	"./DataFormatters/Personal"
	"./Integrations/BlizzardOpenAPI"
	"./Integrations/Raider.io"
	"./Integrations/WarcraftLogs"
	Postgres "./Postgres/PreparedProcedures"
	"./Redis"
	"./Utility/HttpHelper"
	. "./Utility/HttpHelper"
	log "./Utility/Logrus"
	"bytes"
	"github.com/gorilla/mux"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func SetupIndexPage() []byte {
	var buffer bytes.Buffer

	for _, v := range routes {
		buffer.WriteString("\n\n### " + v.Name + "\n")
		buffer.WriteString("##### Route: " + v.Pattern + "\n")
		buffer.WriteString("##### Method: " + v.Method + "\n")

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
	output := blackfriday.Run([]byte(buffer.Bytes()), blackfriday.WithExtensions(blackfriday.FencedCode))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(output)
}

func HandleGetPersonalFull(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("PERSONAL:", id, Personal.Overview{}, Personal.FetchFullPersonal)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}

}

func HandleGetInspectFull(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	charString := vars["region"]+":"+vars["realm"]+":"+vars["name"]

	unwrapFunc := func(id int, obj *interface{}) error {
		e := Personal.FetchFullInspect(vars["name"], vars["realm"], vars["region"], obj)
		return e
	}

	channel := Redis.ServeCacheAndUpdateBehind("INSPECT:"+charString, 0, Personal.Overview{}, unwrapFunc)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}



}

func HandleGetPersonalRaiderio(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("PERSONAL/RAIDERIO:", id, Raider_io.CharacterProfile{}, Personal.FetchRaiderioPersonal)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}

}

func HandleGetPersonalWarcraftLogs(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("PERSONAL/LOGS:", id, WarcraftLogs.Encounter{}, Personal.FetchWarcraftlogsPersonal)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}

}

func HandleGetPersonalBlizzardChar(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("PERSONAL/BLIZZARD:", id, BlizzardOpenAPI.FullCharInfo{}, Personal.FetchBlizzardPersonal)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}
}

func HandleGetPersonalImprovements(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel := Redis.ServeCacheAndUpdateBehind("PERSONAL/IMPROVEMENT:", id, Personal.Improvements{}, Personal.FetchPersonalImprovementsFull)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}

}

func HandleGetGuildOverview(w http.ResponseWriter, r *http.Request, guildid int) {

	channel := Redis.ServeCacheAndUpdateBehind("GUILD/OVERVIEW:", guildid, Guild.Overview{}, Guild.FetchFullGuildOverview)
	result := <-channel
	if result.Error == nil {

		msg, err := json.Marshal(result.Obj)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	} else {

		log.WithLocation().WithError(result.Error).Error("How!")
		InterErrorHeader(w, result.Error)

	}

}

func HandleLogout(w http.ResponseWriter, r *http.Request, id int, region string) {

	e := Redis.DeleteKey("AT:" + strconv.Itoa(id))
	if e != nil {
		InterErrorHeader(w, e)
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK!"))
	}

}

func HandleGuildRegistration(w http.ResponseWriter, r *http.Request, id int, region string, guildstring string) {
	/*
		Checks:
		- Does guild allready exists?
		- Permissions wrapper.
	*/

	split := strings.Split(guildstring, ":")
	GuildFromRedis := struct {
		Name   string
		Realm  string
		Region string
	}{Name: split[0], Realm: split[1], Region: split[2]}

	var Guild struct {
		Officerrank int `json:"officerrank"`
		Raiderrank  int `json:"raiderrank"`
		Trialrank   int `json:"trialrank"`
	}
	HttpHelper.ReadFromRequest(w, r, &Guild)

	if e := Postgres.RegisterGuild(GuildFromRedis.Name, GuildFromRedis.Realm, GuildFromRedis.Region, Guild.Officerrank, Guild.Raiderrank, Guild.Trialrank); e != nil {
		InterErrorHeader(w, e)
		return
	}

	SuccessHeader(w, []byte("Succes"))

}


