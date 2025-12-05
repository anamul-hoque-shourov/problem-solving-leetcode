package leetcode_125

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// convert to lowercase
	s = strings.ToLower(s)

	// remove spaces
	reWhitespace := regexp.MustCompile(`\s+`)
	s = reWhitespace.ReplaceAllString(s, "")

	// remove special characters
	reSpecial := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = reSpecial.ReplaceAllString(s, "")

	left := 0
	right := len(s) - 1
	for left < right && s[left] == s[right] {
		fmt.Println("loop")
		left++
		right--
	}
	if left >= right {
		return true
	} else {
		return false
	}
}

func cleanString(s string) string {
	// convert to lowercase
	s = strings.ToLower(s)

	// remove spaces
	reWhitespace := regexp.MustCompile(`\s+`)
	s = reWhitespace.ReplaceAllString(s, "")

	// remove special characters
	reSpecial := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = reSpecial.ReplaceAllString(s, "")

	return s
}

func TestCode() {
	cleanString := cleanString("A man, a plan, a canal: Panama")
	fmt.Println(cleanString)
	result := isPalindrome(cleanString)
	fmt.Println(result)
}
