package recursion_test

import (
	"fmt"
	"testing"
)

func jiecheng(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * jiecheng(n-1)
}

func TestJiecheng(t *testing.T) {
	ret := jiecheng(5)
	fmt.Println(ret)
}

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func TestTaijie(t *testing.T) {
	ret := taijie(4)
	fmt.Println(ret)
}
