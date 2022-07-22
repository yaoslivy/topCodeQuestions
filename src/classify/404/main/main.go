package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	// return sumOfLeftLeavesPreTraverse(root)

	return sumOfLeftLeavesPostOrder(root)

	// res = 0
	// sumOfLeftLeavesPreOrder(root)
	// return res
}

//后序递归方式
func sumOfLeftLeavesPostOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftVal := sumOfLeftLeavesPostOrder(root.Left)
	rightVal := sumOfLeftLeavesPostOrder(root.Right)
	midVal := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midVal = root.Left.Val
	}
	return leftVal + rightVal + midVal
}

//通过先中后序遍历来统计左叶子节点
func sumOfLeftLeavesPreTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	S := make([]*TreeNode, 0)
	cur := root
	res := 0
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			//当前节点的左孩子不为nil，且左孩子无左右子树，则该左孩子就为左叶子节点
			if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
				res += cur.Left.Val
			}
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			cur = cur.Right
		}
	}
	return res
}

//先中后序递归方式
var res int

func sumOfLeftLeavesPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	sumOfLeftLeavesPreOrder(root.Left)
	sumOfLeftLeavesPreOrder(root.Right)
}
