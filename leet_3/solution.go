package leet3

func lengthOfLongestSubstring(s string) (maxLen int) {
	chars := []byte(s)
	if len(chars) == 0 {
		return
	}

	for startIndex, endIndex, curCharMap := 0, 0, make(map[byte]bool); endIndex < len(chars); endIndex++ {
		if !curCharMap[chars[endIndex]] {
			curCharMap[chars[endIndex]] = true
		} else {
			for ; chars[startIndex] != chars[endIndex]; startIndex++ {
				curCharMap[chars[startIndex]] = false
			}
			startIndex++
		}

		if newLen := endIndex - startIndex + 1; newLen > maxLen {
			maxLen = newLen
		}
	}

	return
}
