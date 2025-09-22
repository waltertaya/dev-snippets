package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	var totalMaximum, maxFrequency int
	frequency := map[int]int{}

	// find frequencies of all elements of the array
	for _, val := range nums {
		frequency[val]++
		if frequency[val] > maxFrequency {
			maxFrequency = frequency[val]
		}
	}

	// find elements that have maximum frequencies and count their total occurrences
	for _, val := range frequency {
		if val == maxFrequency {
			totalMaximum += val
		}
	}

	return totalMaximum
}

func main() {
	// testcase

	nums := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(maxFrequencyElements(nums)) // 4

	nums = []int{1, 2, 3, 4, 5}
	fmt.Println(maxFrequencyElements(nums)) // 5
}
