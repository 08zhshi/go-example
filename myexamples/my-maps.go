package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["name"] = "jion"
	fmt.Println(map1)

	map2 := map[string]string{"io": "jon", "foo": "bar"}
	fmt.Println(map2)
	v, ok := map2["io"]
	if ok {
		fmt.Println(v)
	}
	for k, v := range map2 {
		fmt.Println(k, "=", v)
	}

	//var map3 map[string]string
	//map3["nae"] = "you"  // 赋值为Panic，map3为nil。
	//fmt.Println(map3)

	delete(map2, "io")
	v1, ok1 := map2["io"]
	if ok1 {
		fmt.Println(v1)
	}
}
