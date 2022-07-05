package main

import "fmt"

func main() {
	//Recursion (Ркурсія) + для практики рекурсію пропоную таку задачу:
	//написати функцію з рекурсією, що повертає суму елементів масива
	var res int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		res = makeRecursive(arr[i])
	}
	fmt.Println(res)
}

func makeRecursive(num int) int {
	if num == 1 {
		return 1
	}
	res := num + makeRecursive(num-1)
	return res
}

//fourth FUNC
//	num := 3
//	slises := make([][]int, num)
//	firstnumber := 1
//
//	for countOfSlise := 1; countOfSlise <= num; countOfSlise++ {
//		slise := make([]int, num)
//		firstnumber = countOfSlise
//		for numbersCount := 0; numbersCount < num; numbersCount++ {
//
//			slise[numbersCount] = firstnumber
//			firstnumber+=countOfSlise
//		}
//		slises[countOfSlise-1] = slise
//	}
//	fmt.Println(slises)
//}

//third result
//	var s = "aa b"
//	stringLise := strings.Split(s, " ")
//	numbers := 0
//	var newstring string
//	for _, num := range stringLise{
//		if numbers < summOfNumber(num){
//			numbers=summOfNumber(num)
//			newstring = num
//		}
//	}
//	fmt.Println(newstring)
//}
//
//func summOfNumber(s string) int{
//	count := 1
//	summ := 0
//	for i:='a'; i<='z'; i++{
//		for _, num := range s{
//			if i==num{
//				summ+=count
//			}
//		}
//		count++
//	}
//	return summ
//}

//second result
//func main() {
//	firstStep := getRes("aabccc")
//	printNumber(firstStep)
//}
//
//func printNumber(r []rune) {
//	var num rune
//	num = 0
//		for i := 'a'; i <= 'z'; i++ {
//			for _, run := range r {
//			if i == run {
//				num = i
//				break
//			}
//		}
//		if num != 0 {
//		break
//		}
//	}
//	fmt.Println(num)
//}
//
//func getRes(s string) []rune{
//	var result int
//	var resStr []rune
//	for _, res := range s {
//		min := strings.Index(s, string(res))
//		max := strings.LastIndex(s, string(res))
//		num := max - min
//		if result < num {
//			result = num
//			resStr = nil
//			resStr = append(resStr, res)
//			} else if result == num {
//				resStr = append(resStr, res)
//			}
//		}
//	return resStr
//}
