package main

import "log"

func longestPalindrome(s string) string {
	b := make([]byte, len(s)*2+1)
	b[0] = '#'
	for i := 0; i < len(s); i++ {
		b[i*2] = s[i]
		b[i*2+1] = '#'
	}
	var ret string
	for i := 0; i < len(b); i++ {
		for l := 0; i-l >= 0 && i+l < len(s); l++ {
			if b[i-l] != b[i+l] {
				break
			}
			if len(b[i-l:i+l+1]) > len(ret) {
				ret = string(b[i-l : i+l+1])
			}
		}
	}
	log.Println(ret)
	return ret
}
