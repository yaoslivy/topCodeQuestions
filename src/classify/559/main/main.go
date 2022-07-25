package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	// return maxDepthByLevelTraverse(root)
	return maxDepthByPostOrder(root)

	// maxH = 0
	// maxDepthRecursion(root, 1)
	// return maxH
}

// 递归方式
var maxH int

func maxDepthRecursion(root *Node, curH int) {
	if root == nil {
		return
	}

	for i := 0; i < len(root.Children); i++ {
		maxDepthRecursion(root.Children[i], curH+1)
	}
	if maxH < curH {
		maxH = curH
	}
}

//通过递归方式
func maxDepthByPostOrder(root *Node) int {
	if root == nil {
		return 0
	}
	maxD := 0 //当前子树中最大的高度
	for i := 0; i < len(root.Children); i++ {
		tempD := maxDepthByPostOrder(root.Children[i])
		maxD = max(maxD, tempD)
	}
	return maxD + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//通过层序遍历的方式
func maxDepthByLevelTraverse(root *Node) int {
	if root == nil {
		return 0
	}
	Q := make([]*Node, 0)
	cur := root
	res := 0
	Q = append(Q, cur)
	for len(Q) != 0 {
		size := len(Q)
		for i := 0; i < size; i++ {
			cur = Q[0]
			Q = Q[1:]
			for j := 0; j < len(cur.Children); j++ {
				Q = append(Q, cur.Children[j])
			}
		}
		res++
	}
	return res
}
