package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
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
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Started")
	panic("Fatal error")
	t.Log("End")
}

func addSub(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func TestBibao(t *testing.T) {
	f1, f2 := addSub(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

func ret1() int {
	x := 1
	defer func() {
		x++
	}()
	return x
}

func ret2() (x int) {
	x = 1
	defer func() {
		x++
	}()
	return x
}

func TestRet(t *testing.T) {
	t.Log("ret1(): ", ret1())
	t.Log("ret2(): ", ret2())
}

func calc(index, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func TestDefer2(t *testing.T) {
	a := 1
	b := 2
	defer calc(1, a, calc(2, a, b))
	a = 0
	defer calc(3, b, calc(4, a, b))
	b = 1
}

// defer calc(1, 1, 3)
// defer calc(3, 2, 2)
// 2, 1, 2, 3
// 4, 0, 2, 2
// 3, 2, 2, 4
// 1, 1, 3, 4
