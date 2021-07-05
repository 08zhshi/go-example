package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("a", "b")
	fmt.Println(p)
	fmt.Println(filepath.Join("b//", "c"))
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.IsAbs(p))

	filename := "name.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
