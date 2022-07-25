package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// return binaryTreePathsTraverse(root)

	res = make([]string, 0)
	binaryTreePathsRecursion(root, "")
	return res
}

// 递归方式
var res []string

func binaryTreePathsRecursion(root *TreeNode, oneRes string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil { //找到叶子节点
		oneRes += strconv.Itoa(root.Val)
		res = append(res, oneRes)
		return
	}

	binaryTreePathsRecursion(root.Left, oneRes+strconv.Itoa(root.Val)+"->") //回溯
	binaryTreePathsRecursion(root.Right, oneRes+strconv.Itoa(root.Val)+"->")
}

//迭代方式
func binaryTreePathsTraverse(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode //记录上一次访问的节点
	res := make([]string, 0)
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
				continue
			}
			//到达叶子节点
			oneRes := ""
			if cur.Left == nil && cur.Right == nil {
				for i := 0; i < len(S); i++ {
					if i == 0 {
						oneRes += strconv.Itoa(S[i].Val)
					} else {
						oneRes += "->" + strconv.Itoa(S[i].Val)
					}
				}
				res = append(res, oneRes)
			}
			S = S[:len(S)-1]
			pre = cur
			cur = nil
		}
	}
	return res
}
