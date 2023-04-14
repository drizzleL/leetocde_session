package main

func replaceSpace(s string) string {
	buf := make([]byte, 0, len(s)+20)
	for _, c := range s {
		if c != ' ' {
			buf = append(buf, byte(c))
		} else {
			buf = append(buf, []byte("%20")...)
		}
	}
	return string(buf)
}
