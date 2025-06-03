package palindrome

import (
    "strings"
)

// IsPalindrome checks whether a string is a palindrome (ignoring spaces and case)
func IsPalindrome(s string) bool {
    cleaned := strings.ReplaceAll(strings.ToLower(s), " ", "")
    length := len(cleaned)

    for i := 0; i < length/2; i++ {
        if cleaned[i] != cleaned[length-1-i] {
            return false
        }
    }
    return true
}
