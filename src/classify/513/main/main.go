package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	// return findBottomLeftValueLevelTraverse(root)

	MaxH = 0
	res = 0
	findBottomLeftValueTraverse(root, 1)
	return res
}

//递归方式求解 当遍历到下一层的第一个，且第一个值为叶子节点时就更新结果和当前最大层数
var MaxH int
var res int

func findBottomLeftValueTraverse(root *TreeNode, curH int) { //当前节点和当前高度
	if root == nil {
		return
	}
	findBottomLeftValueTraverse(root.Left, curH+1)
	findBottomLeftValueTraverse(root.Right, curH+1)
	//已经遍历完当前层，到下一层第一个叶子节点时
	if curH > MaxH && root.Left == nil && root.Right == nil {
		MaxH = curH
		res = root.Val
	}
}

//层序遍历方式求解
func findBottomLeftValueLevelTraverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := root
	Q := make([]*TreeNode, 0)
	Q = append(Q, cur)
	res := 0
	for len(Q) != 0 {
		size := len(Q)
		for i := 0; i < size; i++ {
			cur = Q[0]
			Q = Q[1:]
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
			// 该层第一个节点，且该节点左右子孩子为nil
			if i == 0 && cur.Left == nil && cur.Right == nil {
				res = cur.Val
			}
		}
	}
	return res
}
