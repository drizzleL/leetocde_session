package main

import (
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Fields(s)
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strings.Join(strs, " ")
}

func reverseLeftWords(s string, n int) string {
	b := make([]byte, len(s))
	for i := 0; i < n; i++ {
		b[len(s)-n+i] = s[i]
	}
	for i := n; i < len(s); i++ {
		b[i-n] = s[i]
	}
	return string(b)
}
