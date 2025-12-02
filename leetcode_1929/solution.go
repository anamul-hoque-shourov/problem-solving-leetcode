package leetcode_1929

import "fmt"

func getConcatenation(nums []int) []int {
	newArray := make([]int, len(nums)*2)
	j := 0
	for i := range newArray {
		if j == len(nums) {
			j = 0
		}
		newArray[i] = nums[j]
		j++
	}

	return newArray
}

func TestCode() {
	result := getConcatenation([]int{1})
	fmt.Println(result)
}
