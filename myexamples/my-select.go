package main

import "fmt"

func main() {
	sig := make(chan bool)
	sig1 := make(chan bool)
	select {
	case b := <-sig:
		fmt.Println(b)
	case b := <-sig1:
		fmt.Println(b)
	default:
		fmt.Println("default")
	}

	//go func() {
	//	<-sig
	//}()

	// 死锁了，所有线程都sleep。
	select {
	case b := <-sig:
		fmt.Println(b)
	case sig <- true:
		fmt.Println("ok")
		//default:
		//	fmt.Println("no")
	}
}
