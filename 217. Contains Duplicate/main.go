package main

import "fmt"

func containsDuplicate(nums []int) bool {
    m := make(map[int]int, len(nums))

    for _, el := range nums {
        if _, ok := m[el]; ok {
            return true
        } else {
            m[el] = 1
        }
    }
    return false
}

func main() {
    nums1 := []int{1,2,3,1}
    nums2 := []int{1,2,3,4}

    fmt.Println(containsDuplicate(nums1))
    fmt.Println(containsDuplicate(nums2))
}