package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type cc struct {
	country string
	city    string
}

var WgNum = 50

func main() {
	var opt uint64
	var wg sync.WaitGroup
	var entry atomic.Value
	entry.Store(cc{"china", "beijin"})
	var e cc = entry.Load().(cc)
	fmt.Println(e.country, e.city)

	for i := 0; i < WgNum; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&opt, 1)
			}
			wg.Done()
		}()
	}
	go func() {
		fmt.Println("start")
		// 多个Wait等待唤醒
		wg.Wait()
		fmt.Println("end")
	}()

	fmt.Println("wait")
	wg.Wait()
	fmt.Printf("opt=%d\n", atomic.LoadUint64(&opt))

}
