package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	//尽量在根节点放置中间值
	return buildBST(nums, 0, len(nums)-1)
}

//递归方式
func buildBST(nums []int, numsL, numsR int) *TreeNode {
	if numsL > numsR {
		return nil
	}
	//当前根节点值的下标, 每次选择中间值作为当前根的根节点值
	midIndex := (numsL + numsR) / 2
	root := TreeNode{Val: nums[midIndex]}
	root.Left = buildBST(nums, numsL, midIndex-1)
	root.Right = buildBST(nums, midIndex+1, numsR)

	return &root
}
