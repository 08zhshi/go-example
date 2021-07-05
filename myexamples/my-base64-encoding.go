package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "sdf23@#$%&*(')=-+http://ow.orga?q=1&b=3"
	// 加密
	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)
	// 解密
	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(string(dec))

	urlEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEnc)
	urlDec, _ := base64.URLEncoding.DecodeString(urlEnc)
	fmt.Println(string(urlDec))
}
