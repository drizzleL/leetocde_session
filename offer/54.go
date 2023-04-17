package main

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	var vals []int
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, root.Val)
			root = root.Right
		}
	}
	return vals[k-1]
}
