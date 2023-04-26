package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := make([]byte, len(s))
	perRound := numRows*2 - 2
	var j int
	for i := 0; i < len(s); i = i + perRound {
		b[j] = s[i]
		j++
	}
	for row := 1; row < numRows-1; row++ {
		for i := row; i < len(s); i += perRound {
			b[j] = s[i]
			j++
			next := i + perRound - 2*row
			if next < len(s) {
				b[j] = s[next]
				j++
			}
		}
	}
	for i := numRows - 1; i < len(s); i += perRound {
		b[j] = s[i]
		j++
	}
	return string(b)
}
