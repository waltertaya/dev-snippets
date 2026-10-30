package main

import "fmt"

func containsDuplicate(nums []int) bool {
	var seen = make(map[int]bool)

	for _, val := range nums {
		if _, ok := seen[val]; ok {
			return true
		}

		seen[val] = true
	}

	return false
}

func main() {
	var nums = []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(nums))
}
