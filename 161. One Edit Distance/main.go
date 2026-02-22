package main

import "fmt"

func isOneEditDistance(s, t string) bool {
	ns, nt := len(s), len(t)

	if ns > nt {
		return isOneEditDistance(t, s)
	}

	if nt-ns > 1 {
		return false
	}

	for i := range s {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	return ns+1 == nt
}

func main() {
	{
		s := "ab"
		t := "acb"

		fmt.Println(isOneEditDistance(s, t) == true)
	}
	{
		s := ""
		t := ""

		fmt.Println(isOneEditDistance(s, t) == false)
	}
}
