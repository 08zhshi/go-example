package main

import (
	"fmt"
	"net"
	url2 "net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path/sub?k=v&u=h#f"
	u, err := url2.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url2.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
