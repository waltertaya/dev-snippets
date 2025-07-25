/**
* AUTHOR: https://github.com/waltertaya
*
* Variable-size sliding window
 */
package main

import "fmt"

/**
* Find the longest substring without repeating characters
* Args:
* 	- s string
* Rtype:
* 	- integer
 */
func longestUniqueSubstring(s string) int {
	seen := make(map[rune]bool)
	left := 0
	maxLen := 0
	runes := []rune(s)

	for right := 0; right < len(runes); right++ {
		for seen[runes[right]] {
			delete(seen, runes[left])
			left++
		}
		seen[runes[right]] = true
		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(longestUniqueSubstring(s))
}
