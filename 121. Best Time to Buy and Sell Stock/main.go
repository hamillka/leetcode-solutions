func maxProfit(prices []int) int {
	s := prices[0]
	maxProf := 0

	for _, v := range prices {
		if v < s {
			s = v
		}
		if v-s > maxProf {
			maxProf = v - s
		}
	}

	return maxProf
}