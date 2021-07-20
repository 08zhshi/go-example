package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex
var val int

func read() {
	lock.RLock()
	fmt.Printf("read val:%d\n", val)
	lock.RUnlock()
}

func main() {
	fmt.Printf("start\n")

	go func() {
		for i := 0; i < 10; i++ {
			val = i
			fmt.Printf("write val:%d\n", i)
			read()
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Printf("end\n")
}
