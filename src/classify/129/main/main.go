package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	// return sumNumbersTraverse(root)

	res = 0
	sumNumbersRecursion(root, []int{})
	return res
}

//回溯思路
var res int

func sumNumbersRecursion(root *TreeNode, oneRes []int) {
	if root.Left == nil && root.Right == nil {
		oneRes = append(oneRes, root.Val)
		//count a number
		for i := 0; i < len(oneRes); i++ {
			res += oneRes[len(oneRes)-1-i] * int(math.Pow(10, float64(i)))
		}
	}
	oneRes = append(oneRes, root.Val)
	if root.Left != nil {
		sumNumbersRecursion(root.Left, oneRes)
	}
	if root.Right != nil {
		sumNumbersRecursion(root.Right, oneRes)
	}
	oneRes = oneRes[:len(oneRes)-1]
}

//迭代思路
func sumNumbersTraverse(root *TreeNode) int {
	// 使用栈来记录后序遍历栈，当遍历到叶子节点时，生成数字
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	res := 0
	for len(S) != 0 || cur != nil {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
				continue
			}
			if cur.Left == nil && cur.Right == nil {
				//Reaching a leaf node.
				oneRes := 0
				for i := 0; i < len(S); i++ {
					oneRes += S[len(S)-i-1].Val * int(math.Pow(10, float64(i)))
				}
				res += oneRes
			}
			pre = cur
			cur = nil
			S = S[:len(S)-1]
		}
	}
	return res
}
