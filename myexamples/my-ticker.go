package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.NewTicker(200 * time.Millisecond)
	c := make(chan bool)
	go func() {
		for {
			select {
			case t := <-tick.C:
				fmt.Println(t)
			case <-c:
				return
			}
		}
	}()

	time.Sleep(time.Second)
	tick.Stop()
	c <- true
	fmt.Println("over")
}
