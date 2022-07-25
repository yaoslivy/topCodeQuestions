package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// if root == nil {
	//     return true
	// }
	// return isSymmetricTree(root.Left, root.Right)
	return isSymmetricTraverse(root)
}

//迭代解法
func isSymmetricTraverse(root *TreeNode) bool {
	if root == nil {
		return true
	}
	Q := make([]*TreeNode, 0)
	cur := root
	Q = append(Q, cur.Left) //一次入队两个，一次出对两个来比较
	Q = append(Q, cur.Right)
	for len(Q) != 0 {
		size := len(Q)
		for i := 0; i < size; i++ {
			left := Q[0]
			right := Q[1]
			Q = Q[2:]
			if left == nil && right == nil { //满足情况
				break
			}
			if left == nil || right == nil {
				return false
			}
			if left.Val != right.Val {
				return false
			}
			//将当前两节点的的左右子树按照比较顺序入队
			Q = append(Q, left.Left)
			Q = append(Q, right.Right)
			Q = append(Q, left.Right)
			Q = append(Q, right.Left)
		}
	}
	return true
}

//递归判断左右子树是否是对称
func isSymmetricTree(leftTree, rightTree *TreeNode) bool {
	//两个子树为nil
	if leftTree == nil && rightTree == nil {
		return true
	}
	//一个子树为nil
	if leftTree == nil || rightTree == nil {
		return false
	}
	//两个子树都不为nil
	return leftTree.Val == rightTree.Val && isSymmetricTree(leftTree.Left, rightTree.Right) && isSymmetricTree(leftTree.Right, rightTree.Left)
}
