package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 原链表上分割方式
	// reorderListByDivide(head)
	// 双端队列方式
	reorderListByQueue(head)
}

func reorderListByQueue(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//先将所有节点加入队列中
	Q := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		Q = append(Q, cur)
		cur = cur.Next
	}
	//从队列中取值
	cur = head
	var next *ListNode
	Q = Q[1:]
	flag := 1 // 偶数从队头取值，奇数从队尾取值
	for len(Q) != 0 {
		if flag%2 == 0 {
			next = Q[0]
			Q = Q[1:]
		} else {
			next = Q[len(Q)-1]
			Q = Q[:len(Q)-1]
		}
		flag++
		cur.Next = next
		cur = cur.Next
	}
	cur.Next = nil //注意结尾需要指向nil， 因为是在原链表上操作
}

func reorderListByDivide(head *ListNode) {
	//先使用快慢指针，fast走两步，slow走一步，当fast到末尾时，slow分割即是当前链表中心位置
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	var pre *ListNode //指向slow前一个节点，用来分割单链表
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	// 1->2->3->4->5
	// 1->2(pre)->3(slow)->4->5(fast)
	// 1->2->3->4
	// 1->2(pre)->3(slow)->4 fast
	pre.Next = nil
	cur1 := head
	//反转后半段链表
	cur2 := reverseList(slow)
	for cur1 != nil && cur2 != nil {
		// 1(cur1)->2
		// 5(cur2)->4->3
		next := cur1.Next
		cur1.Next = cur2
		// 2(next)
		// 1(cur1)->5(cur2)->4->3
		cur1 = cur2
		cur2 = next
		// 2(next, cur2)
		// 1->5(cur1)->4->3
		// nil(next)
		// 1->5(cur1)->2(cur2)->4->3
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseList(head.Next)
	//1->2
	head.Next.Next = head
	head.Next = nil

	return tail
}
