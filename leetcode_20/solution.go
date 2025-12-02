package leetcode_20

import "fmt"

func isValid(s string) bool {
	var stack []rune
	if len(s)%2 == 1 {
		return false
	}
	for _, bracket := range s {
		if bracket == '(' || bracket == '{' || bracket == '[' {
			stack = append(stack, bracket)
		}
		if bracket == ')' || bracket == '}' || bracket == ']' {
			if (stack[len(stack)-1] == '(' && bracket != ')') ||
				(stack[len(stack)-1] == '{' && bracket != '}') ||
				(stack[len(stack)-1] == '[' && bracket != ']') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func TestCode() {
	result := isValid("[{()]")
	fmt.Println(result)
}
