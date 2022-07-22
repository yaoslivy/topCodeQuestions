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
	return isSymmetricLevelTaverse(root)
}

//迭代判断是否对称
func isSymmetricLevelTaverse(root *TreeNode) bool {
	if root == nil {
		return true
	}
	Q := make([]*TreeNode, 0)
	cur := root
	//一次性入队两个，出队时也一次性出两个
	Q = append(Q, cur.Left)
	Q = append(Q, cur.Right)
	for len(Q) != 0 {
		if len(Q) < 2 {
			return false
		}
		leftNode := Q[0]
		rightNode := Q[1]
		Q = Q[2:]
		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}
		//入队左右子节点
		Q = append(Q, leftNode.Left)
		Q = append(Q, rightNode.Right)
		Q = append(Q, leftNode.Right)
		Q = append(Q, rightNode.Left)
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
