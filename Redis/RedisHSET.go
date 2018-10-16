package Redis

import (
	"errors"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func SetStruct(key string, obj map[string]interface{}){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})

	e := client.HMSet(key, obj).Err()
	if e != nil{
		log.Error(e, " -> Occured in Redis.SetStruct")
	}
}


func GetStruct(key string) (map[string]string, error){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: 0,
	})

	exists, e := client.Exists(key).Result()
	if e != nil{
		log.Error(e, " -> Occured in Redis.GetStruct")
		return nil, e
	}

	if exists == 1{
		value, e := client.HGetAll(key).Result()
		if e != nil{
			log.Error(e, " -> Occured in Redis.GetStruct")
			return map[string]string{}, e
		}
		return value, e
	} else {
		return map[string]string{}, errors.New("The key did not exist")
	}


}
