package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle3(head *ListNode) bool {
	if head == nil {
		return false
	}
	mm := make(map[*ListNode]ListNode, 0) //使用map存储key，后续遍历时判断是否存在
	for head != nil {
		if _, ok := mm[head]; ok {
			return true
		}
		mm[head] = ListNode{}
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true

}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	pre := head
	p := pre
	isFirst := false
	for pre != nil && p != nil {
		if isFirst == true && pre == p {
			return true
		}
		isFirst = true
		pre = pre.Next
		if p.Next != nil {
			p = p.Next.Next
		} else {
			return false
		}
	}
	return false
}
