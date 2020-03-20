package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	u1 := User{
		Name: "longxr",
		Gender: "男",
		Age: 23,
	}
	m1 := map[string]interface{} {
		"name": "long",
		"gender": "男",
		"age": 24,
	}
	hobbyList := []string{
		"唱",
		"跳",
		"rap",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func f1(w http.ResponseWriter, r *http.Request) {
	//定义函数
	k := func(name string) (string, error) {
		return name + "真帅", nil
	}

	//定义模板
	t := template.New("f.tmpl")
	//注册自定义函数
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "longxr"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func tmplDemo1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "longxr"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "longxr"
	//err = t.Execute(w, name)
	err = t.ExecuteTemplate(w, "home.tmpl", name)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	name := "longxr"
	//err = t.Execute(w, name)
	err = t.ExecuteTemplate(w, "index.tmpl", name)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/f1", f1)
	http.HandleFunc("/tmplDemo1", tmplDemo1)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed, err:%v", err)
		return
	}
}
