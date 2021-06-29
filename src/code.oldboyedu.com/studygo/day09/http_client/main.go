package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

func main() {
	//resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=今天&age=18")
	//if err != nil {
	//	fmt.Println("get url failed, err:%v\n", err)
	//	return
	//}
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx")
	data.Set("name", "今天")
	data.Set("age", "18")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	//resp , err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return

	// 禁用KeepAlive的 client(适用于请求不是特别频繁时)
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close() // 关闭resp
	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body")
		return
	}
	fmt.Println(string(b))
}
