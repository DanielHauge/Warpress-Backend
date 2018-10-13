package main

import (
	"./Blizzard"
	"./Raider.io"
	"./Redis"
	"./WarcraftLogs"
	"./Wowprogress"
	"github.com/avelino/slugify"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

func FetchFullPersonal(id int, Profile *PersonalProfile) error{

	charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
	if e != nil{
		log.Error(e, " -> It seems there is no main registered to the requesting user")
		return e
	}
	char := Blizzard.CharacterMinimalFromMap(charMap)

	var wg sync.WaitGroup
	var blizzwait sync.WaitGroup
	blizzwait.Add(1)
	wg.Add(4)

	var blizzChar Blizzard.FullCharInfo
	go func(){
		blizzChar, e = Blizzard.GetBlizzardChar(char)
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
	return e
}
