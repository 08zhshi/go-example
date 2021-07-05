package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()        // 从1970到现在秒
	sec1 := now.Second()     // 当前时间秒（60s）
	nao := now.UnixNano()    // 从1970到现在纳秒
	nao1 := now.Nanosecond() // 当前时间纳秒
	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(sec1)
	fmt.Println(nao)
	fmt.Println(nao1)

	mills := nao / 1000000 // 换算毫秒数
	fmt.Println(mills)

	// 秒数换成指定格式时间
	fmt.Println(time.Unix(sec, 0).Format("2006-01-02 15:04:05"))
	// 指定格式时间转化为秒
	t, err := time.Parse("2006-01-02 15:04:05", "2021-06-30 14:23:53.607692")
	fmt.Printf("t=[%v], err=[%v]", t.Unix(), err)
}
