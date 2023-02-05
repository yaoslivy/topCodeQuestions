package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// return mergeTwoListsTraverse(list1, list2)
	return mergeTwoListsRecursion(list1, list2)
}

// 递归方式
func mergeTwoListsRecursion(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	// 递归的选取出两个链表中较小的头节点，并将剩余节点和另一条链表合并
	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursion(list1, list2.Next)
		return list2
	}
}

// 迭代方式
func mergeTwoListsTraverse(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	// dummyHead->1->1->2->3->4(pre) (list1)
	//      4(list2)
	// pre->
	dummyHead := &ListNode{}
	pre := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	// 最多还有一个链表的剩余没有合并完
	if list1 != nil {
		pre.Next = list1
	} else {
		pre.Next = list2
	}
	return dummyHead.Next
}
