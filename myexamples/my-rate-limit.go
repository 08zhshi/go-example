package main

import (
	"fmt"
	"time"
)

func main() {

	req := make(chan int, 5)
	limit := make(chan time.Time, 3)
	go func() {
		for c := range time.Tick(200 * time.Millisecond) {
			limit <- c
		}
	}()

	for i := 0; i < 5; i++ {
		req <- i
	}
	close(req)

	for r := range req {
		d := <-limit
		fmt.Println(r, d.UnixNano())
	}
}
