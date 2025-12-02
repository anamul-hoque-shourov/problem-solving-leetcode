package leetcode_682

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var score []int
	for idx := range operations {
		isInt, err := strconv.Atoi(operations[idx])
		if err != nil {
			if operations[idx] == "+" {
				score = append(score, score[(len(score)-1)]+score[(len(score)-2)])
				fmt.Println("+ sign", score)
			}
			if operations[idx] == "D" {
				score = append(score, (score[len(score)-1])*2)
				fmt.Println("D sign", score)
			}
			if operations[idx] == "C" {
				score = score[:len(score)-1]
				fmt.Println("C sign", score)
			}
		} else {
			score = append(score, isInt)
		}

	}
	var result int
	for idx := range score {
		result = result + score[idx]
	}
	fmt.Println(score)
	return result
}

func TestCode() {
	result := calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})
	fmt.Println(result)
}
