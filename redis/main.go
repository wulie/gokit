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

	rdb.Set(context.Background(), "key", "wxh", time.Second)

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

	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(s []int) {
	s = append(s, 2048)
	s = append(s, 2048)
	s[0] = 1024
}
