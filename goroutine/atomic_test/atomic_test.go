package atomic_test

import (
	"sync"
	"time"
	"sync/atomic"
	"fmt"
	"testing"
)

var (
	lock sync.Mutex
	wg sync.WaitGroup
	x int64
)

func mutexAdd() {
	defer wg.Done()

	lock.Lock()
	x++
	lock.Unlock()
}

func atomicAdd(){
	defer wg.Done()

	atomic.AddInt64(&x, 1)
}

func TestMutexAdd(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go mutexAdd()
		go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("result: %d, cost: %v", x, end.Sub(start))
}