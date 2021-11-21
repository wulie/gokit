package main

import (
	"fmt"
	"github.com/wulie/gokit/error_pkg/a"
	"time"
)

func main() {
	go func() {
		err := a.B()
		if err != nil {
			fmt.Printf("%+v", err)

		}
	}()

	fmt.Println("333333")
	for true {
		time.Sleep(time.Second)
		fmt.Println("333333")

	}
}
