package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	n := len(intervals)
	prev := 0
	nonOverlap := 1

	for i := 1; i < n; i++ {
		if intervals[i][0] >= intervals[prev][1] {
			prev = i
			nonOverlap++
		}
	}

	return n - nonOverlap
}
