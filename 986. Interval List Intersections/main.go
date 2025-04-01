func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    if len(firstList) == 0 || len(secondList) == 0 {
        return [][]int{}
    }

    res := [][]int{}
    var i, j int

    for i < len(firstList) && j < len(secondList) {
        start, end := max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])

        if start <= end {
            res = append(res, []int{start, end})
        }

        if firstList[i][1] < secondList[j][1] {
            i++
        } else {
            j++
        }
    }
    return res
}