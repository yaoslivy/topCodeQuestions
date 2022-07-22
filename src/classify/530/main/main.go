package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	// minVal = math.MaxInt
	// pre = nil
	// getMinimumDifferenceRecursion(root)
	// return minVal

	return getMinimumDifferenceTraverse(root)
}

//二叉搜索树的中序遍历序列是一段从小到大的序列，最小差值一定在相邻节点处产生
var minVal int    //记录遍历序列中最小差值
var pre *TreeNode //记录上一个访问节点
func getMinimumDifferenceRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	getMinimumDifferenceRecursion(root.Left)
	if pre != nil && minVal > (root.Val-pre.Val) {
		minVal = root.Val - pre.Val
	}
	pre = root
	getMinimumDifferenceRecursion(root.Right)
}

//中序迭代
func getMinimumDifferenceTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := root
	var pre *TreeNode     //记录前一个访问节点
	minVal := math.MaxInt //记录最小差值
	S := make([]*TreeNode, 0)
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			if pre != nil && minVal >= (cur.Val-pre.Val) {
				minVal = cur.Val - pre.Val
			}
			pre = cur
			cur = cur.Right
		}
	}
	return minVal
}
