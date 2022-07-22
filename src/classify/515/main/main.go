package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	Q := make([]*TreeNode, 0)
	p := root
	Q = append(Q, p)
	res := make([]int, 0)
	for len(Q) != 0 {
		//在当层节点个数内寻找最大值
		size := len(Q)
		maxVal := math.MinInt //当层节点最大值
		for size != 0 {
			p = Q[0]
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
			maxVal = max(maxVal, p.Val)
			size--
		}
		res = append(res, maxVal)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
