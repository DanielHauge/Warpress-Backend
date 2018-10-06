package main

import (
	"github.com/go-redis/redis"
)


func IsUserRegistered(accountid int) bool{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	d, _ := client.Exists("AC:"+string(accountid)).Result()
	if d == 1{
		return true
	} else{
		return false
	}
}

func GetUserLoginInfo(accountid int) LoginInfo {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	mainChar, e := client.Get("MAIN-"+string(accountid)).Result()
	alts, e := client.SMembers("ALTS-"+string(accountid)).Result()
	if e != nil { panic(e.Error()) }
	loginInfo := LoginInfo{Main:mainChar, Alts:alts}
	return loginInfo
}
