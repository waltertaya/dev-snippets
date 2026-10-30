package main

/**
* XOR has two properties: a xor a = 0
*		a xor 0 = a
* So the xor nums value the required with cancel each leaving result uncancelled from required
*/
import "fmt"


func missingNumber(nums []int) int {
	var result = len(nums)

	for i, num := range nums {
		result ^= i ^ num
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
