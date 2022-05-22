package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: node1}
	list1 := &ListNode{Val: 1, Next: node2}

	node3 := &ListNode{Val: 4, Next: nil}
	node4 := &ListNode{Val: 3, Next: node3}
	list2 := &ListNode{Val: 1, Next: node4}

	head := mergeTwoLists(list1, list2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val { //因为两个链表有序，若当前list1的值<=list2的值，则当前list1的值<=所有list2的值，所以list1.next应该指向list1之后的节点和list2组合起来的最小的链表的头节点
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	preHead := &ListNode{Val: -1, Next: nil}
	p := preHead //当前p节点前的序列都是有序的，每次比较确定下一个值是在list1，还是list2，切换指向，p再移动
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}

	return preHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val <= list2.Val { //记录两个有序链表合并后的头节点
		head = list1
	} else {
		head = list2
	}
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			nex := list1.Next
			for nex != nil && nex.Val <= list2.Val { //当list1的下一个节点值大于list2的节点值时，才将list1的Next指向list2
				list1 = nex
				nex = nex.Next
			}
			list1.Next = list2
			list1 = nex
		} else {
			nex := list2.Next
			for nex != nil && nex.Val <= list1.Val { //当list2的下一个节点值大于list1时，才将list2的Next指向list1
				list2 = nex
				nex = nex.Next
			}
			list2.Next = list1
			list2 = nex
		}
	}
	return head
}
