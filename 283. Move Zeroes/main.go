func moveZeroes(nums []int)  {
	zeroPtr := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[zeroPtr] = nums[zeroPtr], nums[i]
			zeroPtr++
		}
	}
}
