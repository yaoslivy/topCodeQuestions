package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// return lowestCommonAncestorRecursion(root, p, q)

	return lowestCommonAncestorTraverse(root, p, q)
}

//迭代方式
func lowestCommonAncestorTraverse(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left //p,q在当前节点的左子树中
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right //p,q在当前节点的右子树中
		} else {
			return cur
		}
	}
	return nil
}

//递归方式
func lowestCommonAncestorRecursion(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorRecursion(root.Left, p, q) //p,q都在当前根节点的左子树中
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorRecursion(root.Right, p, q) //p,q都在当前根节点的右子树中
	}
	return root //p, q分布在当前根节点的左右两边
}
