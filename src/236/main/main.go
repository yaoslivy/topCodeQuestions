package main

import (
	"fmt"
)

func main() {
	node1 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 2, Left: node1, Right: node2}
	node4 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5, Left: node4, Right: node3}
	node6 := &TreeNode{Val: 0}
	node7 := &TreeNode{Val: 8}
	node8 := &TreeNode{Val: 1, Left: node6, Right: node7}
	root := &TreeNode{Val: 3, Left: node5, Right: node8}
	fmt.Println(lowestCommonAncestor(root, node5, node8).Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q { //说明子树中存在p或q
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil { // 说明 root的左右子树中都不包含 p,q
		return nil
	}
	if left == nil { // p,q都不在 root的左子树中，直接返回 right
		return right
	}
	if right == nil { //
		return left
	}
	// if left != nil && right != nil { // 说明p,q在root的左右子树，root就是公共节点
	//     return root
	// }
	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) { //递归记录所有节点的父节点
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] { //说明访问到相同节点
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//后序遍历拿到p, q的父序列
	S1, S2 := postTraverse(root, p, q)
	if len(S1) == 0 || len(S2) == 0 {
		return nil
	}
	for len(S1) != len(S2) {
		if len(S1) > len(S2) {
			S1 = S1[:len(S1)-1]
		} else {
			S2 = S2[:len(S2)-1]
		}
	}
	for S1[len(S1)-1] != S2[len(S2)-1] {
		S1 = S1[:len(S1)-1]
		S2 = S2[:len(S2)-1]
	}
	return S1[len(S1)-1]
}

func postTraverse(root, p, q *TreeNode) ([]*TreeNode, []*TreeNode) {
	S := make([]*TreeNode, 0)
	var S1 []*TreeNode
	var S2 []*TreeNode
	r := &TreeNode{} //记录前一个访问节点
	haveP := false
	haveQ := false
	for root != nil || len(S) != 0 {
		if root != nil {
			S = append(S, root)
			root = root.Left
		} else {
			root = S[len(S)-1]
			if root.Right != nil && root.Right != r {
				root = root.Right
			} else {
				if root == p {
					S1 = make([]*TreeNode, len(S))
					copy(S1, S)
					haveP = true
				}
				if root == q {
					S2 = make([]*TreeNode, len(S))
					copy(S2, S)
					haveQ = true
				}
				if haveP && haveQ {
					return S1, S2
				}
				S = S[:len(S)-1]
				r = root
				root = nil
			}
		}
	}
	return nil, nil
}
