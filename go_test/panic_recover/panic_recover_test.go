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
