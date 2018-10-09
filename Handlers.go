package main

import (
	"bytes"
	"encoding/json"
	"github.com/avelino/slugify"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"net/http"
	"reflect"
	"./WarcraftLogs"
	"./Raider.io"
	"./Wowprogress"
	"./Blizzard"
	"sync"
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


func GetPersonalCharInfo(w http.ResponseWriter, r *http.Request) {

	acces, id := DoesUserHaveAccess(w, r)
	if acces {

		var Profile PersonalProfile

		char, e := GetMainChar(id)

		var wg sync.WaitGroup
		var blizzwait sync.WaitGroup
		blizzwait.Add(1)
		wg.Add(4)

		var blizzChar Blizzard.FullCharInfo
		go func(){
			blizzChar, e = GetBlizzardChar(char)
			Profile.Character = blizzChar
			wg.Done()
			blizzwait.Done()
		}()

		var raiderio Raider_io.CharacterProfile
		go func() {
			raiderio, e = Raider_io.GetRaiderIORank(Raider_io.CharInput{Name:char.Name, Realm:char.Realm, Region:FromLocaleToRegion(char.Locale)})
			Profile.RaiderIOProfile = raiderio
			wg.Done()
		}()


		var logs []WarcraftLogs.Encounter
		go func() {
			logs, e = WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name:char.Name, Realm:char.Realm, Region:FromLocaleToRegion(char.Locale)})
			Profile.WarcraftLogsRanks = logs
			wg.Done()
		}()

		var wowprog Wowprogress.GuildRank
		go func() {
			blizzwait.Wait()
			wowprog, e = Wowprogress.GetGuildRank(Wowprogress.Input{Region: FromLocaleToRegion(char.Locale), Realm: slugify.Slugify(char.Realm), Guild: blizzChar.Guild.Name})
			Profile.GuildRank = wowprog
			wg.Done()
		}()

		wg.Wait()

		if e != nil{
			w.WriteHeader(500)
			w.Write([]byte(e.Error()))
			log.Println(e.Error())
		} else {
			msg, err := json.Marshal(Profile); if err != nil{ log.Println(err); w.Write([]byte(err.Error())); return}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(200)
			w.Write(msg)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}




