package prime_test

import (
	"fmt"
	"testing"
)

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func TestPrime(t *testing.T) {
	out := GenerateNatural()
	//in := make(chan int)
	//out := in
	//go func() {
	//	for i := 2; ; i++ {
	//		in <- i
	//	}
	//}()
	for i := 0; i < 100; i++ {
		prime := <-out // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		fmt.Printf("out ptr: %v\n", out)
		out = PrimeFilter(out, prime) // 基于新素数构造的过滤器
	}
}