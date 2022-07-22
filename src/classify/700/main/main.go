package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// if root == nil {
	//     return root
	// }
	// return searchBSTRecursion(root, val)

	return searchBSTTraverse(root, val)
}

//迭代法
func searchBSTTraverse(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		}
		if val > root.Val { //往右边找
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

//递归求解
func searchBSTRecursion(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	if val > root.Val { //在根节点右边
		if root.Right != nil {
			return searchBSTRecursion(root.Right, val)
		}
	} else {
		if root.Left != nil {
			return searchBSTRecursion(root.Left, val)
		}
	}
	return nil
}
