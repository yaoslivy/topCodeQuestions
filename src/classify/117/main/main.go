package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	Q := make([]*Node, 0)
	cur := root
	Q = append(Q, cur)
	for len(Q) != 0 {
		size := len(Q)
		for i := 0; i < size; i++ {
			cur = Q[0]
			Q = Q[1:]
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
			//改变next指针指向
			if i < size-1 { //不是该层的最后一个节点
				cur.Next = Q[0]
			}
		}
	}
	return root
}
