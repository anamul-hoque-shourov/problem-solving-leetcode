package leetcode_283

import "fmt"

func moveZeroes(nums []int) []int {
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	return nums
}

func TestCode() {
	x := []int{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14}
	result := moveZeroes(x)
	fmt.Println(result)
}
