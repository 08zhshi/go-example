package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type respone struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	str := `{"page": 12, "fruits":["apple", "pear"]}`
	var res respone
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

	r, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Marshal failed")
		return
	}
	fmt.Println(string(r))

	var mapRes map[string]interface{}
	json.Unmarshal([]byte(str), &mapRes)
	num, ok := mapRes["page"].(int64)
	if ok {
		fmt.Println(num)
	}

	f, ok := mapRes["fruits"].([]interface{})
	if ok {
		ff := f[0].(string)
		fmt.Println(ff)
	}

	B, _ := json.Marshal(true)
	fmt.Println(string(B))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
