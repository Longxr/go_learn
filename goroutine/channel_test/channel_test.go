package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //放满100个值就关闭通道
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	// for {
	// 	x, ok := <-ch1
	// 	if !ok {
	// 		break
	// 	}
	// 	ch2 <- x * x //多个发送者，如果其中一个关了通道会panic
	// }

	for x := range ch1 {
		ch2 <- x * x
	}

	// once.Do(func() { close(ch2) }) //保证关闭ch2只会调用一次
}

func TestChannel(t *testing.T) {
	a := make(chan int, 50)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	close(b)
	for ret := range b {
		fmt.Println(ret)
	}
}
