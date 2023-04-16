package main

func deleteNode(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	n := head
	for n.Next != nil {
		if n.Next.Val == val {
			n.Next = n.Next.Next
		}
		n = n.Next
	}
	return head
}
