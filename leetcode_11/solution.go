package leetcode_11

import "fmt"

func maxArea(height []int) int {
	var max int
	left := 0
	right := len(height) - 1
	for left < right {
		base := (right - left)
		area1 := height[left] * base
		area2 := height[right] * base
		if height[left] < height[right] {
			if area1 > max {
				max = area1
			}
			left++
		} else {
			if area2 > max {
				max = area2
			}
			right--
		}
	}
	return max
}

func TestCode() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(result)
}
