package leetcode_27

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func TestCode() {
	y := []int{0, 1, 2, 6, 6, 7, 7, 8, 10, 11}
	result := removeElement(y, 7)
	fmt.Println(result)
}
