func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

    for ptr := 0; ptr < len(nums) - 2; ptr++ {
		if ptr > 0 && nums[ptr] == nums[ptr-1] {
			continue
		}

		lp := ptr + 1
		rp := len(nums) - 1

		for lp < rp {
			tripleSum := nums[ptr] + nums[lp] + nums[rp]
			if tripleSum > 0 {
				rp--
			} else if tripleSum < 0 {
				lp++
			} else {
				res = append(res, []int{nums[ptr], nums[lp], nums[rp]})
				lp++
				rp--
                for lp < rp && nums[lp] == nums[lp-1] {
					lp++
				}
				for lp < rp && nums[rp] == nums[rp+1] {
					rp--
				}
			}
		}
	}

	return res
}
