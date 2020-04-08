package customer_type_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

type student struct {
	name string
	age  int
}

func TestRangeStruct(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for i := 0; i < len(stus); i++ {
		stu := stus[i]
		m[stu.name] = &stu
		fmt.Printf("stu address: %p, name: %v\n", &stu, stu.name)
	}

	// for _, stu := range stus {
	// 	m[stu.name] = &stu
	// 	fmt.Printf("stu address: %p, name: %v\n", &stu, stu.name)
	// }
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
