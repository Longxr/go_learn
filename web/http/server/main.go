package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read file failed, err:%v\n", err)))
		return
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	//GET请求参数都在URL中，Body中没有数据
	fmt.Println(r.URL)
	queryParam := r.URL.Query() //自动识别URL中的query参数
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("Ok"))
}

func main() {
	http.HandleFunc("/index/", f1)
	http.HandleFunc("/hello/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
