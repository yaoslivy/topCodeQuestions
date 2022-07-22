package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// return countNodesPostOrder(root)
	return countNodesLevelTraverse(root)
}

//层序遍历方式
func countNodesLevelTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Q := make([]*TreeNode, 0)
	cur := root
	Q = append(Q, cur)
	res := 0
	for len(Q) != 0 {
		cur = Q[0]
		Q = Q[1:]
		res++
		if cur.Left != nil {
			Q = append(Q, cur.Left)
		}
		if cur.Right != nil {
			Q = append(Q, cur.Right)
		}
	}
	return res
}

//递归方式
func countNodesPostOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNode := countNodesPostOrder(root.Left)
	rightNode := countNodesPostOrder(root.Right)

	return 1 + leftNode + rightNode
}
