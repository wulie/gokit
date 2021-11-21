package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func redisConnect() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		//Addr:     redisServer + ":" + port,
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	return
}

func pubMessage(channel, msg string) {
	rdb := redisConnect()
	rdb.Publish(context.Background(), channel, msg)
}

func main() {
	channel := "hello"
	msgList := []string{"hello", "world"}

	// 此处发了两个消息
	i := 0
	for true {
		i++
		for _, msg := range msgList {
			pubMessage(channel, msg+strconv.Itoa(i))
			fmt.Printf("已经发送%d%s到%s\n", i, msg, channel)
		}
		time.Sleep(time.Second)
	}
}

/*密码配置
################################## SECURITY ###################################

# Require clients to issue AUTH <PASSWORD> before processing any other
# commands.  This might be useful in environments in which you do not trust
# others with access to the host running redis-server.
#
# This should stay commented out for backward compatibility and because most
# people do not need auth (e.g. they run their own servers).
#
# Warning: since Redis is pretty fast an outside user can try up to
# 150k passwords per second against a good box. This means that you should
# use a very strong password otherwise it will be very easy to break.
#
# requirepass foobared
requirepass 123456

启动方式 redis-server ../redis.conf
*/
