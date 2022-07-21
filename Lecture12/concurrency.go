package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg.Add(len(n))
	for ind, val := range n {
		value := val
		count := ind + 1
		go func() {
			sum(count, value)
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func sum(count int, sls []int) {
	var sumOfSlice int

	for _, num := range sls {
		sumOfSlice += num
	}
	fmt.Printf("slise %d: %d\n", count, sumOfSlice)
}
