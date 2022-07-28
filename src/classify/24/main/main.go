package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// return swapPairsTraverse(head)

	return swapPairsRecursion(head)
}

//递归方式
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// head->2->next
	// head<->2 next
	// 2->head->next
	res := head.Next
	next := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairsRecursion(next)

	return res
}

//迭代方式
func swapPairsTraverse(head *ListNode) *ListNode {
	// 1->2->...
	// dummyHead(pre)->cur->2->...
	// pre->2->next... //改变2.Next时需要提前记录next位置
	// pre->2->cur next...
	// pre->2->cur->next->...
	// dummyHead->2->pre cur->
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	cur := pre.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next.Next

		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next

		pre = cur
		cur = next
	}
	return dummyHead.Next
}
