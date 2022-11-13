package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	// res = make([][]int, 0)
	// if root == nil {
	//     return res
	// }
	// pathSumTraverse(root, targetSum-root.Val, []int{root.Val})
	// return res

	// return pathSumTraversePostTraverse(root, targetSum)

	if root == nil {
		return nil
	}
	res = make([][]int, 0)
	pathSumRecursion(root, targetSum, []int{})
	return res
}

//后序遍历求解
func pathSumTraversePostTraverse(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode // 记录上一次访问的节点
	res := make([][]int, 0)
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
			} else {
				//当前节点是叶子节点
				if cur.Left == nil && cur.Right == nil {
					sum := 0
					oneRes := make([]int, 0)
					for _, node := range S {
						sum += node.Val
						oneRes = append(oneRes, node.Val)
					}

					if sum == targetSum { //找到一个结果集
						copyArr := make([]int, len(oneRes))
						copy(copyArr, oneRes)
						res = append(res, copyArr)
					}
				}
				S = S[:len(S)-1]
				pre = cur
				cur = nil
			}
		}
	}
	return res
}

//递归求解
var res [][]int

func pathSumTraverse(root *TreeNode, countSum int, oneRes []int) {
	if root.Left == nil && root.Right == nil {
		if countSum == 0 { //已到达叶子节点，且路径总和为目标值
			copyArr := make([]int, len(oneRes))
			copy(copyArr, oneRes)
			res = append(res, copyArr)
		}
		return
	}
	if root.Left != nil {
		oneRes = append(oneRes, root.Left.Val)
		pathSumTraverse(root.Left, countSum-root.Left.Val, oneRes)
		oneRes = oneRes[:len(oneRes)-1]
	}
	if root.Right != nil {
		oneRes = append(oneRes, root.Right.Val)
		pathSumTraverse(root.Right, countSum-root.Right.Val, oneRes)
		oneRes = oneRes[:len(oneRes)-1]
	}
}

//递归思路2
func pathSumRecursion(root *TreeNode, targetSum int, oneRes []int) {
	if root == nil {
		return
	}
	oneRes = append(oneRes, root.Val)
	//判断当前节点是否是叶子节点
	if root.Left == nil && root.Right == nil {
		sum := 0
		for _, val := range oneRes {
			sum += val
		}
		if sum == targetSum {
			copyArr := make([]int, len(oneRes))
			copy(copyArr, oneRes)
			res = append(res, copyArr)
		}
		return
	}
	//向左子树和右子树找叶子节点
	pathSumRecursion(root.Left, targetSum, oneRes)
	pathSumRecursion(root.Right, targetSum, oneRes)
	oneRes = oneRes[:len(oneRes)-1]
}
