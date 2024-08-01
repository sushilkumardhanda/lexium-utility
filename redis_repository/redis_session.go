package redis_repository

import (
	"lexium-utility/config"
	"time"
)

var client = config.GetRedisClient()

func GetSessionId(username string) (string, error) {
	val, err := client.Get(username).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
func SetSessionId(username string, sid string) error {
	_, err := client.Set(username, sid, 24*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}
func Delete(username string) error {
	_, err := client.Del(username).Result()
	if err != nil {
		return err
	}
	return nil
}
