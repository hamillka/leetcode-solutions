package main

func maxSubArray(nums []int) int {
	mxx := nums[0]
	cur := 0

	for _, v := range nums {
		cur += v
		mxx = max(cur, mxx)
		if cur < 0 {
			cur = 0
		}
	}
	return mxx
}
