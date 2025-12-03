package leetcode_15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var triplet [][]int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				triplet = append(triplet, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return triplet
}

func TestCode() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
}
