package leetcode_704

import "fmt"

func search(nums []int, target int) int {
	var left, right, middle int
	left = 0
	right = len(nums) - 1
	for left <= right {
		middle = (left + right) / 2
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func TestCode() {
	result := search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)
	fmt.Println(result)
}
