package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer type, ", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string type, ", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer type, ", v)
	case string:
		fmt.Println("String type, ", v)
	default:
		fmt.Println("Unknow type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

func TestEmptyMap(t *testing.T) {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "张三"
	m1["age"] = 18
	m1["merrid"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
}
