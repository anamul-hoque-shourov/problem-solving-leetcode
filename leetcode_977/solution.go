package leetcode_977

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	var newArray []int
	for _, num := range nums {
		num = num * num
		newArray = append(newArray, num)
	}
	sort.Ints(newArray)
	return newArray
}

func TestCode() {
	x := []int{7, -1, -2, 3}
	result := sortedSquares(x)
	fmt.Println(result)
}
