package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {
	myPool := sync.Pool{
		New: func() interface{} {
			i := 1
			return i
		},
	}

	if v := myPool.Get(); v != 1 {
		fmt.Println("failed")
		os.Exit(1)
	}
	myPool.Put(2)
	runtime.GC()
	if v, ok := myPool.Get().(int); ok {
		fmt.Println(v)
	}
	if v, ok := myPool.Get().(int); ok {
		fmt.Println(v)
	}

	//runtime.GC()

}
