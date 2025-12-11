package leetcode_3531

import (
	"fmt"
	"sort"
)

func countCoveredBuildings(buildings [][]int) int {
	xMap := make(map[int][]int, 0)
	yMap := make(map[int][]int, 0)
	for _, building := range buildings {
		x := building[0]
		y := building[1]
		xMap[x] = append(xMap[x], y)
		yMap[y] = append(yMap[y], x)
	}

	for x := range xMap {
		sort.Ints(xMap[x])
	}
	for y := range yMap {
		sort.Ints(yMap[y])
	}

	var count int
	for _, building := range buildings {
		x := building[0]
		y := building[1]

		xList := xMap[x]
		yList := yMap[y]

		yRow := sort.SearchInts(xList, y)
		xColumn := sort.SearchInts(yList, x)

		if yRow > 0 && yRow < len(xList)-1 && xColumn > 0 && xColumn < len(yList)-1 {
			count++
		}
	}
	return count
}

func TestCode() {
	result := countCoveredBuildings([][]int{{2, 3}, {2, 1}, {2, 5}, {1, 3}, {3, 3}})
	fmt.Println(result)
}
