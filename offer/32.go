package main

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ret []int
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			ret = append(ret, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	return ret
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ret [][]int
	for len(nodes) != 0 {
		var next []*TreeNode
		var line []int
		for _, n := range nodes {
			line = append(line, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
		ret = append(ret, line)
	}
	return ret
}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ret [][]int
	var s bool
	for len(nodes) != 0 {
		var next []*TreeNode
		var line []int
		for _, n := range nodes {
			line = append(line, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if s {
			for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
				line[i], line[j] = line[j], line[i]
			}
		}
		nodes = next
		ret = append(ret, line)
		s = !s
	}
	return ret
}
