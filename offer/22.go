package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var size int
	for n := head; n != nil; n = n.Next {
		size++
	}
	for i := 0; i < size-k; i++ {
		head = head.Next
	}
	return head
}
