package leetcode_3577

import "fmt"

func factorial(n int) int {
	const MOD = 1_000_000_007
	res := int64(1)
	for i := 2; i <= n; i++ {
		res = (res * int64(i)) % MOD
	}
	return int(res)
}

func countPermutations(complexity []int) int {
	if len(complexity) <= 1 {
		return 1
	}
	for i := 1; i < len(complexity); i++ {
		if complexity[0] >= complexity[i] {
			return 0
		}
	}
	return factorial(len(complexity) - 1)
}

func TestCode() {
	result := countPermutations([]int{1, 4, 3})
	fmt.Println(result)
}
