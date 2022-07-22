package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// return hasPathSumPostTraverse(root, targetSum)
	if root == nil {
		return false
	}
	return hasPathSumTraverse(root, targetSum-root.Val)
}

//递归求解
func hasPathSumTraverse(root *TreeNode, countSum int) bool {
	//到叶子节点，且countSum为0，说明找到路径，否则返回false，回溯找下一跳路径
	if root.Left == nil && root.Right == nil {
		if countSum == 0 {
			return true
		}
		return false
	}
	if root.Left != nil {
		if hasPathSumTraverse(root.Left, countSum-root.Left.Val) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSumTraverse(root.Right, countSum-root.Right.Val) {
			return true
		}
	}
	return false
}

//后序迭代求解
func hasPathSumPostTraverse(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
			} else {
				if cur.Left == nil && cur.Right == nil {
					//当前节点是叶子节点，栈中保留了从根节点到叶子节点的路径
					sum := 0
					for _, node := range S {
						sum += node.Val
					}
					if sum == targetSum {
						return true
					}
				}
				S = S[:len(S)-1]
				pre = cur
				cur = nil
			}
		}
	}
	return false
}
