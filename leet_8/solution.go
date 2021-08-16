package leet_8

import (
	"math"
)

func myAtoi(s string) int {
	input := []byte(s)

	resBase := int64(0)
	negPos := int64(1)
	numberCalStart := false

	for i := 0; i < len(input); i++ {
		if !numberCalStart {
			if isNumber(input[i]) || input[i] == '+' || input[i] == '-' {
				if isNumber(input[i]) {
					resBase = resBase*10 + (int64(input[i]) - int64('0'))
				} else if input[i] == '-' {
					negPos = -1
				}
				numberCalStart = true
			} else if input[i] == ' ' {
				continue
			} else {
				return 0
			}
		} else {
			if isNumber(input[i]) {
				resBase = resBase*10 + (int64(input[i]) - int64('0'))
				if resBase*negPos > math.MaxInt32 {
					return math.MaxInt32
				} else if resBase*negPos < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				return int(resBase * negPos)
			}
		}
	}

	return int(resBase * negPos)
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func MyAtoi(s string) int {
	return myAtoi(s)
}
