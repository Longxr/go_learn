package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unkonw"
	t.Log(Q2)
	t.Log(year)
}

// func TestSliceCompare(t *testing.T) {
// 	a := []int{1, 2, 3, 4}
// 	b := []int{1, 2, 3, 4}
// 	if a == b {
// 		t.Log("a==b")
// 	}
// }

func TestSliceAppend(t *testing.T) {
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("a ptr: %p\n", a)
	}
	fmt.Println(a)
}

func TestSliceAppend2(t *testing.T) {
	s := []int{1, 2, 3, 4}
	a := make([]int, 3, 6)
	b := append(a, 10)
	a[0] = 50
	fmt.Printf("a: %v\tptr: %p\tfirst: %v\n", a, a, a[0])
	fmt.Printf("b: %v\tptr: %p\tfirst: %v\n", b, b, b[0])

	b = append(a, s...)
	a[0] = 100
	fmt.Printf("a: %v\tptr: %p\tfirst: %v\n", a, a, a[0])
	fmt.Printf("b: %v\tptr: %p\tfirst: %v\n", b, b, b[0])
}

func TestSliceexpression(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	// b := a[1:3:7]
	b := a[1:3:5]
	fmt.Printf("b:%v len(b):%v cap(b):%v\n", b, len(b), cap(b))
}
