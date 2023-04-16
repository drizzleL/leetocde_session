package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	n := head
	helper := func(l *ListNode) *ListNode {
		n.Next = l
		n = n.Next
		return l.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1 = helper(l1)
		} else {
			l2 = helper(l2)
		}
	}
	for l1 != nil {
		l1 = helper(l1)
	}
	for l2 != nil {
		l2 = helper(l2)
	}
	return head.Next
}
