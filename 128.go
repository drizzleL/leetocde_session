package main

func longestConsecutive(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		if _, used := dict[num]; used {
			continue
		}
		group := num                  // default group
		if v, ok := dict[num-1]; ok { // use smaller group
			group = v
		}
		if _, ok := dict[num+1]; ok { // found bigger group, change
			dict[num+1] = group
		}
		dict[num] = group
	}
	var find func(x int) int
	savedMapping := map[int]int{}
	find = func(x int) (m int) {
		if v, ok := savedMapping[x]; ok {
			return v
		}
		defer func() {
			savedMapping[x] = m
		}()
		if dict[x] == x {
			return x
		}
		return find(dict[x])
	}
	var ret int
	tmp := map[int]int{}
	for k := range dict {
		finalG := find(k)
		tmp[finalG]++
		if tmp[finalG] > ret {
			ret = tmp[finalG]
		}
	}
	return ret
}
