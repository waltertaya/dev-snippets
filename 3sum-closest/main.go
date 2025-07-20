/**
* Author:- https://github.com/waltertaya
* 3 sum closest
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
* Solved 3 sum using fixed i and two pointer instead of brute-force
* Complexity:- O(n-squared)
*/

func threeSumClosest(nums []int, target int) int {
	// returned sum closest
	output := nums[0] + nums[1] + nums[2]

	// sort nums slice
	sort.Ints(nums)

	// n-squared complexity
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			// logic bug:- skipping valid tripplets
			// for left < right && nums[left] == nums[left+1] {
			// 	left++
			// }
			// for left < right && nums[right] == nums[right-1] {
			// 	right--
			// }

			sum := nums[i] + nums[left] + nums[right]

			difference := sum - target
			if math.Abs(float64(difference)) < math.Abs((float64(output) - float64(target))) {
				output = sum
			}
			if difference == 0 {
				return output
			} else if difference < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return output
}

func main() {
	nums := []int{-1,2,1,-4}
	fmt.Println(threeSumClosest(nums, 1)) // expected: 2

	nums = []int{0,0,0}
	fmt.Println(threeSumClosest(nums, 1)) // expected: 0

	nums = []int{-1000,-5,-5,-5,-5,-5,-5,-1,-1,-1}
	fmt.Println(threeSumClosest(nums, -14)) // expected: -15
}
