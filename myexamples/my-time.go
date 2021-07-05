package main

import (
	"fmt"
	"time"
)

func main() {
	// 秒数换成指定格式时间
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	// 指定格式时间转化为秒
	t, err := time.Parse("2006-01-02 15:04:05", "2021-06-30 14:23:53.607692")
	fmt.Printf("t=[%v], err=[%v]\n", t.Unix(), err)

	<-time.After(time.Second)
	fmt.Println("after")
	tn := time.NewTimer(200 * time.Millisecond)
	c := make(chan bool)
	go func() {
		for {
			select {
			case <-c:
				return
			case t := <-tn.C:
				fmt.Println(t)
			}
		}
	}()

	fmt.Println("dfa")
	tn.Stop()
	c <- true
	fmt.Println("over")

}
