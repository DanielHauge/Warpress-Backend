package Redis

import (
	"reflect"
	"strconv"
)

func ServeCacheAndUpdateBehind(key string, id int, expectedType interface{}, fetcher func(id int, obj *interface{}) error) (chan interface{}, chan error) {
	channel := make(chan interface{})
	errorcheck := make(chan error)
	rediskey := key + strconv.Itoa(id)

	go func() {
		result := reflect.New(reflect.TypeOf(expectedType)).Interface()
		var e error
		if DoesKeyExist(rediskey) {
			e = CacheGetResult(rediskey, &result)
			if !CacheTimeout(rediskey) { // If key is not in timeout, update cache.
				go func() {
					var Caching interface{}
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
		if e != nil {
			errorcheck <- e
		} else {
			channel <- result
		}

	}()

	return channel, errorcheck

}
