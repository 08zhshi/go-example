package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("hello server start")
	defer fmt.Println("hello server end")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server err:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":9000", nil)
}
