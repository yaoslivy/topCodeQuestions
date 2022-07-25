package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// return getIntersectionNode1(headA, headB)

	return getIntersectionNode2(headA, headB)
}

//双指针方法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	// a+c +b == b+c +a
	curA, curB := headA, headB
	for curA != curB { //两者先链表一条链表，再遍历第二条链表
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
	//如果没有交点，退出for循环时会同时为nil，因为都遍历了两条链表
	return curA
}

//在两条链表长度相同的情况下查找
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	//求出两条链表的长度，让指向链表长度长的指针先走长度差个节点,
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var fast, slow *ListNode
	if lenA > lenB {
		fast = headA
		slow = headB
	} else {
		fast = headB
		slow = headA
	}
	gaps := math.Abs(float64(lenA - lenB))
	for gaps != 0 {
		gaps--
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
