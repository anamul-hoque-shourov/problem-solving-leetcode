package leetcode_2824

import (
	"fmt"
)

func countPairs(nums []int, target int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				fmt.Println(i, j)
				count++
			}
		}
	}
	return count
}

func TestCode() {
	result := countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2)
	fmt.Println(result)
}
