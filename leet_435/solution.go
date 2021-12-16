package main

import (
	"fmt"
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	inters := Spaces(intervals)
	sort.Sort(inters)

	selectNum, maxRigthNumber := 0, math.MinInt32
	for _, s := range inters {
		if s[0] >= maxRigthNumber {
			selectNum++
			maxRigthNumber = s[1]
		}
	}

	return inters.Len() - selectNum
}

type Spaces [][]int

func (s Spaces) Len() int {
	return len(s)
}

func (s Spaces) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s Spaces) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {3, 4}, {2, 3}, {1, 2}}))
}
