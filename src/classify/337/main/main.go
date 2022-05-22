package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob2(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

//长度为2的slice，0:记录不偷该节点所得到的最大金钱，1:记录偷该节点所得到的最大金钱
func robTree(cur *TreeNode) []int {
	if cur == nil { //遇到空节点，无论偷与不偷都是取值0
		return []int{0, 0}
	}
	left := robTree(cur.Left)
	right := robTree(cur.Right)
	//偷cur
	val1 := cur.Val + left[0] + right[0]
	//不偷cur
	val2 := max(left[0], left[1]) + max(right[0], right[1])
	return []int{val2, val1}
}

var mm = make(map[*TreeNode]int) //记录节点为*TreeNode的子树所能偷到的最大值

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if val, ok := mm[root]; ok {
		return val
	}
	//偷父节点
	val1 := root.Val
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right) //跳过root.Left相当于不考虑左子节点
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right) // 不考虑右子节点
	}
	//不偷父节点
	val2 := rob(root.Left) + rob(root.Right) //考虑root的左右孩子
	mm[root] = max(val1, val2)               //记录以当前节点为子树情况所能偷得最大金额
	return max(val1, val2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
