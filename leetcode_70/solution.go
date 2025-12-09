package leetcode_70

import "fmt"

// TLE issue
/*
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
*/

func climbStairs(n int) int {
	hashmap := make(map[int]int, 0)
	var recurring func(int) int
	recurring = func(num int) int {
		if num == 0 || num == 1 {
			return 1
		}
		if val, ok := hashmap[num]; ok {
			return val
		}
		hashmap[num] = recurring(num-1) + recurring(num-2)
		return hashmap[num]
	}
	return recurring(n)
}

func TestCode() {
	result := climbStairs(44)
	fmt.Println(result)
}
