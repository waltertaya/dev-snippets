// Check if Binary String Has at most one segment of ones
package main

import (
	"fmt"
	"strings"
)

/** 
Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.

	Example 1:

		Input: s = "1001"
		Output: false
		Explanation: The ones do not form a contiguous segment.
		
	Example 2:

		Input: s = "110"
		Output: true
*/
func checkOnesSegment(s string) bool {

	if strings.Contains(s, "01") {
		return false
	}

	return true
}

func main() {
	var testCases = map[string]bool{"1001": false, "110": true, "10": true, "1": true, "101": false}
	
	for k, v := range testCases {
		if checkOnesSegment(k) == v {
			fmt.Printf("Testcase {%v: %v} passed ✅\n", k, v)
		} else {
			fmt.Printf("Testcase {%v: %v} failed ❌\n", k, v)
		}
	}
}
