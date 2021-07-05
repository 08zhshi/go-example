package main

import "fmt"

const s = "name"

func main() {
	const i = 40
	fmt.Println("s=", s)
	fmt.Println("i=", i)
	const n = 500000000
	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
