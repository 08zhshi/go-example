package main

import (
	"fmt"
	"time"
)

func work(w int, jobs chan int, res chan int) {
	for j := range jobs {
		time.Sleep(time.Second)
		res <- 2 * j
	}
}
func main() {
	jobs := make(chan int, 5)
	res := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go work(i, jobs, res)
	}

	for j := 0; j < 3; j++ {
		jobs <- j
	}
	close(jobs)
	for j := 0; j < 3; j++ {
		fmt.Println(<-res)
	}
}
