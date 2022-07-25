package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// return swapPairsTraverse(head)

	return swapPairsRecursion(head)
}

//递归方式
func swapPairsRecursion(head *ListNode) *ListNode {
	// 1->2->3->4
	if head == nil || head.Next == nil {
		return head
	}
	//记录一个组合中的第一节点
	first := head.Next                         //2
	head.Next = swapPairsRecursion(first.Next) //1->() //将1.Next指向下一个组合的第一个节点
	first.Next = head                          //2->1

	return first //返回一个交换组合中的第一个节点
}

// 迭代方式
func swapPairsTraverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //一个节点或为nil
		return head
	}
	// 1->2->3->4 //一个组合两个值，需要记录一个组合前一个位置，和后一个位置，以及当前组合的第一个位置
	// 1->2->3->4 // pre指向1的前一个位置，cur指向1，next指向3
	// pre.Next = 2  //前一个组合不为nil时
	// 2->1
	// 1->3
	// 2->1->3->4
	root := head.Next
	cur := head
	var pre *ListNode
	for cur != nil && cur.Next != nil {
		next := cur.Next.Next //3

		if pre != nil {
			pre.Next = cur.Next // ()->2
		}
		cur.Next.Next = cur //()->2->1
		cur.Next = next     //()->2->1->3
		//2->1->3->4
		pre = cur
		cur = next
	}
	return root
}
