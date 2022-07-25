package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// return reverseListTraverse(head)

	return reverseListRecursion(nil, head)

	// return reverseListRecursion2(head)
}

//递归方式
// pre->cur
func reverseListRecursion(pre, cur *ListNode) *ListNode {
	//1->2->3
	// <-1 2->3 // 记录2的位置
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre

	return reverseListRecursion(cur, next)
}

func reverseListRecursion2(cur *ListNode) *ListNode {
	if cur == nil {
		return nil
	}
	if cur.Next == nil {
		return cur
	}
	//反转第二个节点往后的节点 1,2,3,4,5
	last := reverseListRecursion2(cur.Next)

	cur.Next.Next = cur
	cur.Next = nil //此时的cur节点为尾节点，next需要指向nil
	return last
}

//迭代方式
func reverseListTraverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre, next *ListNode //保存前一个节点和后一个节点
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
