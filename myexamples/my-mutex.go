package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var mtx sync.Mutex
	map1 := make(map[int]int)
	var readOps uint64
	var writeOps uint64

	for i := 0; i < 10; i++ {
		go func() {
			key := rand.Intn(5)
			mtx.Lock()
			_ = map1[key]
			mtx.Unlock()
			atomic.AddUint64(&readOps, 1)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mtx.Lock()
			map1[key] = val
			mtx.Unlock()
			atomic.AddUint64(&writeOps, 1)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadUint64(&readOps))
	fmt.Println(atomic.LoadUint64(&writeOps))
	fmt.Println(map1)

}
