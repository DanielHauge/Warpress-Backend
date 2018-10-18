package Redis

import (
	log "../Logrus"
	"github.com/go-redis/redis"
)

func Set(key string, value string) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	client.Set(key, value, 0)
}

func Get(key string) string {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	v, e := client.Get(key).Result()
	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
	}
	return v
}
