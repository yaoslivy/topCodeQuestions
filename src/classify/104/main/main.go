package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// return maxDepthByLevelTraverse(root)
	return maxDepthByPostOrder(root)
}

func maxDepthByPostOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := maxDepthByPostOrder(root.Left)
	rightD := maxDepthByPostOrder(root.Right)
	//做后序遍历，通过递归函数的返回值做计算
	return max(leftD, rightD) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//通过层序遍历获得二叉树层数
func maxDepthByLevelTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Q := make([]*TreeNode, 0)
	cur := root
	Q = append(Q, cur)
	res := 0
	for len(Q) != 0 {
		for size := len(Q); size > 0; size-- {
			cur = Q[0]
			Q = Q[1:]
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
		}
		res++
	}
	return res
}
