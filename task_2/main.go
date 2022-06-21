package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 2 3 4 5"
	var result string
	var min, max int32

	inputSlip := strings.Split(input, " ")

	m, _ := strconv.Atoi(inputSlip[0])
	min = int32(m)

	for _, val := range inputSlip {
		i, _ := strconv.Atoi(val)
		if min >= int32(i) {
			min = int32(i)
		}
		if max <= int32(i) {
			max = int32(i)
		}
	}
	minString := strconv.Itoa(int(min))
	maxString := strconv.Itoa(int(max))
	result += maxString + " " + minString
	fmt.Println(result)
}
