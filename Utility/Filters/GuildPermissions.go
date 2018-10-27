package Filters

import (
	"../../DataFormatters/Internal"
	"../../Integrations/BlizzardOpenAPI"
	"../../Postgres"
	"../../Redis"
	. "../../Utility/HttpHelper"
	log "../Logrus"
	"net/http"
	"strconv"
	"strings"
)

func RequireGuildMaster(HandleFunction func(w http.ResponseWriter, r *http.Request, id int, region string, guildstring string)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil{
				InterErrorHeader(w, e)
				return
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <- GuildRosterChannel

			if Result.Error == nil{

				guildwithmembers, _ := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isGM := false
				for _, member := range guildwithmembers.Members {
					if member.Rank == 0 && member.Character.Name == charactername {
						isGM = true
						break
					}
				}
				if isGM {
					HandleFunction(w, r, id, region, guildstring)
				} else {
					log.Info("User was not guildmaster of guild")
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("User was not guildmaster, which this operation requires"))
				}

			} else {

				log.Info("User is not in a detectable guild")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(Result.Error.Error()))

			}

		} else {
			log.Info("User did not have a main")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User did not have a main"))
		}

	}
}

func RequireOfficer(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil{
				InterErrorHeader(w, e)
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
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(e.Error() + " --- Guild might not be registered"))
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <- GuildRosterChannel
			if Result.Error == nil{
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
					log.Info("User was not guildmaster of guild")
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("User was not guildmaster, which this operation requires"))
				}
			} else {
				log.Info("User is not in a detectable guild")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(Result.Error.Error()))
			}

		} else {
			log.Info("User did not have a main")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User did not have a main"))
		}

	}
}

func RequireRaider(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil{
				InterErrorHeader(w, e)
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
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(e.Error() + " --- Guild might not be registered"))
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <- GuildRosterChannel
			if Result.Error == nil{
				guildwithmembers := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isOfficer := false
				for _, member := range guildwithmembers.Members {
					if member.Rank <= GuildStruct.Raider && member.Character.Name == charactername {
						isOfficer = true
						break
					}
				}
				if isOfficer {
					HandleFunction(w, r, GuildStruct.Id)
				} else {
					log.Info("User was not guildmaster of guild")
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("User was not guildmaster, which this operation requires"))
				}
			} else {
				log.Info("User is not in a detectable guild")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(Result.Error.Error()))
			}

		} else {
			log.Info("User did not have a main")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User did not have a main"))
		}

	}
}

func RequireTrial(HandleFunction func(w http.ResponseWriter, r *http.Request, guildid int)) func(w http.ResponseWriter, r *http.Request, id int, region string) {
	return func(w http.ResponseWriter, r *http.Request, id int, region string) {
		charactername, _, _, e := Postgres.GetMain(id)
		if e == nil {

			guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
			if e != nil{
				InterErrorHeader(w, e)
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
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(e.Error() + " --- Guild might not be registered"))
				return
			}

			GuildRosterChannel := Redis.ServeCacheAndUpdateBehind(guildstring, id, BlizzardOpenAPI.GuildWithMembers{}, Internal.FetchGuildRooster)
			Result := <- GuildRosterChannel
			if Result.Error == nil{
				guildwithmembers := Result.Obj.(*BlizzardOpenAPI.GuildWithMembers)
				isOfficer := false
				for _, member := range guildwithmembers.Members {
					if member.Rank <= GuildStruct.Trial && member.Character.Name == charactername {
						isOfficer = true
						break
					}
				}
				if isOfficer {
					HandleFunction(w, r, GuildStruct.Id)
				} else {
					log.Info("User was not guildmaster of guild")
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("User was not guildmaster, which this operation requires"))
				}
			} else {
				log.Info("User is not in a detectable guild")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(Result.Error.Error()))
			}

		} else {
			log.Info("User did not have a main")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User did not have a main"))
		}

	}
}
