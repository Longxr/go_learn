package struct_combination_test

import (
	"fmt"
	"testing"
)

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address
}

type company struct {
	name string
	address
}

func TestCombination(t *testing.T) {
	p1 := person{
		name: "张三",
		age:  18,
		address: address{
			province: "江苏",
			city:     "南京",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.city)
}
