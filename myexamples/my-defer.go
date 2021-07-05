package main

import (
	"fmt"
	"os"
)

func creatFile(p string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
func main() {
	fn := "defer.txt"
	f := creatFile(fn)
	defer closeFile(f)
	writeFile(f)
	//defer func() {
	//	v := recover()
	//	fmt.Println("3",v)
	//}()
	defer func() {
		defer func() {
			v := recover()
			fmt.Println("2", v)
		}()
		defer func() {
			defer func() {
				fmt.Println("1")
			}()
			panic("test")
		}()
	}()
	panic("main")

}
