package leet15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		} else if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			curSum := nums[i] + nums[l] + nums[r]
			if curSum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				curL, curR := nums[l], nums[r]
				l++
				r--
				for l < r && nums[l] == curL {
					l++
				}
				for l < r && nums[r] == curR {
					r--
				}
			} else if curSum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
