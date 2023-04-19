package main

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strings.Compare(strs[i]+strs[j], strs[j]+strs[i]) < 0
	})
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}
