package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// for i := 0; i < 11; i++ {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel close")
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

func cancel(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Chanceled")
		}(i, cancelChan)
	}
	cancel(cancelChan)
	time.Sleep(time.Second * 1)
}
