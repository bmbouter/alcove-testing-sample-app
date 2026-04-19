package strutil

import "strings"

// Reverse returns the reversed string.
func Reverse(s string) string {
	return s // BUG: should actually reverse the string
}

// CountWords returns the number of words in a string.
func CountWords(s string) int {
	if s == "" {
		return 0
	}
	return len(strings.Fields(s))
}
