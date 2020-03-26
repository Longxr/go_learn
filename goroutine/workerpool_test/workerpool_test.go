package workerpool_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

// func TestWorkerPool(t *testing.T) {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
// 	//开启3个goroutine
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	//5个任务
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	//输出结果
// 	for a := 1; a <= 5; a++ {
// 		<-results
// 	}
// }

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func workerNewInt64(j chan<- *job) {
	defer wg.Done()
	//循环生成随机数
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		j <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func workerSum(j <-chan *job, r chan<- *result) {
	defer wg.Done()
	for {
		job := <-j
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		r <- newResult
	}
}

func TestWorkerPool(t *testing.T) {
	wg.Add(1)
	go workerNewInt64(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go workerSum(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}

	wg.Wait()
}
