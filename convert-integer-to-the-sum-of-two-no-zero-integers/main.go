/**
* Solution by https://github.com/waltertaya
* Convert Integer to the Sum of Two No-Zero Integers
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* Given an integer n, return a list of two integers [a, b] where:
*
* 	a and b are No-Zero integers.
* 	a + b = n
 */

func getNoZeroIntegers(n int) []int {
	var results []int
	i, j := 1, n-1

	for i <= j {
		if !strings.Contains(strconv.Itoa(i), "0") && !strings.Contains(strconv.Itoa(j), "0") {
			results = append(results, i, j)

			break
		}
		i++
		j--
	}

	return results
}

// Test case
func main() {
	fmt.Println(getNoZeroIntegers(11))
	fmt.Println(getNoZeroIntegers(2))
}
