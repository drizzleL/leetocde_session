package main

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var helper func(n *TreeNode, curr []int, val int)
	helper = func(n *TreeNode, curr []int, val int) {
		curr = append(curr, n.Val)
		val += n.Val
		if val == target && n.Left == nil && n.Right == nil {
			clone := make([]int, len(curr))
			copy(clone, curr)
			ret = append(ret, clone)
		}
		if n.Left != nil {
			helper(n.Left, curr, val)
		}
		if n.Right != nil {
			helper(n.Right, curr, val)
		}
		return
	}
	helper(root, nil, 0)
	return ret
}
