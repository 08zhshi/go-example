package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(2)
	fmt.Println(rand.Intn(10))
	//rand.Seed(1)
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float32())

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println(r.Intn(10))
}
