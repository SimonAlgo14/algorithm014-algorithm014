package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return  nil
	}

	return _mergeKLists(lists, 0 , len(lists) -1 )
}

func _mergeKLists(lists []*ListNode, low int, high int) *ListNode {
	if low >= high {
		return lists[low]
	}

	mid   := low + (high-low) / 2
	left  := _mergeKLists(lists, low,       mid)
	right := _mergeKLists(lists,  mid + 1,  high)

	return _merge(left, right)
}

func _merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = _merge(l1.Next, l2)
		return l1
	}

	l2.Next = _merge(l1, l2.Next)
	return l2
}
