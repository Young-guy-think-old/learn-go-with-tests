package main

import "fmt"


func Sum(arrayInt []int) int  {
	result := 0
	for _, number := range arrayInt {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int  {
	sum := []int{}

	for _, arrayNumber := range numbersToSum {
		sum = append(sum, Sum(arrayNumber))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int  {
	sum := []int{}
	for _, arrayNumber := range numbersToSum {
		if len(arrayNumber) > 0 {
			tail := arrayNumber[1:]
			sum = append(sum, Sum(tail))
		} else {
			sum = append(sum, 0)
		}
	}
	return sum
}

func main() {
	sums:= SumAllTails([]int{}, []int{}, []int{1,2}, []int{5,6}, []int{1,2}, []int{5,6})
	fmt.Printf("%v", sums)
}
