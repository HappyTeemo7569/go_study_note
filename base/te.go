package main

import (
	"fmt"
	"sync"
)

const total = 10 // 1000000000

func worker(id int) {
	results <- id * 2
	defer wg.Done()
}

func do() {
	c := <-jobs
	worker(c)
}

func send() {
	for i := 1; i <= total; i++ {
		wg.Add(1)
		go func(i int) {
			jobs <- i
		}(i)
	}
}

var jobs = make(chan int)
var results = make(chan int)

var wg sync.WaitGroup

func main() {

	go do()
	go send()

	wg.Wait()

	var sum int32 = 0
	for r := range results {
		sum += int32(r)
	}

	fmt.Println(sum)
}
