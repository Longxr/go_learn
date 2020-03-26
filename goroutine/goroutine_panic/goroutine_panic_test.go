package goroutine_panic_test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("主协程 defer")

		if r := recover(); r != nil {
			fmt.Println("主协程恢复")
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("子协程 defer")
			if r := recover(); r != nil {
				fmt.Println("子协程恢复")
			}
			wg.Done()
		}()

		panic("子线程出错了")
	}()
	// panic("主线程出错了")
	wg.Wait()
	fmt.Println("主程序结束")
}
