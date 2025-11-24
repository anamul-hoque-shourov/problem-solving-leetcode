package validateSubsequence

import "fmt"

func IsSubsequence(s string, t string) bool {
	j := 0
	if len(s) < 1 {
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
	result1 := IsSubsequence(subString1, mainString) //true

	subString2 := "cntna"
	result2 := IsSubsequence(subString2, mainString) //false

	fmt.Println(result1, result2) // true, false
}
