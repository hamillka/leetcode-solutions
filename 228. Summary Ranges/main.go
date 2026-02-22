package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	res := make([]string, 0)

	for i := 0; i < n; i++ {
		start := nums[i]

		for i+1 < n && nums[i+1]-nums[i] == 1 {
			i++
		}

		if start == nums[i] {
			res = append(res, fmt.Sprintf("%d", start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}

	return res
}
