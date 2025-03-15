package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i] - 'a']++
		}
		m[k] = append(m[k], s)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs2 := []string{""}
	strs3 := []string{"a"}

	fmt.Println(groupAnagrams(strs1))
	fmt.Println(groupAnagrams(strs2))
	fmt.Println(groupAnagrams(strs3))
}
