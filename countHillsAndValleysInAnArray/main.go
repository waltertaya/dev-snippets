/**
* SOLUTION BY:- https://github.com/waltertaya
*
* Count Hills and Valleys in an Array
*/
package main

import "fmt"

/**
* An index i ia part of a hill if the closest non-equal neighbors of i are smaller than nums[i]
* An index i is part of a valley in nums if the closest non-equal neighbors of i are larger than nums[i]
* Adjacent indices i and j are part of the same hill or valley if nums[i] == nums[j]
*
* Args:
* 	- nums []int{}
*
* Rtype:
* 	- int (Return the number of hills and valleys in nums)
*/
func countHillValley(nums []int) int {
	count := 0

	for i := 1; i < len(nums)-1; i++ {
		left := nums[i-1]
		right := nums[i+1]

		// Hill
		if left < nums[i] {
			j := i + 1
			for right == nums[i] && j+1 < len(nums) {
				j++
				right = nums[j]
			}
			if right < nums[i] {
				count++
			}
		}

		// Valley
		if left > nums[i] {
			j := i + 1
			for right == nums[i] && j+1 < len(nums) {
				j++
				right = nums[j]
			}
			if right > nums[i] {
				count++
			}
		}
	}

	return count
}

func main() {
	nums := []int{2, 4, 1, 1, 6, 5}
	fmt.Println(countHillValley(nums))

	nums = []int{6, 6, 5, 5, 4, 1}
	fmt.Println(countHillValley(nums))

	nums = []int{5, 5, 1, 5, 5}
	fmt.Println(countHillValley(nums))

	nums = []int{5, 5, 1, 1, 5, 5}
	fmt.Println(countHillValley(nums))

	nums = []int{5, 1, 1, 1, 5}
	fmt.Println(countHillValley(nums))
}
