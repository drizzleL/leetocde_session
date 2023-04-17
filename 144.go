package main

func preorderTraversal(root *TreeNode) []int {
	var vals []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			vals = append(vals, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return vals
}
