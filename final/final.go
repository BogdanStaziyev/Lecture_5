package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	ch := make(chan int, len(n))
	var sum int

	for _, val := range n {
		go func() {
			sumOfValues(val, ch)
		}()
		sum += <-ch
	}
	close(ch)
	fmt.Println("result:", sum)
	//мені здається що реалізація невірна)) але через брак досвіду не можу висловити проблему
}

func sumOfValues(sls []int, ch chan int) {
	var sumOfSlice int

	for _, num := range sls {
		sumOfSlice += num
	}
	ch <- sumOfSlice
}
