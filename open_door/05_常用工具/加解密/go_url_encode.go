package main

import (
	"fmt"
	"net/url"
)

func main() {

	var urlStr string = "http://wwer.werwer.5w5er5w0.w"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:", escapeUrl)

	enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:", enEscapeUrl)
}
