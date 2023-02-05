package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	// return connectTraverse(root)
	return connectRecursion(root)
}

// 递归
func connectRecursion(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connectRecursion(root.Left)
	connectRecursion(root.Right)
	return root
}

// 迭代
func connectTraverse(root *Node) *Node {
	if root == nil {
		return root
	}
	Q := make([]*Node, 0)
	p := root
	Q = append(Q, p)
	for len(Q) != 0 {
		//当层节点个数-1，最后一个节点之前的节点都需要指向后一个节点
		//先将当层节点都存在一个slice中
		size := len(Q)
		oneRes := make([]*Node, 0)
		for i := 0; i < size; i++ {
			p = Q[0]
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
			oneRes = append(oneRes, p)
		}
		//oneRes中将最后一个节点之前的节点连接起来
		for i := 0; i < len(oneRes)-1; i++ {
			oneRes[i].Next = oneRes[i+1]
		}
	}
	return root
}
