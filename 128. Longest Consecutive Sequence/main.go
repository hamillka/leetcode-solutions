package main

import "fmt"

func longestConsecutive(nums []int) int {
    m := make(map[int]struct{})
    for _, v := range nums {
    	m[v] = struct{}{}
    }

    maxLen := 0

    for key := range m {
    	if _, ok := m[key - 1]; !ok {
    		sequenceLen := 1
    		for {
    			if _, ok = m[key + sequenceLen]; ok {
    				sequenceLen++
    				continue
    			}
    			maxLen = max(sequenceLen, maxLen)
    			break
    		}
    	}
    }

    return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums3 := []int{1, 0, 1, 2}

	fmt.Println(longestConsecutive(nums1))
	fmt.Println(longestConsecutive(nums2))
	fmt.Println(longestConsecutive(nums3))
}