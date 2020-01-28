package q001

func twoSum(nums []int, target int) []int {
	indexCache := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := indexCache[num]; ok {
			return []int{i, idx}
		}
		indexCache[target-num] = i
	}
	return []int{}
}
