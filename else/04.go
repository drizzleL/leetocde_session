package main

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	if start == target {
		return true
	}
	v := map[int][]int{}
	seen := map[int]bool{}
	for _, pair := range graph {
		v[pair[0]] = append(v[pair[0]], pair[1])
	}
	nodes := []int{start}
	seen[start] = true
	for len(nodes) != 0 {
		n := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		for _, next := range v[n] {
			if next == target {
				return true
			}
			if seen[next] {
				continue
			}
			nodes = append(nodes, next)
			seen[next] = true
		}
	}
	return false
}
