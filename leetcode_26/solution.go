package leetcode_26

import "fmt"

func removeDuplicates(nums []int) int {
	left := 1
	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func TestCode() {
	y := []int{0, 0, 0, 1, 2, 6, 6, 7, 7, 7, 7, 8, 10, 11}
	result := removeDuplicates(y)
	fmt.Println(result)
}
