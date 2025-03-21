package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l < r {
		curSum := numbers[l] + numbers[r]
		if curSum < target {
			l++
		} else if curSum > target {
			r--
		} else {
			return []int{l+1, r+1}
		}
	}

	return []int{}
}

func main() {
	numbers1, target1 := []int{2, 7, 11, 15}, 9
	numbers2, target2 := []int{2, 3, 4}, 6
	numbers3, target3 := []int{-1, 0}, -1

	fmt.Println(twoSum(numbers1, target1))
	fmt.Println(twoSum(numbers2, target2))
	fmt.Println(twoSum(numbers3, target3))
}