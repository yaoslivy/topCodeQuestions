package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// return trimBSTRecursion(root, low, high)

	return trimBSTRecursionTraverse(root, low, high)
}

//迭代方式
func trimBSTRecursionTraverse(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//找到在区间范围内的根节点
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low { //则当前根节点和左子树上的元素都不符合条件
			root = root.Right
		} else {
			root = root.Left
		}
	}
	//根节点在区间范围内，则保证了左子树的所有元素都<high，保证了右子树的所有元素都>low
	//去除左子树中<low的节点
	cur := root
	for cur != nil {
		//找到满足的左子树根节点，不满足则跳到右边，因为左边比根节点的值还小
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left //下一个满足的左子树根节点
	}
	//去除右子树中>high的节点
	cur = root
	for cur != nil {
		//找到满足的右子树的根节点，不满足则跳到左边，因为右边比根节点的值还大
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
	//二叉搜索树特点是中序遍历序列是从小到大的
	if root.Val < low {
		//如果当前根节点的值小于low，则当前根节点和左子树元素都不符合条件，返回符合的右子树
		return trimBSTRecursion(root.Right, low, high)
	}
	if root.Val > high {
		//如果当前根节点值大于high，则当前根节点和右子树元素都不符合条件，返回符合的左子树
		return trimBSTRecursion(root.Left, low, high)
	}
	root.Left = trimBSTRecursion(root.Left, low, high)
	root.Right = trimBSTRecursion(root.Right, low, high)
	return root
}
