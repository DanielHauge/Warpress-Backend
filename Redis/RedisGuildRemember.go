package Redis

import (
	"../Integrations/BlizzardOpenAPI"
	"../Postgres"
	log "../Utility/Logrus"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

func Set(key string, value string) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	client.Set(key, value, 0)
}

func Get(key string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	v, e := client.Get(key).Result()
	if e != nil {
		log.WithLocation().WithError(e).Warn("Hov!")
		v, e = refreshGuild(key)
	}
	return v, e
}

func refreshGuild(key string) (string, error) {
	split := strings.Split(key, ":")
	id, e := strconv.Atoi(split[1])
	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return "", e
	}
	blizzChar, e := BlizzardOpenAPI.GetBlizzardChar(realm, name, region)
	go Set("GUILD:"+strconv.Itoa(id), blizzChar.Guild.Name+":"+blizzChar.Guild.Realm+":"+region)
	return blizzChar.Guild.Name + ":" + blizzChar.Guild.Realm + ":" + region, e
}
