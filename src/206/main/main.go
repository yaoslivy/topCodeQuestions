package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代方式
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	p := head
	for p != nil {
		nex := p.Next
		p.Next = pre
		pre = p
		p = nex
	}
	return pre
}

/*
 可以使用递归解法的条件：
 1. 大问题拆成两个子问题
 2. 子问题求解方式和大问题一样
 3. 存在最小子问题


*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
