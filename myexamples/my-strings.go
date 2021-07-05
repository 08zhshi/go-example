package main

import (
	"fmt"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("|%-6.2f|%6.2f|\n", 12.345, 12.345)

	fmt.Println(strings.Split("12|er|dfe|34", "|"))
	fmt.Println(strings.Join([]string{"1", "2", "3"}, ","))
}
