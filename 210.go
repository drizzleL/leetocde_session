package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var ret []int
	out := make([]int, numCourses)
	relas := map[int][]int{}
	for _, pair := range prerequisites {
		out[pair[0]]++
		relas[pair[1]] = append(relas[pair[1]], pair[0])
	}
	var zeros []int
	for i, v := range out {
		if v == 0 {
			zeros = append(zeros, i)
		}
	}
	for len(zeros) != 0 {
		v := zeros[len(zeros)-1]
		ret = append(ret, v)
		zeros = zeros[:len(zeros)-1]
		for _, k := range relas[v] {
			out[k]--
			if out[k] == 0 {
				zeros = append(zeros, k)
			}
		}
	}
	for _, v := range out {
		if v != 0 {
			return nil
		}
	}
	return ret
}
