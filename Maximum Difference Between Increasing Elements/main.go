package main

func maximumDifference(nums []int) int {
    results := -1
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] >= nums[j] {
                continue
            }
            temp := nums[j] - nums[i]
            if temp >= results {
                results = temp
            }
        }
    }

    return results
}

func main() {
	nums := []int{7, 1, 5, 4}
	result := maximumDifference(nums)
	println(result)
}
