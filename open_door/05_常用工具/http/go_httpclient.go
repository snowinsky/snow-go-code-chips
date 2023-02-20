package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//get
	getUrl := "https://ddap.uat.homecreditcfc.cn/health"
	getResp, getErr := http.Get(getUrl)
	if nil != getErr {
		panic("get fail")
	}
	defer getResp.Body.Close()
	getRespBody, _ := io.ReadAll(getResp.Body)
	getRespBodyString := string(getRespBody)
	fmt.Println("get url response body string=", getRespBodyString)

	getUrlWithBasicAuth := "https://sss"
	urlQuery := url.Values{}
	urlQuery.Add("login", "abc")
	urlQuery.Add("passw", "pas")
	getReqWithBasicAuth, getReqErr := http.NewRequest(http.MethodGet, getUrlWithBasicAuth, strings.NewReader(urlQuery.Encode()))
	if nil != getReqErr {
		panic("new get req with basic auth fail")
	}
	getReqHeader := getReqWithBasicAuth.Header
	getReqHeader.Add("Content-Type", "application/json")
	getReqWithBasicAuth.SetBasicAuth("userName", "password")

	getResWithAuth, getResErr := http.DefaultClient.Do(getReqWithBasicAuth)
	if nil != getResErr {
		panic("get res with basic auth fail")
	}
	defer getResWithAuth.Body.Close()
	getResWithAuthBody, _ := io.ReadAll(getResWithAuth.Body)
	getResWithAuthBodyString := string(getResWithAuthBody)
	fmt.Println("get url response body string=", getResWithAuthBodyString)

	//postJson

	//postForm
}
