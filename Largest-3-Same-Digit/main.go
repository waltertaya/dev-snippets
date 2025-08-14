package main

import (
	"fmt"
	"strconv"
)

/**
* Solution by: https://github.com/waltertaya
*
* Largest 3-Same-Digit Number in String
 */

func main() {
	num := "6777133339"
	fmt.Println(largestGoodInteger(num))
	num = "2300019"
	fmt.Println(largestGoodInteger(num))
	num = "42352338"
	fmt.Println(largestGoodInteger(num))
}

// fixed (3) sliding window solution

/**
* Given a string `num` representing a large integer.
* Integer is good if it meets the following conditions:
* 	- It is a substring of `num` with length `3`.
* 	- It consists of only one unique digit.
* Substring is a contiguous sequence of characters within a string
* There may be leading zeroes in `num` or a good integer.
*
* args:-
* 	- `num` string
*
* rtype:-
* 	- string
 */

func largestGoodInteger(num string) string {
	var maxGoodInt string
	var window string

	for i := 0; i < len(num); i++ {
		window += string(num[i])

		if i >= 2 {
			if window[0] != window[1] || window[0] != window[2] || window[1] != window[2] {
				window = window[1:]
				continue
			}
			if maxGoodInt == "" {
				maxGoodInt = window
			}
			windowInt, _ := strconv.Atoi(window)
			goodInt, _ := strconv.Atoi(maxGoodInt)
			if windowInt > goodInt {
				maxGoodInt = window
			}
			window = window[1:]
		}
	}

	return maxGoodInt
}
