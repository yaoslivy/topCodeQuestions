package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//设置快慢指针，快指针先走n步
	fast := head
	for fast != nil && n != 0 {
		fast = fast.Next
		n--
	}
	// 1->2->3(fast)->4->5
	if fast == nil && n != 0 {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	slow := dummyHead
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummyHead.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 设置快慢指针，使得两指针间隔n个节点，当快指针走到nil时，慢指针的下一个节点就是需要删除的节点
	if head == nil {
		return nil
	}
	//n范围不合法
	if n <= 0 || n > getLength(head) {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	var fast, slow *ListNode
	fast = dummyHead.Next
	for n != 0 && fast != nil {
		fast = fast.Next
		n--
	}
	slow = dummyHead
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummyHead.Next
}

func getLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	res := 1
	for head != nil {
		head = head.Next
		res++
	}
	return res
}
