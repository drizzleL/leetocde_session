package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func copyRandomList(head *Node) *Node {
	dict := map[*Node]*Node{}
	newPre := &Node{}
	for n, m := head, newPre; n != nil; n, m = n.Next, m.Next {
		m.Next = &Node{
			Val: n.Val,
		}
		dict[n] = m.Next
	}
	for n := head; n != nil; n = n.Next {
		if n.Random == nil {
			continue
		}
		dict[n].Random = dict[n.Random]
	}
	return newPre.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func treeToDoublyList(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}
	var helper func(root *ListNode) (*ListNode, *ListNode)
	helper = func(n *ListNode) (*ListNode, *ListNode) {
		if n.Prev == nil && n.Next == nil {
			return n, n
		}
		lHead, lTail := helper(root.Prev)
		rHead, rTail := helper(root.Next)
		lTail.Next = root
		root.Prev = lTail
		root.Next = rHead
		rHead.Prev = root
		return lHead, rTail
	}
	head, tail := helper(root)
	head.Prev = tail
	tail.Next = head
	return head
}
