package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	head := build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
	return head
}

func build(preorder []int, preL, preR int, inorder []int, inL, inR int) *TreeNode {
	if preL > preR {
		return nil
	}
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// 先序遍历序列的第一个值即为当前根节点
	root := TreeNode{Val: preorder[preL]}
	//在中序序列中找到当前根节点位置，然后划分左右子树
	mid := inL
	for inorder[mid] != root.Val {
		mid++
	}
	//左子树的个数
	leftNum := mid - inL
	root.Left = build(preorder, preL+1, preL+leftNum, inorder, inL, inL+leftNum-1)
	root.Right = build(preorder, preL+leftNum+1, preR, inorder, inL+leftNum+1, inR)
	return &root
}
