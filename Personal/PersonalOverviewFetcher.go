package Personal

import (
	"../Integrations/BlizzardOauthAPI"
	"../Integrations/BlizzardOpenAPI"
	"../Integrations/Raider.io"
	"../Integrations/WarcraftLogs"
	"../Integrations/Wowprogress"
	"../Redis"
	"github.com/avelino/slugify"
	"github.com/jinzhu/copier"
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
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)

	var wg sync.WaitGroup
	var blizzwait sync.WaitGroup
	blizzwait.Add(1)
	wg.Add(4)

	var blizzChar BlizzardOpenAPI.FullCharInfo
	go func(){
		blizzChar, e = BlizzardOpenAPI.GetBlizzardChar(char.Realm, char.Name, char.Region)
		Profile.Character = blizzChar
		wg.Done()
		blizzwait.Done()
	}()

	var raiderio Raider_io.CharacterProfile
	go func() {
		raiderio, e = Raider_io.GetRaiderIORank(Raider_io.CharInput{Name:char.Name, Realm:char.Realm, Region:char.Region})
		Profile.RaiderIOProfile = raiderio
		wg.Done()
	}()


	var logs []WarcraftLogs.Encounter
	go func() {
		logs, e = WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name:char.Name, Realm:char.Realm, Region:char.Region})
		Profile.WarcraftLogsRanks = logs
		wg.Done()
	}()

	var wowprog Wowprogress.GuildRank
	go func() {
		blizzwait.Wait()
		wowprog, e = Wowprogress.GetGuildRank(Wowprogress.Input{Region: char.Region, Realm: slugify.Slugify(char.Realm), Guild: blizzChar.Guild.Name})
		Profile.GuildRank = wowprog
		wg.Done()
	}()

	wg.Wait()
	return e
}

func FetchRaiderioPersonal(id int, Profile *Raider_io.CharacterProfile) error{

	charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
	if e != nil{
		log.Error(e, " -> It seems there is no main registered to the requesting user")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)
	prof, e := Raider_io.GetRaiderIORank(Raider_io.CharInput{Name:char.Name, Realm:char.Realm, Region:char.Region})
	copier.Copy(Profile, &prof)
	return e
}

func FetchWarcraftlogsPersonal(id int, Logs *[]WarcraftLogs.Encounter) error{
	charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
	if e != nil{
		log.Error(e, " -> It seems there is no main registered to the requesting user")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)

	logs, e := WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name:char.Name, Realm:char.Realm, Region:char.Region})
	copier.Copy(Logs, &logs)
	return e
}

func FetchBlizzardPersonal(id int, Profile *BlizzardOpenAPI.FullCharInfo) error{

	charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
	if e != nil{
		log.Error(e, " -> It seems there is no main registered to the requesting user")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)
	blizzChar, e := BlizzardOpenAPI.GetBlizzardChar(char.Realm, char.Name, char.Region)
	copier.Copy(Profile, &blizzChar)
	return e
}
