package leetcode_1470

import "fmt"

func shuffle(nums []int, n int) []int {
	newArray := make([]int, 2*n)
	for i := 0; i < n; i++ {
		newArray[2*i] = nums[i]
		newArray[2*i+1] = nums[n+i]
	}
	return newArray
}

func TestCode() {
	result := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	fmt.Println(result)
}
