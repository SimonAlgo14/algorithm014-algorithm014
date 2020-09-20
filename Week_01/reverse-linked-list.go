package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Input:  1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

 */

func reverseList(head *ListNode) *ListNode {
	var newList  ListNode

	for head != nil {
		nextNode := head.Next // Keep

		head.Next = newList.Next  // Move head to new List head
		newList.Next = head       //

		head = nextNode // Next
	}

	return newList.Next
}

func reverse2(head *ListNode, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}

	nextNode := head.Next
	head.Next = newHead

	return reverse2(nextNode, head)
}
