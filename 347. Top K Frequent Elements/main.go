func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, el := range nums {
		cnt[el]++
	}

	buckets := make([][]int, len(nums)+1)
	for el, freq := range cnt {
		buckets[freq] = append(buckets[freq], el)
	}

	ans := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, el := range buckets[i] {
			if k > 0 {
				ans = append(ans, el)
				k--
			}
		}
	}

	return ans
}