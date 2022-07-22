package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	//二叉搜索树的特点是中序遍历序列是一段从小到大的序列
	// pre = nil
	// return isValidBSTRecursion(root)

	return isValidBSTTraverse(root)
}

//迭代判断
func isValidBSTTraverse(root *TreeNode) bool {
	if root == nil {
		return true
	}
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode //记录前一个访问节点
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			if pre != nil && pre.Val >= cur.Val {
				return false
			}
			pre = cur
			cur = cur.Right
		}
	}
	return true
}

//递归判断
var pre *TreeNode //记录前一个访问节点
func isValidBSTRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isValidBSTRecursion(root.Left)
	if pre != nil {
		if pre.Val >= root.Val { //中序遍历序列不是正序
			return false
		}
	}
	pre = root
	right := isValidBSTRecursion(root.Right)

	return left && right
}
