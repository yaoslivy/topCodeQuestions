package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// return isSameTreeTraverse(p, q)
	return isSameTreeTraverse(p, q)
}

// 递归思路
func isSameTreeTraverse(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 迭代思路
func isSameTreeRecursion(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// 使用一个队列记录
	Q := make([]*TreeNode, 0)
	Q = append(Q, p)
	Q = append(Q, q)
	for len(Q) != 0 {
		//出两个节点比较
		p = Q[0]
		Q = Q[1:]
		q = Q[0]
		Q = Q[1:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		Q = append(Q, p.Left)
		Q = append(Q, q.Left)
		Q = append(Q, p.Right)
		Q = append(Q, q.Right)
	}
	return true
}
