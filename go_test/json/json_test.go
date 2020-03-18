package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func TestJson(t *testing.T) {
	p1 := person{
		Name: "张三",
		Age:  18,
	}

	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	//反序列化
	str := `{"name":"李四","age":19}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
