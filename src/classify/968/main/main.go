package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	res = 0
	if find(root) == 0 { // root无覆盖
		res++
	}
	return res
}

var res int

func find(root *TreeNode) int {
	// 局部最优：让叶子节点的父节点安装摄像头，所用摄像头最少
	// 整体最优：全部摄像头数量所用最少
	// 每个节点状态：0无覆盖，1有摄像头，2有覆盖
	// 空节点的状态为有覆盖，这样就可以在叶子节点的父节点上放摄像头了
	if root == nil {
		return 2
	}
	left := find(root.Left)
	right := find(root.Right)

	//左右节点都有覆盖
	if left == 2 && right == 2 {
		return 0 //当前节点无需覆盖
	} else if left == 0 || right == 0 {
		// 左右节点中有一个未覆盖，需要一个摄像头
		res++
		return 1
	}
	// 左右节点中有一个摄像头，一个已覆盖，当前节点就覆盖
	return 2
}
