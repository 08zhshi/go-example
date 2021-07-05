package main

import (
	"fmt"
)

type myerror struct {
	Err  string
	Code int32
}

func NewMyError(code int32, err string) error {
	return &myerror{Err: err, Code: code}
}

func (e *myerror) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Err)
}

func main() {
	err := NewMyError(12, "Not Found")

	if err != nil {
		fmt.Println(err.Error())
	}
	switch err.(type) {
	case *myerror:
		fmt.Println("*myerror")
	default:
		fmt.Println("undefined")
	}
}
