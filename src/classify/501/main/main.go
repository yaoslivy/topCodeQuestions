package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	//二叉搜索树的特点是：中序遍历序列是一段从小到大的序列，如果出现重复值，则一定是相邻的
	// preCount, curCount = 0, 0
	// pre = nil
	// res = make([]int, 0)
	// findModeRecursion(root)
	// return res

	return findModeTraverse(root)
}

//迭代方式
func findModeTraverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	S := make([]*TreeNode, 0)
	cur := root
	res := make([]int, 0) //结果集
	var pre *TreeNode     //记录前一个访问节点
	var preCount int      //记录上一个加入结果集的值的出现个数
	var curCount int      //记录当前值出现的个数
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			// 统计当前值出现的个数
			if pre == nil {
				curCount = 1
			} else if pre.Val == cur.Val {
				curCount++
			} else {
				curCount = 1
			}
			// 如果当前值出现的个数>= 结果集合中值出现的个数
			if curCount > preCount {
				res = make([]int, 0)
				res = append(res, cur.Val)
				preCount = curCount
			} else if curCount == preCount {
				res = append(res, cur.Val)
			}
			pre = cur
			cur = cur.Right
		}
	}
	return res
}

//递归方式
var preCount int  //前一个众数的出现个数
var curCount int  //当前值的出现个数
var pre *TreeNode // 记录前一个访问节点
var res []int     //记录结果集
func findModeRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	findModeRecursion(root.Left)
	//确定当前值出现的个数
	if pre == nil {
		curCount = 1
	} else if pre.Val == root.Val {
		curCount++
	} else {
		curCount = 1
	}
	//如果当前值>=上一个加入结果集的值出现的个数，则加入结果集
	if curCount > preCount { //当前值出现的个数超过了之前值出现的个数
		res = make([]int, 0)
		res = append(res, root.Val)
		preCount = curCount
	} else if curCount == preCount { //相等则直接加入结果集
		res = append(res, root.Val)
	}

	pre = root
	findModeRecursion(root.Right)
}
