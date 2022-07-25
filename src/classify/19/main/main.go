package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//设置两个指针，隔n个节点，如此，下一节点就是需要删除的节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast := dummyHead
	fast = fast.Next
	for fast != nil && n != 0 {
		fast = fast.Next
		n--
	}
	// 1->2->3    // n = 2时， fast指向3，slow.Next是1，当fast为nil时,slow指向1，删除后一个节点
	slow := dummyHead //slow和fast间隔n个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
