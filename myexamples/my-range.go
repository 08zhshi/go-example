package main

import "fmt"

func main() {

	num := []int{1, 3, 4}
	for i, n := range num {
		fmt.Println(i, "=", n)
	}
	for i, n := range "good" {
		fmt.Printf("%d=%c\n", i, n)
	}

	ch := make(chan string, 2)
	ch <- "one"
	ch <- "two"
	close(ch)
	for elem := range ch {
		fmt.Println(elem)
	}

}
