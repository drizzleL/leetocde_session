package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				for j < k {
					j++
					if nums[j] != nums[j-1] {
						break
					}
				}
				for j < k {
					k--
					if nums[k] != nums[k+1] {
						break
					}
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ret
}
