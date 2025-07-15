/**
* Solution by https://github.com/waltertaya
* 3 Sum code challenge
*/

package main

import (
	"fmt"
	"sort"
)

/*
* Two-pointer solution with i being fixed
* Optimal solution
* Faster than Brute force and solve the triplet duplicate problem
*/

func threeSum(nums []int) [][]int {
	var results = [][]int{}

	// sort the slice
	sort.Ints(nums)

	// [-4, -1, -1, 0, 1, 2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				
				for left < right && nums[left] == nums[left + 1] {
					left++
				}
				for left < right && nums[right] == nums[right - 1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return results
}

/*
* 1st Attempt, Brute Force Approach
* Try all combinations of 3 numbers using 3 nested loops
* Check the sum is zero
*
* PROBLEM:
*		- This is O(n3) - inefficient for large arrays
*		- Duplicates triplets
 */

// func threeSum(nums []int) [][]int {
// 	var results = [][]int{}

// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				sum := nums[i] + nums[j] + nums[k]
// 				if sum == 0 {
// 					results = append(results, []int{nums[i], nums[j], nums[k]})
// 				}
// 			}
// 		}
// 	}

// 	return results
// }

func main() {
	testcase := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(testcase))
}
