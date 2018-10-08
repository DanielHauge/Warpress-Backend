package main

import (
	"errors"
	"github.com/go-redis/redis"
	"golang.org/x/oauth2"
	"log"
	"strconv"
	"time"
)


func IsUserRegistered(accountid int) bool{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	d, _ := client.Exists("MAIN:"+strconv.Itoa(accountid)).Result()
	if d == 1{
		return true
	} else{
		return false
	}
}

func GetAccessToken(accountid int) (oauth2.Token, error){
	AID := strconv.Itoa(accountid)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	isRegistered, e := client.Exists("AT:"+AID).Result()
	if isRegistered == 0 {
		log.Println("User does not have any accessToken stored in the system.")
		return oauth2.Token{}, errors.New("User does not have any accessToken stored in system")
	}
	value, e := client.HGetAll("AT:"+AID).Result()
	time, e := time.Parse(time.RFC3339,value["expire"])
	accessToken := oauth2.Token{
		Expiry: time,
		TokenType: value["tokentype"],
		RefreshToken: value["refreshtoken"],
		AccessToken: value["accesstoken"],
	}
	if e != nil {
		log.Println(e.Error())
		return oauth2.Token{}, e
	}
	return accessToken, nil
}

func CacheAccesToken(accountId int,accessToken *oauth2.Token){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	m := map[string]interface{}{
		"accesstoken": accessToken.AccessToken,
		"expire": accessToken.Expiry.Format(time.RFC3339),
		"refresh": accessToken.RefreshToken,
		"tokentype": accessToken.TokenType,
	}
	expiredur := accessToken.Expiry.Sub(time.Now())
	AID := strconv.Itoa(accountId)
	client.HMSet("AT:"+AID,m)

	client.Expire("AT:"+AID, expiredur)
}

func GetMainChar(accountid int) (charRequest, error){
	AID := strconv.Itoa(accountid)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	value, e := client.HGetAll("MAIN:"+AID).Result()
	if e != nil{
		log.Println(e.Error())
		return charRequest{}, e
	}
	d := charRequest{Name:value["name"], Realm:value["realm"], Locale:value["locale"]}
	return d, nil
}

func SetMainChar(accountid int, mainChar charRequest){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	m := map[string]interface{}{
		"name": mainChar.Name,
		"realm": mainChar.Realm,
		"locale": mainChar.Locale,
	}
	AID := strconv.Itoa(accountid)
	client.HMSet("MAIN:"+AID, m)

}
