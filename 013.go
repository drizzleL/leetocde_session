package main

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ret int
	for i := 0; i < len(s); i++ {
		v := dict[s[i]]
		if i+1 < len(s) && dict[s[i+1]] > v {
			ret -= v
		} else {
			ret += v
		}
	}
	return ret
}
