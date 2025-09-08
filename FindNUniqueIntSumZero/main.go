/**
* Solution by https://github.com/waltertaya
* Find N Unique Integers Sum upto Zero
 */

package main

import "fmt"

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(4))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(2))
	fmt.Println(sumZero(1))
}

/**
* Given an integer n, return any array containing n unique integers such that they add up to 0
 */
func sumZero(n int) []int {
	var result []int
	// base case n = 1
	if n == 1 {
		result = append(result, 0)
		return result
	}
	if n%2 != 0 { // add zero centroid for the n is odd number
		result = append(result, 0)
		n -= 1
	}
	for i := 1; i <= n/2; i++ {
		result = append(result, i)
		result = append(result, -(i)) // cancel the +ve i appended
	}

	return result
}
