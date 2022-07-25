package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

//递归构造
func build(preorder []int, preL, preR int, inorder []int, inL, inR int) *TreeNode {
	if preL > preR {
		return nil
	}
	//preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	//用前序遍历序列中的根元素在中序遍历序列中划分左右子树
	root := TreeNode{Val: preorder[preL]}
	midIndex := inL
	for inorder[midIndex] != root.Val {
		midIndex++
	}
	//在中序序列中计算左子树个数，来划分前序遍历序列
	leftNum := midIndex - inL
	root.Left = build(preorder, preL+1, preL+leftNum, inorder, inL, midIndex-1)
	root.Right = build(preorder, preL+leftNum+1, preR, inorder, midIndex+1, inR)
	return &root
}
