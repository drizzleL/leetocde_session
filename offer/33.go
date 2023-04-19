package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	rootVal := postorder[len(postorder)-1]
	// 是否可以分成<x,>x
	bigger := -1
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > rootVal { // 找到第一个比rootVal大的
			bigger = i
			break
		}
	}
	if bigger == -1 { // all smaller than root
		return verifyPostorder(postorder[:len(postorder)-1])
	}
	// check last all bigger than root
	for i := bigger; i < len(postorder)-1; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}
	return verifyPostorder(postorder[:bigger]) && verifyPostorder(postorder[bigger:len(postorder)-1])
}
