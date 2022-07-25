package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	// return minDepthPostOrder(root)
	//return minDepthLevelTraverse(root)

	minH = math.MaxInt
	minDepthRecursion(root, 1)
	if minH == math.MaxInt {
		return 0
	}
	return minH
}

//递归解法
var minH int //保存最小高度
func minDepthRecursion(root *TreeNode, curH int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if minH > curH { //到叶子节点且高度最小
			minH = curH
		}
	}
	minDepthRecursion(root.Left, curH+1)
	minDepthRecursion(root.Right, curH+1)
}

//迭代解法
func minDepthLevelTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Q := make([]*TreeNode, 0)
	cur := root
	Q = append(Q, cur)
	res := 0
	for len(Q) != 0 {
		size := len(Q)
		res++
		for i := 0; i < size; i++ {
			cur = Q[0]
			Q = Q[1:]
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
			//当找到一个节点的左右孩子都为nil，则找到了最小深度
			if cur.Left == nil && cur.Right == nil {
				return res
			}
		}
	}
	return res
}

//递归解法
func minDepthPostOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := minDepthPostOrder(root.Left)
	rightD := minDepthPostOrder(root.Right)

	//一个分支为nil的情况
	if root.Left == nil && root.Right != nil {
		return rightD + 1
	}
	if root.Left != nil && root.Right == nil {
		return leftD + 1
	}
	//分支都不为nil，或都为nil
	return min(leftD, rightD) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
