package leetcode_1523

import "fmt"

func countOdds(low int, high int) int {
	var count int
	if low == high && high%2 == 1 {
		count = 1
	} else {
		count = 0
	}
	if low%2 == 0 && high%2 == 0 {
		count = (high - low) / 2
	} else {
		count = (high-low)/2 + 1
	}
	return count
}

func TestCode() {
	result := countOdds(9, 11)
	fmt.Println(result)
}
