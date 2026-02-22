package main

import "slices"

func maxProduct(nums []int) int {
	minP, maxP, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		cand := []int{
			nums[i],
			nums[i] * minP,
			nums[i] * maxP,
		}
		minP = slices.Min(cand)
		maxP = slices.Max(cand)

		res = max(res, maxP)
	}

	return res
}
