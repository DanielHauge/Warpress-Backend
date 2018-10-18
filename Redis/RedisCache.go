package Redis

import (
	log "../Logrus"
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
