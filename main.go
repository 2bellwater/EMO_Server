package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/cmd"
	"github.com/go-redis/redis/v7"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/redis"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:			"",
		DB:                 0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil{
		panic(err)
	}

	router := gin.Default()

	cmd.SetRestAPI(router)
	router.Run(":8080")

}