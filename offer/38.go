package main

import "sort"

func permutation3(s string) []string {
	vis := make([]bool, len(s))
	dict := make(map[string]bool)
	var dfs func(str string)
	dfs = func(currstr string) {
		if len(currstr) == len(s) {
			dict[currstr] = true
			return
		}
		for j := 0; j < len(s); j++ {
			if !vis[j] {
				vis[j] = true
				dfs(currstr + string(s[j]))
				vis[j] = false
			}
		}
	}
	dfs("")
	var ret []string
	for v := range dict {
		ret = append(ret, v)
	}
	return ret
}

func permutation(s string) []string {
	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	vis := make([]bool, len(sb))
	var dfs func(b []byte)
	var ret []string
	dfs = func(curr []byte) {
		if len(curr) == len(sb) {
			ret = append(ret, string(curr))
			return
		}
		for j := 0; j < len(sb); j++ {
			if j > 0 && !vis[j-1] && sb[j] == sb[j-1] {
				continue
			}
			if !vis[j] {
				vis[j] = true
				dfs(append(curr, sb[j]))
				vis[j] = false
			}
		}
	}
	dfs(nil)
	return ret
}
