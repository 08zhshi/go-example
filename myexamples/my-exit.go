package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("qwe")
	}()
	fmt.Println("123")

	os.Exit(2) // 退出不会执行defer函数
}
