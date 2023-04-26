package main

func longestPalindrome(s string) string {
	b := make([]byte, len(s)*2+1)
	b[0] = '#'
	for i := 0; i < len(s); i++ {
		b[i*2+1] = s[i]
		b[i*2+2] = '#'
	}
	var ret string
	for i := 0; i < len(b); i++ {
		for l := 0; i-l >= 0 && i+l < len(b); l++ {
			if b[i-l] != b[i+l] {
				break
			}
			if len(b[i-l:i+l+1]) > len(ret) {
				ret = string(b[i-l : i+l+1])
			}
		}
	}
	var bb []byte
	for _, c := range ret {
		if c == '#' {
			continue
		}
		bb = append(bb, byte(c))
	}
	return string(bb)
}

func longestPalindrome2(s string) string {
	b := make([]byte, len(s)*2+1)
	b[0] = '#'
	for i := 0; i < len(s); i++ {
		b[i*2+1] = s[i]
		b[i*2+2] = '#'
	}
	var maxRight, center int
	var start, maxSize int
	dp := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		if i < maxRight {
			mirror := center*2 - i
			dp[i] = min(dp[mirror], maxRight-i)

		}
		for l, r := i-dp[i]-1, i+dp[i]+1; l >= 0 && r < len(b) && b[l] == b[r]; l, r = l-1, r+1 {
			dp[i] = dp[i] + 1
		}
		if dp[i]+i > maxRight {
			maxRight = dp[i] + i
			center = i
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
			start = (i - maxSize) / 2
		}
	}
	return s[start : start+maxSize]
}
