/*
массив целых чисел и число k
вернуть k наиболее встречающихся по частоте

[1 1 1 2 2 3]
k = 2
---
res = [1 2]
*/

package main

import (
	"fmt"
	"sort"
)

func solution(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, el := range nums {
		cnt[el]++
	}

	type pair struct {
		num int
		freq int
	}

	pairs := make([]pair, 0, len(cnt))
	for num, freq := range cnt {
		pairs = append(pairs, pair{num, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	ans := make([]int, 0, k)

	for i := 0; i < k && i < len(pairs); i++ {
		ans[i] = pairs[i].num
	}

	return ans
}


func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(solution(nums, k))
}
