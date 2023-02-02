package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	// 记录头节点，然后按k组反转
	dummyHead := &ListNode{Next: head}
	// dummyHead->1->2->3->4->5
	gHead := head    //找到一组节点的头节点
	pre := dummyHead //记录一组节点的头节点的前一个节点
	// (pre)->1(gHead)->2(gTail)->3(next)->4->5
	// (pre)->2(gTail)->1(gHead) 3(next)->4->5
	// (dummyHead)->2->1(gHead,pre) 3(next)->4->5
	// (dummyHead)->2->1(pre)->3(gHead)->4(gTail)->5
	for gHead != nil {
		gTail := pre
		for i := 0; i < k; i++ {
			gTail = gTail.Next
			if gTail == nil {
				return dummyHead.Next
			}
		}
		next := gTail.Next //下一组节点
		gTail.Next = nil
		pre.Next = reverseList(gHead) //此时cur即是当前组的最后一个节点
		pre = gHead
		gHead.Next = next
		gHead = next
	}

	return dummyHead.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tail
}
