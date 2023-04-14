package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // empty list no insersection
		return nil
	}
	var lastA, lastB *ListNode // end of list
	l1, l2 := 1, 1             // length of list
	for lastA = headA; lastA.Next != nil; lastA = lastA.Next {
		l1++
	}
	for lastB = headB; lastB.Next != nil; lastB = lastB.Next {
		l2++
	}
	if lastA != lastB { // no common end
		return nil
	}
	if l1 < l2 { // which longer goes first
		for i := 0; i < l2-l1; i++ {
			headB = headB.Next
		}
	} else if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			headA = headA.Next
		}
	}
	for ; headA != headB; headA, headB = headA.Next, headB.Next {
	}
	return headA
}
