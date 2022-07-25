package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	//设置fast，slow都指向head，fast走两步，slow走一步，
	// 如果有环则fast和slow一定会在环中相遇
	// 相遇后在设置一个节点从head出发，一定可以和slow相遇
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			resNode := head
			for resNode != slow {
				resNode = resNode.Next
				slow = slow.Next
			}
			return resNode
		}
	}
	return nil
}
