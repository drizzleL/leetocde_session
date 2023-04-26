package main

import "strconv"

func countAndSay(n int) string {
	ret := "1"
	for i := 1; i < n; i++ {
		b := []byte{}
		for j := 0; j < len(ret); j++ {
			c := ret[j]
			cnt := 1
			for j+cnt < len(ret) && ret[j+cnt] == c {
				cnt++
			}
			b = append(b, []byte(strconv.Itoa(cnt))...)
			b = append(b, c)
			j += cnt - 1
		}
		ret = string(b)
	}
	return ret
}
