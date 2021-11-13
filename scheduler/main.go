package main

import (
	"fmt"
	"github.com/hslam/scheduler"
	"sync"
)

func main() {
	s := scheduler.New(64, nil)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		task := func() {
			fmt.Println(i)
			wg.Done()
		}
		s.Schedule(task)
	}
	wg.Wait()
	s.Close()
}
