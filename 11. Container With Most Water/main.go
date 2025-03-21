package main

import "fmt"

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


func maxArea(height []int) int {
	var (
        l, r = 0, len(height) - 1
        maxArea = 0
    )

    for l < r {
        curArea := (r-l) * min(height[l], height[r])
        maxArea = max(maxArea, curArea)
        if height[r] < height[l] {
            r--
        } else {
            l++
        }
    }

    return maxArea
}

func main() {
    height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    height2 := []int{1, 1}

    fmt.Println(maxArea(height1))
    fmt.Println(maxArea(height2))
}