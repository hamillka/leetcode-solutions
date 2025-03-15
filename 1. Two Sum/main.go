package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        if idx, ok := m[target - v]; ok {
            return []int{i, idx}
        }
        m[v] = i
    }

    return []int{}
}

func main() {
    nums1 := []int{2, 7, 11, 15}
    target1 := 9

    nums2 := []int{3, 2, 4}
    target2 := 6

    nums3 := []int{3, 3}
    target3 := 6

    fmt.Println(twoSum(nums1, target1))
    fmt.Println(twoSum(nums2, target2))
    fmt.Println(twoSum(nums3, target3))
}