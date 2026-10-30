package main

// index marking for the existing numbers

import (
	"fmt"
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	var results []int

	for _, val := range nums {
		var idx = int(math.Abs(float64(val)) - 1)

		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}

	for i, num := range nums {
		if num > 0 {
			results = append(results, i+1)
		}
	}

	return results
}


func main() {
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDisappearedNumbers(nums)) // expect: [5,6]

	nums1 := []int{1,1}
	fmt.Println(findDisappearedNumbers(nums1)) // [2]
}
