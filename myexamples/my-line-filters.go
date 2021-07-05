package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 将输入字符串，将每行字符串变成大写，直到碰到EOF结束。
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
