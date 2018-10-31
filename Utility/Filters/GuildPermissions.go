package Filters

import (
	"../../DataFormatters/Internal"
	"../../Integrations/BlizzardOpenAPI"
	Postgres "../../Postgres/PreparedProcedures"
	"../../Redis"
	. "../../Utility/HttpHelper"
	log "../Logrus"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func RequireGuildMaster(HandleFunction func(w http.ResponseWriter, r *http.Request, id int, region string, guildstring string)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil {
				InterErrorHeader(w, e, "The requesting users guild cannot be detected, this can be solved by loading a page that force the players guild to be fetched", GetStatusCodeByError(e))
				return
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <-GuildRosterChannel

			if Result.Error == nil {

				guildwithmembers, _ := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isGM := false
				for _, member := range guildwithmembers.Members {
					if /*member.Rank == 0 &&  */ member.Character.Name == charactername { // TODO: make guild master only.
						isGM = true
						break
					}
				}
				if isGM {
					HandleFunction(w, r, id, region, guildstring)
				} else {
					log.Info("User was not guildmaster of guild")
					e = errors.New("Insufficient rank in guild")
					InterErrorHeader(w, e, "User was not Guild master", GetStatusCodeByError(e))
				}
			} else {
				log.Info("User is not in a detectable guild")
				InterErrorHeader(w, e, "Unable to get guild roster", GetStatusCodeByError(e))
			}

		} else {
			log.Info("User did not have a main -> therefor guild to fetch is unknown")
			InterErrorHeader(w, e, "User did not have a main -> therefor guild to fetch is unknown", GetStatusCodeByError(e))
		}

	}
}

func RequireOfficer(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil {
				InterErrorHeader(w, e, "The requesting users guild cannot be detected, this can be solved by loading a page that force the players guild to be fetched", GetStatusCodeByError(e))
				return
			}
			split := strings.Split(guildstring, ":")
			guild := struct {
				Name   string
				Realm  string
				Region string
			}{Name: split[0], Realm: split[1], Region: split[2]}
			GuildStruct, e := Postgres.GetGuildByComposite(guild.Name, guild.Realm, guild.Region)
			if e != nil {
				log.WithLocation().WithError(e).Error("Could not get guild, might not be registered")
				InterErrorHeader(w, e, "Was unable not able to fetch guild", GetStatusCodeByError(e))
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <-GuildRosterChannel
			if Result.Error == nil {
				guildwithmembers := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isOfficer := false
				for _, member := range guildwithmembers.Members {
					if member.Rank <= GuildStruct.Officer && member.Character.Name == charactername {
						isOfficer = true
						break
					}
				}
				if isOfficer {
					HandleFunction(w, r, GuildStruct.Id)
				} else {
					log.Info("User was not atleast officer of guild")
					e = errors.New("Insufficient rank in guild")
					InterErrorHeader(w, e, "User was not Officer", GetStatusCodeByError(e))
				}
			} else {
				log.Info("User is not in a detectable guild")
				InterErrorHeader(w, e, "Unable to get guild roster", GetStatusCodeByError(e))
			}

		} else {
			log.Info("User did not have a main -> therefor guild to fetch is unknown")
			InterErrorHeader(w, e, "User did not have a main -> therefor guild to fetch is unknown", GetStatusCodeByError(e))
		}

	}
}

func RequireRaider(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil {
				InterErrorHeader(w, e, "The requesting users guild cannot be detected, this can be solved by loading a page that force the players guild to be fetched", GetStatusCodeByError(e))
				return
			}
			split := strings.Split(guildstring, ":")
			guild := struct {
				Name   string
				Realm  string
				Region string
			}{Name: split[0], Realm: split[1], Region: split[2]}
			GuildStruct, e := Postgres.GetGuildByComposite(guild.Name, guild.Realm, guild.Region)
			if e != nil {
				log.WithLocation().WithError(e).Error("Could not get guild, might not be registered")
				InterErrorHeader(w, e, "Was unable not able to fetch guild", GetStatusCodeByError(e))
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <-GuildRosterChannel
			if Result.Error == nil {
				guildwithmembers := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isRaider := false
				for _, member := range guildwithmembers.Members {
					if member.Rank <= GuildStruct.Raider && member.Character.Name == charactername {
						isRaider = true
						break
					}
				}
				if isRaider {
					HandleFunction(w, r, GuildStruct.Id)
				} else {
					log.Info("User was not atleast raider of guild")
					e = errors.New("Insufficient rank in guild")
					InterErrorHeader(w, e, "User was not Raider", GetStatusCodeByError(e))
				}
			} else {
				log.Info("User is not in a detectable guild")
				InterErrorHeader(w, e, "Unable to get guild roster", GetStatusCodeByError(e))
			}

		} else {
			log.Info("User did not have a main -> therefor guild to fetch is unknown")
			InterErrorHeader(w, e, "User did not have a main -> therefor guild to fetch is unknown", GetStatusCodeByError(e))
		}

	}
}

func RequireTrial(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil {
				InterErrorHeader(w, e, "The requesting users guild cannot be detected, this can be solved by loading a page that force the players guild to be fetched", GetStatusCodeByError(e))
				return
			}
			split := strings.Split(guildstring, ":")
			guild := struct {
				Name   string
				Realm  string
				Region string
			}{Name: split[0], Realm: split[1], Region: split[2]}
			GuildStruct, e := Postgres.GetGuildByComposite(guild.Name, guild.Realm, guild.Region)
			if e != nil {
				log.WithLocation().WithError(e).Error("Could not get guild, might not be registered")
				InterErrorHeader(w, e, "Was unable not able to fetch guild", GetStatusCodeByError(e))
				return
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <-GuildRosterChannel
			if Result.Error == nil {
				var guildwithmembers *BlizzardOpenAPI.GuildWithMembers
				switch t := Result.Obj.(type) {
				case BlizzardOpenAPI.GuildWithMembers:
					guildwithmembers = &t
				case *BlizzardOpenAPI.GuildWithMembers:
					guildwithmembers = t
				}
				isTrial := false
				for _, member := range guildwithmembers.Members {
					if member.Rank <= GuildStruct.Trial && member.Character.Name == charactername {
						isTrial = true
						break
					}
				}
				if isTrial {
					HandleFunction(w, r, GuildStruct.Id)
				} else {
					log.Info("User was not atleast trial of guild")
					e = errors.New("Insufficient rank in guild")
					InterErrorHeader(w, e, "User was not Raider", GetStatusCodeByError(e))
				}
			} else {
				log.Info("User is not in a detectable guild")
				InterErrorHeader(w, e, "Unable to get guild roster", GetStatusCodeByError(e))
			}

		} else {
			log.Info("User did not have a main -> therefor guild to fetch is unknown")
			InterErrorHeader(w, e, "User did not have a main -> therefor guild to fetch is unknown", GetStatusCodeByError(e))
		}

	}
}
