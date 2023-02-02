package main

import "math"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 通过链表长度处理
	// return getIntersectionNodeByLen(headA, headB)
	// 通过遍历两遍链表处理
	return getIntersectionNodeTraverse(headA, headB)
}

func getIntersectionNodeTraverse(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	// a+c + b
	// b+c + a
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func getIntersectionNodeByLen(headA, headB *ListNode) *ListNode {
	ALen := getLen(headA)
	BLen := getLen(headB)
	var fast, slow *ListNode
	if ALen > BLen {
		fast = headA
		slow = headB
	} else {
		fast = headB
		slow = headA
	}
	// fast先走ALen-BLen
	for i := 0; i < int(math.Abs(float64(ALen-BLen))); i++ {
		fast = fast.Next
	}
	//     a1(slow)->a2->c1->c2->c3
	// b1->b2(fast)->b3->c1->c2->c3
	// fast和slow同时走
	for fast != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

func getLen(head *ListNode) int {
	if head == nil {
		return 0
	}
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
