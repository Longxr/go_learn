package syncmap_test

import (
	"sync"
	"testing"
	"strconv"
)

var c M

// M
type M struct {
    // Map    map[string]int
    Map    sync.Map
}

// Set ...
func (m *M) Set(key string, value int) {
	// m.Map[key] = value
	m.Map.Store(key, value)
}

// Get ...
func (m *M) Get(key string) int {
	// return m.Map[key]
	v, _ := m.Map.Load(key)
	value, _ := v.(int)
	return value
}

func TestSyncMap(t *testing.T) {
	// c = M{Map: make(map[string]int)}
	c = M{Map: sync.Map{}}
    wg := sync.WaitGroup{}
    for i := 0; i < 21; i++ {
        wg.Add(1)
        go func(n int) {
            k := strconv.Itoa(n)
            c.Set(k, n)
            t.Logf("k=:%v,v:=%v\n", k, c.Get(k))
            wg.Done()
        }(i)
    }
    wg.Wait()
    t.Log("ok finished.")
}