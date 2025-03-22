func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	m := len(matrix)

	l, r := 0, m*n-1

	for l <= r {
		mid := (l + r) / 2
		row := mid / n
		col := mid - row*n
		if matrix[row][col] < target {
			l = mid + 1
		} else if matrix[row][col] > target {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}
