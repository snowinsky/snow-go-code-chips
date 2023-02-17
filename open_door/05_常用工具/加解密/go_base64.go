package main

import (
	"encoding/base64"
	"fmt"
)

func base64UrlEncode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func base64UrlDecode(data string) string {
	a, b := base64.URLEncoding.DecodeString(data)
	if nil != b {
		panic("base64 url decode fail")
	}
	return string(a)
}

func base64StdEncode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func base64StdDecode(data string) string {
	a, b := base64.StdEncoding.DecodeString(data)
	if nil != b {
		panic("base64 standard decode fail")
	}
	return string(a)
}

func main() {

	fmt.Println("http://234.234.com", base64StdEncode("http://234.234.com"), base64StdDecode(base64StdEncode("http://234.234.com")))
	fmt.Println("http://234.234.com", base64UrlEncode("http://234.234.com"), base64UrlDecode(base64UrlEncode("http://234.234.com")))

}
