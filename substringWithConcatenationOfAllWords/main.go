/**
* SOLUTION BY:- https://github.com/waltertaya
*
* Substring with Concatenation of All Words
*/
package main

import (
	"fmt"
	"slices"
	"sort"
)

func findSubstring(s string, words []string) []int {
	result := []int{}
	wordLen := len(words[0])
	wordCount := len(words)
	sLen := len(s)
	windowSize := wordLen * wordCount

	// base case scenario
	if sLen < windowSize {
		return result
	}

	sortedWords := make([]string, wordCount)
	copy(sortedWords, words)
	sort.Strings(sortedWords)

	// sliding window soln
	for i := 0; i <= sLen-windowSize; i++ {
		var windowString []string
		for j := 0; j < wordCount; j++ {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			windowString = append(windowString, word)
		}

		sortedWindowString := make([]string, wordCount)
		copy(sortedWindowString, windowString)
		sort.Strings(sortedWindowString)

		if slices.Equal(sortedWindowString, sortedWords) {
			result = append(result, i)
		}
		windowString = windowString[1:]

	}

	return result
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))

	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s, words))

	s = "barfoofoobarthefoobarman"
	words = []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s, words))
}
