package main

func findDisappearedNumbers(nums []int) []int {
	tmp := make([]bool, len(nums))
	for _, num := range nums {
		tmp[num-1] = true
	}
	ret := make([]int, 0, len(nums))
	for k, v := range tmp {
		if v {
			continue
		}
		ret = append(ret, k+1)
	}
	return ret
}

func findDisappearedNumbersV2(nums []int) []int {
	var trySwap func(i int)
	trySwap = func(i int) {
		j := nums[i] - 1
		if nums[j] == j+1 { // target is good
			return
		}
		nums[i], nums[j] = nums[j], nums[i] // swap
		if nums[i] == i+1 {                 // self is good
			return
		}
		trySwap(i)
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == i+1 { // 刚好符合
			continue
		}
		trySwap(i)
	}
	ret := make([]int, 0, len(nums))
	for i, num := range nums {
		if num == i+1 {
			continue
		}
		ret = append(ret, i+1)
	}
	return ret
}
