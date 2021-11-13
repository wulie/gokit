package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(rdb)

	rdb.Set(context.Background(), "key", "wxh", time.Second*5)

	time.Sleep(time.Second * 3)
	val, err := rdb.Get(context.Background(), "key").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
	fmt.Println(val)
	ctx := context.Background()
	err = rdb.Publish(ctx, "mychannel1", "payload").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("1111111")
	pubsub := rdb.Subscribe(ctx, "mychannel1")
	fmt.Println("1111111")

	// Close the subscription when we are done.
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {

		}
	}(pubsub)
	fmt.Println("1111111")

	for {
		fmt.Println("1111111")

		msg, err := pubsub.ReceiveMessage(ctx)
		fmt.Println("1111111")

		if err != nil {
			panic(err)
		}
		fmt.Println("1111111")

		fmt.Println(msg.Channel, msg.Payload)
	}
}
