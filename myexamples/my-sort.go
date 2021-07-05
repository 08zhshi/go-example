package main

import (
	"fmt"
	"sort"
)

type mySort struct {
	Name string
	Age  int
}

type mySortS []mySort

func (m mySortS) Len() int {
	return len(m)
}
func (m mySortS) Less(i, j int) bool {
	return m[i].Age < m[j].Age
}
func (m mySortS) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	s := []string{"q", "c", "d"}
	sort.Strings(s)
	fmt.Println(s)

	var aa mySortS
	aa = append(aa, []mySort{{Name: "you", Age: 12}, mySort{Name: "me", Age: 11}, mySort{Name: "she", Age: 14}}...)
	sort.Sort(aa)
	fmt.Println(aa)
}
