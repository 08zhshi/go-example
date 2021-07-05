package main

import (
	"flag"
	"fmt"
	"os"
)

func testArgs() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

func testFlags() {
	wordPtr := flag.String("word", "fou", "a string")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("svar:", svar)
}

func testEnv() {
	os.Setenv("boo", "you")
	fmt.Println(os.Getenv("boo"))
	fmt.Println(os.Getenv("bar"))
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

func main() {
	//testArgs()
	//testFlags()
	testEnv()
}
