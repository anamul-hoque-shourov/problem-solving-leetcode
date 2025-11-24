package main

import (
	"fmt"
	"playground/validateSubsequence"
)

func main() {

	mainString := "constant"

	subString1 := "cntn"
	result1 := validateSubsequence.IsSubsequence(subString1, mainString) //true

	subString2 := "cntna"
	result2 := validateSubsequence.IsSubsequence(subString2, mainString) //false

	fmt.Println(result1, result2) // true, false
}
