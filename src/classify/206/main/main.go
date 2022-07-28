package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// return reverseListTraverse(head)

	return reverseListRecursion2(nil, head)

	// return reverseListRecursion2(head)
}

//递归方式
// pre->cur
func reverseListRecursion2(pre, cur *ListNode) *ListNode {
	//1->2->3
	// <-1 2->3 // 记录2的位置
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre

	return reverseListRecursion2(cur, next)
}

// 递归方式
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 1->2->3
	// head->next
	// 1<->2 3 // 改变2.Next指向前，3的位置已经记录
	// <-1->2 3
	next := reverseListRecursion(head.Next)

	head.Next.Next = head
	head.Next = nil
	return next
}

//迭代方式
func reverseListTraverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 1->2->3
	// nil cur->...
	// nil<-cur next->... //改变cur.Next时需要先记录next值，不然无法找到下一个节点
	// nil<-pre cur->...
	cur := head
	var pre *ListNode //记录cur的上一个节点
	for cur != nil {
		next := cur.Next //记录cur的下一个节点

		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}
