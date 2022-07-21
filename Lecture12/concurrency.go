package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	var wg sync.WaitGroup

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg.Add(len(n))
	for _, val := range n {
		value := val
		go func() {
			sum(value)
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func sum(sls []int) {
	var sumOfSlice int
	count++

	for _, num := range sls {
		sumOfSlice += num
	}
	fmt.Printf("slise %d: %d\n", count, sumOfSlice)
}
