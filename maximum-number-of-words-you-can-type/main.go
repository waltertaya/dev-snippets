/**
* Solution by: https://github.com/waltertaya
* 
* Maximum number of words you can type
*/

package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text, brokenLetters string) int {
	var count int

	// break the text into word slice
	textSlice := strings.Split(text, " ")

	// base case
	if len(brokenLetters) == 0 {
		return len(textSlice)
	}

	for _ , word := range textSlice {
		if strings.ContainsAny(word, brokenLetters) {
			continue
		}
		count++
	}

	return count
}

func main() {
	text, brokenLetters := "hello world", "ad"
	fmt.Println(canBeTypedWords(text, brokenLetters))

	text, brokenLetters = "leet code", "lt"
	fmt.Println(canBeTypedWords(text, brokenLetters))

	text, brokenLetters = "leet code", "e"
	fmt.Println(canBeTypedWords(text, brokenLetters))
}
