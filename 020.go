package main

func isValid(s string) bool {
	var stack []rune
	dict := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		default:
			if len(stack) == 0 {
				return false
			}
			if dict[c] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
