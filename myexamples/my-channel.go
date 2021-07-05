package main

import (
	"fmt"
	"time"
)

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func job(c chan int, d chan bool) {
	for {
		j, more := <-c
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			d <- true
			break
		}
	}
}

func main() {
	message := make(chan string, 2)

	message <- "hello"
	message <- "world"

	fmt.Println(<-message)
	fmt.Println(<-message)

	recv := make(chan string, 1)
	send := make(chan string, 1)
	recv <- "you"
	pong(recv, send)
	fmt.Println(<-send)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	c := make(chan int, 5)
	d := make(chan bool)
	go job(c, d)
	for j := 1; j <= 3; j++ {
		c <- j
		fmt.Println("sent job", j)
	}
	close(c)
	fmt.Println("sent all jobs")
	<-d

	e := make(<-chan bool)  // 从通道接收
	e1 := make(chan<- bool) // 发送到通道
	<-e
	e1 <- true
}
