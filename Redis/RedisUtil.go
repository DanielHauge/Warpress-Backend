package Redis

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)



var Addr string = os.Getenv("CONNECTION_STRING")
var Port string = ":6379"
var Password string = ""
var DB int = 0


// TODO: If availability ever becomes a problem, look into ClusterClient.
// TODO: If redis becomes cache only and availability becomes a problem, look into Ring for multiple redis servers.

func CanIConnect() error{
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})
	_, e := client.Ping().Result()

	return e
}

func DoesKeyExist(key string) bool{
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})
	d, e := client.Exists(key).Result()
	if e != nil{
		log.Println(e.Error())
	}

	if d == 1{
		return true
	} else {
		return false
	}
}

