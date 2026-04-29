package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9" //Official Go client library for Redis.Which let's Go to talk to Redis server
)

var RedisClient *redis.Client //var for storing redis object for connction
var Ctx = context.Background()

func ConnectRedis() { //connects redis to go app
	RedisClient = redis.NewClient(&redis.Options{ //Create connection config.
		Addr: "localhost:6379", //redis server location
	})

	_, err := RedisClient.Ping(Ctx).Result() //Testing redis by sending a message
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis connected")
}
