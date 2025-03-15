package main

import "fmt"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    res[0] = 1

    for i := 1; i < n; i++ {
        res[i] = res[i-1] * nums[i-1]
    }

    right := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
    }

    return res
}

func main() {
    nums1 := []int{1, 2, 3, 4}
    nums2 := []int{-1, 1, 0, -3, 3}

    fmt.Println(productExceptSelf(nums1))
    fmt.Println(productExceptSelf(nums2))
}
