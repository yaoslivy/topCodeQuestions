package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	head := node1
	k := 2
	p := head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	p = reverseKGroup(head, k)
	fmt.Println()
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	headPre := &ListNode{Next: head} //生成一个新的头节点指向当前头节点
	p := head
	pre := headPre
	for p != nil { //p节点指向第一个节点，后逐渐后移，直至到最后一个节点，在移动的过程中每k步反转k个节点
		tail := pre // 指向第一个节点的前一个节点，这样移动k步后，tail会刚好指向第k个节点，
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return headPre.Next
			}
		}
		//找到需要反转的头尾序列后，进行局部反转，反转前需要记录下一段序列的位置
		nex := tail.Next
		p, tail = reverseOneGroup(p, tail)
		//局部反转完后，需要将前后连接起来
		pre.Next = p
		tail.Next = nex
		//开始下一段序列的操作， pre和tail需要指向下一段序列的前一个位置，p需要指向下一段序列的第一个位置
		pre = tail
		p = nex
	}
	return headPre.Next
}

func reverseOneGroup(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	p := head
	lastNode := tail.Next //记录p移动的终止位置，不能直接使用tail.Next，因为p会移动到tail，然后让tail.Next指向pre
	for p != lastNode {   //p移动完区间内的所有节点
		nex := p.Next //每次将p指向前一个节点时，都需要记录p后一个节点位置，方便p后移
		p.Next = pre
		pre = p
		p = nex
	}
	return tail, head
}
