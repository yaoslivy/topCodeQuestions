package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	//先通过中序遍历，将二叉搜索树转换为有序数组
	nodes = make([]int, 0)
	inOrder(root)
	// 根据中序遍历序列转换为二叉搜索树
	return buildBiTree(nodes, 0, len(nodes)-1)
}

func buildBiTree(nodes []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 1 2 3 4
	//每次取中间元素作为根节点
	mid := left + (right-left)/2
	root := &TreeNode{Val: nodes[mid]}
	root.Left = buildBiTree(nodes, left, mid-1)
	root.Right = buildBiTree(nodes, mid+1, right)
	return root
}

var nodes []int

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	nodes = append(nodes, root.Val)
	inOrder(root.Right)
}
