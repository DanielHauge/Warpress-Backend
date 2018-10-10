package Redis

import (
	"github.com/go-redis/redis"
	"log"
)

func SetStruct(key string, obj map[string]interface{}){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})

	e := client.HMSet(key, obj).Err()
	if e != nil{
		log.Println(e.Error())
	}
}


func GetStruct(key string) (map[string]string, error){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: 0,
	})

	value, e := client.HGetAll(key).Result()
	if e != nil{
		log.Println(e.Error())
		return map[string]string{}, e
	}
	return value, e
}
