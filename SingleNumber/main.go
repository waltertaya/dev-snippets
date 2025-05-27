package main

import (
	"fmt"
)

func main() {
	// test case 1
	test := []int{2, 2, 1}
	fmt.Println(singleNumber(test))

	// test case 2
	test = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(test))

	// test case 3
	test = []int{1}
	fmt.Println(singleNumber(test))
}

func singleNumber(nums []int) int {
	singleMap := make(map[int]bool)
	result := 0

	for _, val := range nums {
		if !singleMap[val] {
			singleMap[val] = true
		} else {
			delete(singleMap, val)
		}
	}
	for k, _ := range singleMap {
		result = k
	}

	return result
}
