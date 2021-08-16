package leet_16

import "fmt"

func maxArea(height []int) int {
	maxAreaRes := 0
	for i, j := 0, len(height)-1; i < j; {
		newArea := calcArea(i, j, height)
		if newArea > maxAreaRes {
			fmt.Printf("maxArea candidate: %v, index: %v, %v\n", newArea, i, j)
			maxAreaRes = newArea
		}
		if height[i] > height[j] {
			j--
		} else if height[i] < height[j] {
			i++
		} else {
			i++
			j--
		}
	}
	return maxAreaRes
}

func calcArea(x1, x2 int, height []int) int {
	h := height[x1]
	if height[x2] < h {
		h = height[x2]
	}
	return h * (x2 - x1)
}
