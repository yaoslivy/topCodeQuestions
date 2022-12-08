package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//使用快慢指针fast, slow
	//fast一次走两步，slow一次走一步，则fast到末尾时，slow正好到中间
	fast, slow := head, head
	var pre *ListNode //记录慢指针的前一个节点，用来分割链表
	for fast != nil && fast.Next != nil {
		pre = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 1->2->2->1
	// 1>2(pre)->2(slow)->1 fast
	// 1->2->1
	// 1(pre)->2(slow)->1(fast)
	//反转后半部分的链表
	pre.Next = nil
	cur1 := head
	cur2 := reverseList(slow)
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//1(head)->2
	next := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}

// 使用slice记录链表值
func isPalindrome2(head *ListNode) bool {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	if len(arr)%2 == 0 {
		return judge(arr, len(arr)/2-1, len(arr)/2)
	} else {
		return judge(arr, len(arr)/2-1, len(arr)/2+1)
	}
}

//判断是否为回文
func judge(arr []int, left, right int) bool {
	for left >= 0 && right < len(arr) {
		if arr[left] != arr[right] {
			return false
		}
		left--
		right++
	}
	return true
}
