package main

/**
* calculate the expected sum and substract the values available
* remainder gives the correct missing
*/

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	expected := (n * (n + 1)) / 2

	for _, num := range nums {
		expected -= num
	}

	return expected
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

