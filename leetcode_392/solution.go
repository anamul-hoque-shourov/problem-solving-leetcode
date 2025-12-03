package leetcode_392

import "fmt"

func isSubsequence(s string, t string) bool {
	j := 0
	if len(s) < 2 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}

func TestCode() {
	mainString := "constant"

	subString1 := "cntn"
	result1 := isSubsequence(subString1, mainString) //true

	subString2 := "cntna"
	result2 := isSubsequence(subString2, mainString) //false

	fmt.Println(result1, result2) // true, false
}
