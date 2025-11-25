package leetcode_509

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestCode() {
	result := fib(7)
	fmt.Println(result)
}
