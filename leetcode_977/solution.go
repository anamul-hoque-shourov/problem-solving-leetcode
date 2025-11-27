package leetcode_977

import (
	"fmt"
	"math"
)

/*
func sortedSquares(nums []int) []int {
	var newArray []int
	for _, num := range nums {
		num = num * num
		newArray = append(newArray, num)
	}
	sort.Ints(newArray)
	return newArray
}
*/

func sortedSquares(nums []int) []int {
	newArray := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	for i := len(newArray) - 1; i >= 0; i-- {
		if int(math.Abs(float64(nums[left]))) > int(math.Abs(float64(nums[right]))) {
			newArray[i] = nums[left] * nums[left]
			left++
		} else {
			newArray[i] = nums[right] * nums[right]
			right--
		}
	}
	return newArray
}

func TestCode() {
	x := []int{-7, -3, 2, 3, 11}
	result := sortedSquares(x)
	fmt.Println(result)
}
