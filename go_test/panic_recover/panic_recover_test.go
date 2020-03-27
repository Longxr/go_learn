package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
		if err := recover(); err != nil {
			fmt.Println("recoverd from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	fmt.Println("Panic after")
	// os.Exit(-1)
}

func TestPanicDefer(t *testing.T) {
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer 2")
	}()

	defer func() {
		fmt.Println("defer 3")
	}()

	panic("触发异常")

	defer func() {
		fmt.Println("defer 4")
	}()
}
