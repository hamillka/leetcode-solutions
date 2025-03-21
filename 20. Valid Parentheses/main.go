package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	m := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	for _, el := range s {
		if _, ok := m[el]; ok {
			stack = append(stack, el)
		} else if len(stack) == 0 || m[stack[len(stack)-1]] != el {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "(]"
	s4 := "([])"

	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
	fmt.Println(isValid(s4))
}
