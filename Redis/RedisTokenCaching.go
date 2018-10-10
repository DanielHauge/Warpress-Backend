package Redis

import (
	"errors"
	"github.com/go-redis/redis"
	"golang.org/x/oauth2"
	"log"
	"time"
)

func GetAccessToken(key string) (oauth2.Token, error){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})
	isRegistered, e := client.Exists(key).Result()
	if isRegistered == 0 {
		log.Println("User does not have any accessToken stored in the system.")
		return oauth2.Token{}, errors.New("User does not have any accessToken stored in system")
	}
	value, e := client.HGetAll(key).Result()
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

func CacheAccesToken(key string,accessToken *oauth2.Token){
	client := redis.NewClient(&redis.Options{
		Addr: Addr+Port,
		Password: Password,
		DB: DB,
	})

	m := map[string]interface{}{
		"accesstoken": accessToken.AccessToken,
		"expire": accessToken.Expiry.Format(time.RFC3339),
		"refresh": accessToken.RefreshToken,
		"tokentype": accessToken.TokenType,
	}
	expireDuration := accessToken.Expiry.Sub(time.Now())
	client.HMSet(key,m)
	client.Expire(key, expireDuration)
}

