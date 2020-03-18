package series

import "fmt"

func init() {
	fmt.Println("Init1")
}

func init() {
	fmt.Println("Init2")
}

func Square(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
