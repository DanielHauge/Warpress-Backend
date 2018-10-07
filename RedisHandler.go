package main

import (
	"github.com/go-redis/redis"
	"golang.org/x/oauth2"
	"log"
	"time"
)


func IsUserRegistered(accountid int) bool{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	d, _ := client.Exists(string(accountid)).Result()
	if d == 1{
		return true
	} else{
		return false
	}
}

func GetAccessToken(accountid int) (bool, string){
	AID := string(accountid)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	isRegistered, e := client.Exists(AID).Result()
	if isRegistered == 0 {return false, ""}
	accesToken, e := client.Get("AT:"+AID).Result()
	if e != nil { log.Println(e.Error())	}
	return true, accesToken
}

func CacheAccesToken(accountId int,accessToken *oauth2.Token){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	m := map[string]interface{}{
		"accessToken": accessToken.AccessToken,
		"Expire": accessToken.Expiry.String(),
		"Refresh": accessToken.RefreshToken,
		"Type": accessToken.TokenType,
	}
	AID := string(accountId)
	client.HMSet("AT:"+AID,m)
	client.Expire("AT:"+AID, accessToken.Expiry.Sub(time.Now()))
}

/*
func GetUserLoginInfo(accountid int) LoginInfo {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	mainChar, e := client.HGetAll("MAIN-"+string(accountid)).Result()
	alts, e := client.SMembers("ALTS-"+string(accountid)).Result()
	if e != nil { panic(e.Error()) }
	loginInfo := LoginInfo{Main:CharInfo{Name:mainChar["Name"], Realm:mainChar["Realm"], Locale:mainChar["Locale"]}, Alts:alts}
	return loginInfo
}
*/
