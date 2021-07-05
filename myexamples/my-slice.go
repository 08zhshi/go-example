package main

import "fmt"

func main() {
	s := make([]string, 2)
	s[0] = "a"
	s[1] = "b"
	fmt.Printf("len=%d, cap=%d,&s=%p,s=%p\n", len(s), cap(s), &s, s)
	s = append(s, "d")
	fmt.Printf("len=%d, cao=%d,&s=%p,s=%p\n", len(s), cap(s), &s, s)

	c := make([]string, len(s))
	copy(c, s)
	c[0] = "aa"
	fmt.Printf("c=%v,s=%v\n", c, s)

	fmt.Println(s[:0])
	fmt.Println(s[:2])
	fmt.Println(s[1:])

	l := s[1:]
	l[0] = "bb"
	fmt.Printf("l=%p,l=%v,s=%p,s=%v\n", l, l, s, s)
}
