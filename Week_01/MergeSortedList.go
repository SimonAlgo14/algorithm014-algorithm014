package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	h := ListNode{} //head
	t := &h         //pointer

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}

	if l1 != nil {
		t.Next = l1
	} else {
		t.Next = l2
	}

	return h.Next
}
