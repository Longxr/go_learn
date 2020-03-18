package nilfun_test

import (
	"fmt"
	"testing"
)

type Dog struct {
	name string
}

func (d *Dog) speak() {
	fmt.Println("wangwang~")
}

func callSpeak() {
	var d *Dog
	defer d.speak()
	if d == nil {
		fmt.Println("d is nil")
		return
	}
	// defer d.speak()
}

func TestNilFunc(t *testing.T) {
	callSpeak()
}