package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// return hasPathSumPostTraverse(root, targetSum)

	/* if root == nil {
		return false
	}
	return hasPathSumTraverse(root, targetSum-root.Val) */

	res = false
	hasPathSumRecursion(root, targetSum, 0)
	return res
}

//递归解法
var res bool

func hasPathSumRecursion(root *TreeNode, targetSum int, curSum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil { //找到叶子节点
		if curSum+root.Val == targetSum {
			res = true
		}
		return
	}
	hasPathSumRecursion(root.Left, targetSum, curSum+root.Val)
	hasPathSumRecursion(root.Right, targetSum, curSum+root.Val)
}

//递归方式 回溯
func hasPathSumRecursion2(root *TreeNode, targetSum int) bool {
	if root != nil && targetSum-root.Val == 0 && root.Left == nil && root.Right == nil { //如果目标值可以被减到0
		return true
	}
	var left, right bool
	if root != nil {
		left = hasPathSumRecursion2(root.Left, targetSum-root.Val)
		right = hasPathSumRecursion2(root.Right, targetSum-root.Val)
	}

	return left || right
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
