package leet_275

func hIndex(citations []int) int {
	h := 0
	if len(citations) == 0 {
		return h
	}

	k := len(citations)
	for i := 0; i < k; i++ {
		h = k - i
		if i > 0 && citations[i] >= k-i && citations[i-1] <= k-i {
			return h
		} else if i == 0 && citations[i] >= k-i {
			return h
		}
	}
	return 0
}

func HIndex(citations []int) int {
	return hIndex(citations)
}
