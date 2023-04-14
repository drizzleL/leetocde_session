package main

func firstUniqChar(s string) byte {
	b := make([]int, 26)
	for i, c := range s {
		b[c-'a'] = i // last seen
	}
	for i, c := range s {
		if b[c-'a'] == -1 {
			continue
		}
		if b[c-'a'] == i {
			return byte(c)
		} else {
			b[c-'a'] = -1
		}
	}
	return ' '
}
