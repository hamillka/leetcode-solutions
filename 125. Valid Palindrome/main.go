package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	f := func(c rune) rune {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			return -1
		}

		return unicode.ToLower(c)
	}
	s = strings.Map(f, s)

	beg, end := 0, len(s) - 1
	for beg < end {
		if s[beg] != s[end] {
			return false
		}
		beg++
		end--
	}

	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "

	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
	fmt.Println(isPalindrome(s3))
}