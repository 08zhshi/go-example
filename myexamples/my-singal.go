package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	go func() {
		s := <-sig
		fmt.Println(s)
		done <- true
	}()

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("waiting")
	<-done
	fmt.Println("finished")
}
