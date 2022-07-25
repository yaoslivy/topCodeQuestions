package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//使用虚拟头节点
func removeElements(head *ListNode, val int) *ListNode {
	//创建一个虚拟头节点，然后让虚拟头节点指向head
	// head = [1,2,6,3,4,5,6], val = 6
	// begin->1
	begin := &ListNode{}
	begin.Next = head
	cur := begin
	for cur.Next != nil {
		if cur.Next.Val == val { //删除cur.Next节点
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return begin.Next
}

//在原链表上操作
func removeElements2(head *ListNode, val int) *ListNode {
	// head = [1,2,6,3,4,5,6], val = 6
	for head != nil && head.Val == val {
		head = head.Next
	}
	//此时head为nil或者值不等于val
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
