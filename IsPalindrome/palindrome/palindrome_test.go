package palindrome_test

import (
	"testing"

	"github.com/waltertaya/dev-snippets/palindrome"
	"fmt"
)

// Basic Unit Tests: go test -v
func TestIsPalindromeTrueBasic(t *testing.T) {
	ans := palindrome.IsPalindrome("racecar")

	if !ans {
		t.Errorf("IsPalindrome(%v) Expected = %v", "racecar", true)
	}
}

func TestIsPalindromeFalseBasic(t *testing.T) {
	ans := palindrome.IsPalindrome("Hello")

	if ans {
		t.Errorf("IsPalindrome(%v) Expected = %v", "Hello", false)
	}
}

// Table Driven Tests: go test -v
func TestIsPalindromeTableDriven(t *testing.T) {
	tests := []struct{
		s string
		expected bool
	} {
		{"Was it a car or a cat I saw", true},
		{"No lemon no melon", true},
		{"Go is awesome", false},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.s)

		t.Run(testname, func(t *testing.T) {
			ans := palindrome.IsPalindrome(test.s)
			if ans != test.expected {
				t.Errorf("IsPalindrome(%v) Expected = %v", test.s, test.expected)
			}
		})
	}
}

// benchmark tests for performance: go test -bench=.
func BenchmarkIsPalindrome(b *testing.B) {
    input := "Was it a car or a cat I saw"

    for i := 0; i < b.N; i++ {
        palindrome.IsPalindrome(input)
    }
}

func BenchmarkIsPalindromeVarious(b *testing.B) {
    cases := []string{
        "madam",
        "Was it a car or a cat I saw",
        "Go is fun",
        "Never odd or even",
    }

    for _, input := range cases {
        b.Run(input, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                palindrome.IsPalindrome(input)
            }
        })
    }
}

