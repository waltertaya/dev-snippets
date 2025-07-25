/**
* SOLUTION BY: https://github.com/waltertaya
* Fixed sized sliding window
 */
package main

import "fmt"

/**
* Given a list numbers,
* find the maximum total of any k numbers next to each other
* 
* Args:
* 	- nums []int
* 	- k int
*
* Rtype:
* 	- int
 */
func maxSumSubarray(nums []int, k int) int {
	maxSum := 0
	windowSum := 0

	for i := 0; i < len(nums); i++ {
		windowSum += nums[i]

		if i >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= nums[i-(k-1)]
		}
	}
	return maxSum
}

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	fmt.Println(maxSumSubarray(nums, 3))
}
