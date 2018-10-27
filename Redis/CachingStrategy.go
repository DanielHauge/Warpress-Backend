package Redis

import (
	"reflect"
	"strconv"
)

type ChannelResult struct {
	Obj interface{}
	Error error
}

func ServeCacheAndUpdateBehind(key string, id int, expectedType interface{}, fetcher func(id int, obj *interface{}) error) (chan ChannelResult) {
	channel := make(chan ChannelResult)
	rediskey := key + strconv.Itoa(id)

	go func() {
		result := reflect.New(reflect.TypeOf(expectedType)).Interface()
		var e error
		if DoesKeyExist(rediskey) {
			e = CacheGetResult(rediskey, &result)
			if !CacheTimeout(rediskey) { // If key is not in timeout, update cache.
				go func() {
					Caching := reflect.New(reflect.TypeOf(expectedType)).Interface()
					e = fetcher(id, &Caching)
					CacheSetResult(rediskey, Caching)
				}()
			}
		} else {
			e = fetcher(id, &result)
			if e == nil {
				go CacheSetResult(rediskey, result)
			}
		}
		channel <- ChannelResult{result, e}


	}()

	return channel

}
