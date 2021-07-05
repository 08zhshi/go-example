package main

import "fmt"

func myval(v int) {
	fmt.Println(&v)
	v = 10
}

func myptr(p *int) {
	fmt.Println(p)
	*p = 10
}
func main() {
	i := 2
	fmt.Println(&i)
	myval(i)
	fmt.Println(i)
	myptr(&i)
	fmt.Println(i)
}
