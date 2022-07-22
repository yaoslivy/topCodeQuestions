package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// return trimBSTTraverse2(root, low, high)

	// return trimBSTRecursion(root, low, high)
	return trimBSTTraverse(root, low, high)
}

//迭代方式
func trimBSTTraverse(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 先处理root
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low { //当前根节点和左子树的都不符合条件
			root = root.Right
		} else { //当前根节点和右子树都不符合条件
			root = root.Left
		}
	}
	//根节点在区间中，则左子树所有节点一定小于high，右子树所有节点一定大于low，判断左右子树符合条件的值
	cur := root
	for cur != nil {
		//处理左子树中小于low的节点
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root
	for cur != nil {
		//处理右子树中大于high的节点
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}

//递归方式
func trimBSTRecursion(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low { //当前节点的值小于可取值区间，则当前节点和左子树上节点都不符合
		right := trimBSTRecursion(root.Right, low, high)
		return right
	}
	if root.Val > high { //当前节点的值大于可取值区间，则当前节点和右子树上节点都不符合
		left := trimBSTRecursion(root.Left, low, high)
		return left
	}
	root.Left = trimBSTRecursion(root.Left, low, high)
	root.Right = trimBSTRecursion(root.Right, low, high)
	return root
}

//迭代方式 全部遍历
func trimBSTTraverse2(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//直接遍历所有节点，判断是否需要删除
	if root.Val < low || root.Val > high { //当前根需要删除
		root = deleteRootNode(root)
		root = trimBSTTraverse(root, low, high)
	}
	if root != nil {
		root.Left = trimBSTTraverse(root.Left, low, high)
		root.Right = trimBSTTraverse(root.Right, low, high)
	}
	return root
}

//删除二叉搜索树的根节点，并返回重组后的二叉搜索树
//策略：在根节点右子树不为nil时，将当前根的左子树作为当前根右子树最左边节点的左孩子
func deleteRootNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	cur := root.Right //找到根节点的右子树中最左边的节点
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = root.Left
	return root.Right
}
