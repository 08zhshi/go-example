package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("12")
	if err == nil {
		fmt.Println(i)
	}
	str := strconv.Itoa(456)
	fmt.Println(str)

	// base参数处理10进制为10，
	i64, err := strconv.ParseInt("12", 10, 0)
	if err == nil {
		fmt.Println(i64)
	}
	i64, err = strconv.ParseInt("0xc8", 16, 0)
	if err == nil {
		fmt.Println(i64)
	}
	var bi []byte
	bi = strconv.AppendInt(bi, 345, 10)
	fmt.Println(string(bi)) // 345
	var bi1 []byte
	bi1 = strconv.AppendInt(bi1, 80, 16)
	fmt.Println(string(bi1)) // 50

	f, err := strconv.ParseFloat("13.44", 64)
	if err == nil {
		fmt.Println(f)
	}
	var ft []byte
	ft = strconv.AppendFloat(ft, 13.5424, 'e', 2, 32)
	fmt.Println(string(ft))
}
