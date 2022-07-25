package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// return deleteNodeRecursion(root, key)

	return deleteNodeTraverse(root, key)
}

//迭代
func deleteNodeTraverse(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	//处理当前根节点
	if root.Val == key {
		return deleteRootNode(root)
	}
	cur := root
	var pre *TreeNode //记录cur的父节点
	for cur != nil && cur.Val != key {
		pre = cur
		if key < cur.Val { //key在左子树中
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur != nil { //树中存在key
		//找到key的节点cur
		if pre.Val > cur.Val { //cur在pre的左边
			pre.Left = deleteRootNode(cur)
		} else {
			pre.Right = deleteRootNode(cur)
		}
	}

	return root
}

//删除当前根节点，并返回重组后的二叉搜索树
func deleteRootNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//当前根节点右子树为nil，直接返回左子树
	if root.Right == nil {
		return root.Left
	}
	//根节点右子树不为nil，将根节点的左子树作为根节点右子树最左边节点的最孩子
	cur := root.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = root.Left
	return root.Right
}

//递归
func deleteNodeRecursion(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key { //删除当前根节点
		//如果当前根节点右子树为nil或者左右子树都为nil
		if root.Right == nil {
			return root.Left
		}
		//右子树不为nil，则将当前根节点的左子树作为当前根节点右子树最左边节点的左孩子，
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}
	//在左右子树中找到key删除，并返回重组后的二叉搜索树
	root.Left = deleteNodeRecursion(root.Left, key)
	root.Right = deleteNodeRecursion(root.Right, key)
	return root
}
