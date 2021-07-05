package main

import (
	"fmt"
	"reflect"
)

// 接收可变个参数...
func myfunc(v ...interface{}) {
	fmt.Println(reflect.ValueOf(v).Kind())
	// v与v...不一样，后者将切片变为多个元素
	myfunc1(v)
	myfunc1(v...)
}
func myfunc1(v ...interface{}) {
	for _, c := range v {
		switch c.(type) {
		case string:
			fmt.Println("string=", c)
		case int64:
			fmt.Println("int64=", c)
		default:
			fmt.Printf("other=%+v\n", c)
		}
	}
}
func main() {
	// 最基础的方式，单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 经典的初始/条件/后续 `for` 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 不带条件的 `for` 循环将一直重复执行，
	// 直到在循环体内使用了 `break` 或者 `return` 跳出循环。
	for {
		fmt.Println("loop")
		break
	}
	// 你也可以使用 `continue` 直接进入下一次循环。
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	myfunc("string", int64(1), true)
}
