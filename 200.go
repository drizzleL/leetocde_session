package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '#'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	var ret int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ret++
				dfs(i, j)
			}
		}
	}
	return ret
}
func numIslands2(grid [][]byte) int {
	getId := func(i, j int) int {
		return i*len(grid[0]) + j
	}
	uf := NewUFWithGrid(grid)
	join := func(i, j int, oldId int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] != '1' {
			return
		}
		uf.union(oldId, getId(i, j))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			id := getId(i, j)
			join(i-1, j, id)
			join(i+1, j, id)
			join(i, j-1, id)
			join(i, j+1, id)
		}
	}
	return uf.count()
}

type UF struct {
	parent []int
	size   []int
	cnt    int
}

func NewUFWithGrid(grid [][]byte) *UF {
	n := len(grid) * len(grid[0])
	parent := make([]int, n)
	size := make([]int, n)
	var cnt int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			id := i*len(grid[0]) + j
			size[id] = 1
			parent[id] = id
			cnt++
		}
	}
	return &UF{parent: parent, cnt: cnt, size: size}
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UF{parent: parent, cnt: n, size: size}
}
func (u *UF) union(a, b int) {
	rootA, rootB := u.find(a), u.find(b)
	if rootA == rootB {
		return
	}
	// smaller merge into bigger
	if u.size[rootA] > u.size[rootB] {
		u.parent[rootB] = rootA
		u.size[rootA] += u.size[rootB]
	} else {
		u.parent[rootA] = rootB
		u.size[rootB] += u.size[rootA]
	}
	u.cnt--
}

func (u *UF) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]] // merge when not found
		x = u.parent[x]
	}
	return x
}

func (u *UF) count() int {
	return u.cnt
}
