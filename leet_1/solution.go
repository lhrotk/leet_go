package leet1

func twoSum(nums []int, target int) (res []int) {
	resStore := make(map[int]int)
	for idx, num := range nums {
		if pos, ok := resStore[num]; ok {
			return []int{pos, idx}
		}
		resStore[target-num] = idx
	}
	return nil
}
