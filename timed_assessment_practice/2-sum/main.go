package main

import (
	"fmt"
)

// 14-minute
func twoSum (nums []int, target int) []int {
	var hash = make(map[int]int)
	var results []int

	for i := 0; i < len(nums); i++ {
		var compliment = target - nums[i]

		if idx, ok := hash[compliment]; ok {
			results = append(results, idx, i)
			break
		}

		hash[nums[i]] = i
	}
	
	return results
}


func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	results := twoSum(nums1, target1)
	fmt.Println("output:", results, "expected: [0, 1]")
}
