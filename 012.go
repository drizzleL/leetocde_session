package main

import "bytes"

func intToRoman(num int) string {
	dict := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	base := 1
	b := []byte{}
	for num != 0 {
		v := num % 10
		switch v {
		case 1, 2, 3:
			b = append(b, bytes.Repeat([]byte{dict[base*1]}, v)...)
		case 5:
			b = append(b, dict[base*5])
		case 4:
			b = append(b, dict[base*5], dict[base*1])
		case 6, 7, 8:
			b = append(b, bytes.Repeat([]byte{dict[base*1]}, v-5)...)
			b = append(b, dict[base*5])
		case 9:
			b = append(b, dict[base*10], dict[base*1])
		}
		num /= 10
		base *= 10
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
