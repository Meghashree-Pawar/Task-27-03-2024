package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func squareWorker(nums []int, ch chan int) {
	defer wg.Done()
	for _, num := range nums {
		ch <- num * num
	}
	close(ch)
}

func aggregateSquares(ch chan int) {
	defer wg.Done()
	sum := 0
	for square := range ch {
		sum += square
	}
	fmt.Println("Aggregated squared result:", sum)
}

func main() {
	nums := []int{1, 2, 3,4,5}
	ch := make(chan int)

	wg.Add(2)
	go squareWorker(nums, ch)
	go aggregateSquares(ch)

	wg.Wait()
}
