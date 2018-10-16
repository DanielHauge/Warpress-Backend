package Redis

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"time"
)

func CacheSetResult(key string, obj interface{}){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})
	codec := &cache.Codec{
		Redis: client,

		Marshal: func (v interface{}) ([]byte, error){
			return msgpack.Marshal(v)
		},
		Unmarshal: func (b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	err := codec.Set(&cache.Item{
		Key: key,
		Object: obj,
		Expiration: time.Minute*10,
	})
	if err != nil {
		log.Error(err, " -> Occured in redis.CacheSet!")
	}
}


func CacheGetResult(Key string, obj interface{}) error{
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})

	codec := &cache.Codec{
		Redis: client,

		Marshal: func (v interface{}) ([]byte, error){
			return msgpack.Marshal(v)
		},
		Unmarshal: func (b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	err := codec.Get(Key, obj)
	if err != nil {
		log.Error(err, " -> Occured in redis.CacheGet!")
	}
	return err
}
