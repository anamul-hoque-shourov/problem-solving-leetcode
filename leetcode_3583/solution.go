package leetcode_3583

import "fmt"

func specialTriplets(nums []int) int {
	const mod = 1_000_000_007
	var count int
	leftMap := make(map[int]int, 0)
	rightMap := make(map[int]int, 0)
	for _, num := range nums {
		rightMap[num]++
	}
	for j := 0; j < len(nums); j++ {
		rightMap[nums[j]]--
		if j > 0 && j < len(nums)-1 {
			leftValue := leftMap[nums[j]*2]
			rightValue := rightMap[nums[j]*2]
			count = (count + leftValue*rightValue) % mod
		}
		leftMap[nums[j]]++
	}
	return count
}

func TestCode() {
	result := specialTriplets([]int{0, 1, 0, 0})
	fmt.Println(result)
}
