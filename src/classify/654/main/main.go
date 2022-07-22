package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	head := build(nums, 0, len(nums)-1)
	return head
}

//返回当前序列最大值作为根节点的子树
func build(nums []int, numsL, numsR int) *TreeNode {
	if numsL > numsR {
		return nil
	}
	index := numsL
	for i := numsL; i <= numsR; i++ {
		if nums[index] < nums[i] {
			index = i
		}
	}
	root := TreeNode{Val: nums[index]}
	//找到最大值后，划分左右子树
	// nums = [3,2,1,6,0,5]
	root.Left = build(nums, numsL, index-1)
	root.Right = build(nums, index+1, numsR)
	return &root
}
