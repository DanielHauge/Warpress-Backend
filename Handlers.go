package main

import (
	"./DataFormatters/Guild"
	"./DataFormatters/Personal"
	"./Redis"
	. "./Utility/HttpHelper"
	log "./Utility/Logrus"
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
	"reflect"
	"strconv"
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

	channel, error := Redis.ServeCacheAndUpdateBehind("PERSONAL:", id, Personal.FetchFullPersonal)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

	}

}

func HandleGetPersonalRaiderio(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel, error := Redis.ServeCacheAndUpdateBehind("PERSONAL/RAIDERIO:", id, Personal.FetchRaiderioPersonal)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

	}

}

func HandleGetPersonalWarcraftLogs(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel, error := Redis.ServeCacheAndUpdateBehind("PERSONAL/LOGS:", id, Personal.FetchWarcraftlogsPersonal)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

	}

}

func HandleGetPersonalBlizzardChar(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel, error := Redis.ServeCacheAndUpdateBehind("PERSONAL/BLIZZARD:", id, Personal.FetchBlizzardPersonal)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

	}
}

func HandleGetPersonalImprovements(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel, error := Redis.ServeCacheAndUpdateBehind("PERSONAL/IMPROVEMENT:", id, Personal.FetchPersonalImprovementsFull)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

	}

}

func HandleGetGuildOverview(w http.ResponseWriter, r *http.Request, id int, region string) {

	channel, error := Redis.ServeCacheAndUpdateBehind("GUILD/OVERVIEW:", id, Guild.FetchFullGuildOverview)

	select {

	case result := <-channel:
		msg, err := json.Marshal(result)
		if err != nil {
			log.WithLocation().WithError(err).Error("Was not able to marshal raider.io profile")
			InterErrorHeader(w, err)
		} else {
			SuccessHeader(w, msg)
		}

	case e := <-error:
		log.WithLocation().WithError(e).Error("How!")
		InterErrorHeader(w, e)

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
