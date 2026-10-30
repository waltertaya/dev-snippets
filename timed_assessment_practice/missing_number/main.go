package main

/**
*	Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return
*	the only number in the range that is missing from the array
*/

import "fmt"

func missingNumber(nums []int) int {
	// use index marking by computing and replacing existing with [-1, -(n+1)] => [0, n]
	var n = len(nums)
	var result = n

	for _, val := range nums {
		if val < 0 {
			// decode back
			val = -(val + 1)
		}
		if val == n {
			continue
		}

		nums[val] = -(nums[val] + 1)
	}

	for i, val := range nums {
		if val >= 0 {
			result = i
			break
		}
	}

	return result
}


func main() {
	var nums = []int{1, 0}
	fmt.Println("got:", missingNumber(nums), "expect: 2")

	nums = []int{3, 0, 1}
	fmt.Println("got:", missingNumber(nums), "expect: 2")

	nums = []int{0, 1}
	fmt.Println("got:", missingNumber(nums), "expect: 2")

	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("got:", missingNumber(nums), "expect: 8")

	nums = []int{2, 0}
	fmt.Println("got:", missingNumber(nums), "expect: 1")

	nums = []int{1, 2}
	fmt.Println("got:", missingNumber(nums), "expect: 0")
}

