package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// return treePathPreOrder(root)
	if root == nil {
		return []string{}
	}
	res = make([]string, 0)
	binaryTreePathsOrder(root, "")
	return res

}

//回溯求解
var res []string

func binaryTreePathsOrder(root *TreeNode, oneRes string) {
	if root.Left == nil && root.Right == nil { //找到叶子节点
		oneRes += strconv.Itoa(root.Val)
		res = append(res, oneRes)
		return
	}
	oneResLen := len(oneRes)
	oneRes += strconv.Itoa(root.Val) + "->"
	//将当前节点加入
	if root.Left != nil {
		binaryTreePathsOrder(root.Left, oneRes)
	}
	if root.Right != nil {
		binaryTreePathsOrder(root.Right, oneRes)
	}
	oneRes = oneRes[:oneResLen]
}

//通过后序遍历方式求解
func treePathPreOrder(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	cur := root
	var pre *TreeNode
	S := make([]*TreeNode, 0)
	res := make([]string, 0)
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
			} else {
				if cur.Left == nil && cur.Right == nil { //cur为叶子节点
					oneRes := ""
					for i := 0; i < len(S); i++ {
						oneRes += strconv.Itoa(S[i].Val)
						if i < len(S)-1 {
							oneRes += "->"
						}
					}
					res = append(res, oneRes)
				}
				S = S[:len(S)-1]
				pre = cur
				cur = nil
			}
		}
	}
	return res
}
