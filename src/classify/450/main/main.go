package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// return deleteNodeTraverse(root, key)

	return deleteNodeRecursion(root, key)
}

//递归方式
func deleteNodeRecursion(root *TreeNode, key int) *TreeNode {
	if root == nil { //没找到删除节点，直接返回nil
		return nil
	}
	if root.Val == key {
		//当前删除节点存在左右子树一端为nil，将左右子树替代当前节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else { //左右子树都不为nil
			//找到右子树最左边的节点
			next := root.Right
			for next.Left != nil {
				next = next.Left
			}
			//找到右子树最左边的节点后，让当前节点的左子树成为最左边节点的左孩子，然后返回当前节点的右孩子
			next.Left = root.Left
			return root.Right
		}
	}
	if root.Val > key { //删除的key在左边
		root.Left = deleteNodeRecursion(root.Left, key)
	} else {
		root.Right = deleteNodeRecursion(root.Right, key)
	}
	return root
}

// 迭代方式
func deleteNodeTraverse(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	cur := root
	var pre *TreeNode //记录上一个访问节点
	for cur != nil {
		if cur.Val > key {
			pre = cur
			cur = cur.Left
		} else if cur.Val < key {
			pre = cur
			cur = cur.Right
		} else {
			break
		}
	}
	if cur == root { //需要删除头节点
		return deleteRootNode(root)
	}
	if cur != nil { //该删除节点存在于左右子树中
		//找到需要删除节点cur
		if pre.Val > cur.Val { // cur在pre左边
			pre.Left = deleteRootNode(cur)
		} else {
			pre.Right = deleteRootNode(cur)
		}
	}
	return root
}

//删除二叉搜索树中的根节点，并返回删除后的剩余节点构成的二叉搜索树
//策略：在根节点右孩子不为nil时，将根节点的左子树作为根节点右子树中最左边孩子的左孩子节点，并返回根节点的右孩子
func deleteRootNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root.Left
	}
	cur := root.Right
	for cur.Left != nil { //一直找到根节点右子树中的最左边的节点
		cur = cur.Left
	}
	cur.Left = root.Left
	return root.Right
}
