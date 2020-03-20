package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//公用一个client 适用于比较频繁的请求
// var (
// 	client = http.Client {
// 		Transport: &&http.Transport{
// 			DisableKeepAlives: true,
// 		},
// 	}
// )

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/hello/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Printf("get url failed, err:%v\n", err)
	// 	return
	// }
	urlObj, _ := url.Parse("http://127.0.0.1:9090/hello/")
	data := url.Values{}
	data.Set("name", "张三")
	data.Set("age", "18")
	queryStr := data.Encode() //编码转义后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close() //请求使用完后要关闭

	//读取请求数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
