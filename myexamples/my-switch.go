package main

import (
	"fmt"
	"reflect"
)

func main() {

	f := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		default:
			fmt.Println(reflect.ValueOf(i).Kind())
		}
	}
	f(true)
	f("dasd")
	f(1)
}
