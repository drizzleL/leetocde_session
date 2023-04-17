package main

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	// 从后面找e
	idx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'e' || s[i] == 'E' {
			idx = i
			break
		}
	}
	if idx == -1 {
		return isInt(s) || isDecimal(s)
	}
	return isInt(s[idx+1:]) && (isInt(s[:idx]) || isDecimal(s[:idx]))
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '+', '-':
		return isPureInt(s[1:])
	}
	return isPureInt(s)
}

func isPureInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true

}
func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '+', '-':
		return isPureDeci(s[1:])
	}
	return isPureDeci(s)
}

func isPureDeci(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '.' {
		return isPureInt(s[1:])
	}
	idx := -1
	for i, c := range s {
		if c == '.' {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}
	if idx == len(s)-1 {
		return isPureInt(s[:idx])
	}
	return isPureInt(s[:idx]) && isPureInt(s[idx+1:])
}
