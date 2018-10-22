package Redis

import (
	log "../Utility/Logrus"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"time"
)

func CacheSetResult(key string, obj interface{}) {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	codecs := &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	err := codecs.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Minute * 10,
	})
	if err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
	}
}

func CacheGetResult(Key string, obj interface{}) error {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})

	codecs := &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	err := codecs.Get(Key, obj)
	if err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
	}
	return err
}

func CacheTimeout(key string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     Addr + Port,
		Password: Password,
		DB:       DB,
	})
	remaingDur, e := client.TTL(key).Result()
	if e != nil{
		log.WithLocation().WithError(e).Error("Hov!")
	}
	if remaingDur > time.Minute*10-time.Second*30{
		return true
	}
	return false

}
