package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	m := make(map[int]int)
	m[arr[0]] = arr[0]
	result = append(result, arr[0])
	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = m[val]
			result = append(result, val)
		}
	}
	fmt.Println(result)
}
