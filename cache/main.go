package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	c := cache.New(0, 0)
	now := time.Now()
	for i := 0; i < 1000000; i++ {
		c.Set(strconv.Itoa(i+1052), &Realtime{
			ID: int64(1052 + i),
			GN: "W3.D." + strconv.Itoa(i),
			AV: rand.Float32(),
			TM: 0,
			DS: 0,
		}, 0)
	}

	fmt.Println(time.Since(now))
	time.Sleep(time.Second * 10)
	for i := 0; i < 1000000; i++ {
		get, b := c.Get(strconv.Itoa(i + 1052))
		fmt.Println(get, b, "222222222222")
	}

	time.Sleep(time.Minute)
	for i := 0; i < 1000000; i++ {
		get, b := c.Get(strconv.Itoa(i + 1052))
		fmt.Println(get, b, "11111111111111111111111")
	}
}

type Realtime struct {
	ID int64
	GN string
	AV interface{}
	TM int64
	DS int16
}
