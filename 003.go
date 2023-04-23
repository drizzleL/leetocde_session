package main

func lengthOfLongestSubstring(s string) int {
	dict := make([]int, 256)
	ret, currStart := 0, 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		old := dict[v]
		if old != 0 && (old-1) >= currStart {
			currStart = old
		}
		curr := i - currStart + 1
		if curr > ret {
			ret = curr
		}
		dict[v] = i + 1
	}
	return ret
}
