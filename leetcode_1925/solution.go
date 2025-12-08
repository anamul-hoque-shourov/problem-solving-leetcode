package leetcode_1925

import (
	"fmt"
	"math"
)

/*
	condition: 0 <= a, b, c <= n
	return: number of triplets possible
*/

func countTriples(n int) int {
	var array [][]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			square := i*i + j*j
			sqrtFloat := math.Sqrt(float64(square))
			sqrtInt := int(sqrtFloat)
			if square == sqrtInt*sqrtInt && sqrtInt <= n {
				array = append(array, []int{i, j, sqrtInt})
			}
		}
	}
	return len(array)
}

func TestCode() {
	result := countTriples(150)
	fmt.Println(result)
}
