/**
* SOLUTION BY: https://github.com/waltertaya
*
* 3487. Maximum Unique Subarray Sum After Deletion
 */
package main

import (
	"fmt"
	"sort"
)

/**
* You are allowed to delete any number of elements from `nums` without making it empty.
* After performing the deletions, select subarray of nums such that:
* 	1. All elements in the subarray are unique.
* 	2. The sum of the elements in the subarray is maximized.
* Return the maximum sum of such a subarray.
 */
func maxSum(nums []int) int {
	var subarray []int
	// base case
	if len(nums) == 1 {
		return nums[0]
	}
	// sort
	sort.Ints(nums)

	// unique subarray
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		subarray = append(subarray, nums[i])
	}
	sum := subarray[len(subarray)-1]

	// base case
	if len(subarray) == 1 {
		return subarray[0]
	}

	// max sum
	for j := len(subarray) - 2; j >= 0; j-- {
		if sum >= sum+subarray[j] {
			break
		}
		sum += subarray[j]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maxSum(nums))

	nums = []int{1, 1, 0, 1, 1}
	fmt.Println(maxSum(nums))

	nums = []int{1, 2, -1, -2, 1, 0, -1}
	fmt.Println(maxSum(nums))
}
