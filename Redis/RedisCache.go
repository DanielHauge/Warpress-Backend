package Redis

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"log"
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
		log.Println(err.Error())
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
		log.Println(err.Error())
	}
	return err
}
