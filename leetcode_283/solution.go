package leetcode_283

import "fmt"

/*
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
*/

func moveZeroes(nums []int) []int {
	i := 0
	for _, num := range nums {
		if num != 0 {
			nums[i] = num
			i++
		}
	}

	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
	return nums
}

func TestCode() {
	x := []int{0, 5, 0, 20, 0, 10, 0, 14}
	result := moveZeroes(x)
	fmt.Println(result)
}
