package select_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		for i := 0; i < numOfRunner; i++ {
			go func(i int) {
				ret := runTask(i)
				ch <- ret
			}(i)
		}
	}
	return <-ch
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		for i := 0; i < numOfRunner; i++ {
			go func(i int) {
				ret := runTask(i)
				ch <- ret
			}(i)
		}
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	// t.Log(FirstResponse())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
