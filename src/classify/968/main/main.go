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
	// 后序遍历判断每个节点状态然后再考虑是否放摄像头
	// 放摄像头的策略是：叶子结点不放置，尽量放置在中间节点
	// 定义每个节点的状态为 0 未放置摄像头且无摄像头覆盖， 1 未放置摄像头但有摄像头覆盖， 2 该节点放置摄像头
	if root == nil {
		return 1 //空节点的状态为无摄像头，但是有摄像头覆盖，如此，叶子节点就无需放置摄像头
	}
	// 判断左右子节点的状态，然后决定当前节点的状态
	left := find(root.Left)
	right := find(root.Right)

	if left == 1 && right == 1 { //左右子节点都有覆盖
		return 0
	} else if left == 0 || right == 0 { //只有左右子节点存在一个无覆盖，当前根节点都需要放置一个摄像头
		res++
		return 2
	}
	return 1
}
