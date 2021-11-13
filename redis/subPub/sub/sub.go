package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func redisConnect() (rdb *redis.Client) {

	var (

	//password    string= "123456"
	)

	//password = os.Getenv("RedisPass")

	rdb = redis.NewClient(&redis.Options{
		//Addr:     redisServer + ":" + port,
		Addr: "localhost:6379",

		Password: "123456",
		DB:       0, // use default DB
	})

	return
}

func subMessage(channel string) {
	rdb := redisConnect()
	pubsub := rdb.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func main() {
	channel := "hello"
	subMessage(channel)
}

//https://zhuanlan.zhihu.com/p/330602361
