package string

import "strings"

func NoRepeatLongestString(s string) int {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 && (i+1) > end {
			end = i + 1
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
