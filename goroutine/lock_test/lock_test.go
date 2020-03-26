package lock_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	x      = 0
	lock   sync.Mutex
	rwlock sync.RWMutex
	wg     sync.WaitGroup
)

func write() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	// lock.Unlock()
	rwlock.Unlock()
}

func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.RLock()
	// fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.RUnlock()
}

func TestLock(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Printf("result:%d, cost:%v", x, time.Now().Sub(start))
}
