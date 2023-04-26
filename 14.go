package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	b := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || c != strs[j][i] { // exceed length
				return string(b)
			}
		}
		b = append(b, c)
	}
	return string(b)
}
