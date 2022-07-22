package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	//递归找根节点
	head := build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
	return head
}

func build(inorder []int, inL, inR int, postorder []int, postL, postR int) *TreeNode {
	if inL > inR {
		return nil
	}
	// inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	// post末尾就是当前子树的根节点，用根节点在inorder中划分左右子树
	root := TreeNode{Val: postorder[postR]}
	mid := inL
	for inorder[mid] != root.Val {
		mid++
	}
	//当前根左子树个数
	leftNum := mid - inL
	root.Left = build(inorder, inL, mid-1, postorder, postL, postL+leftNum-1)
	root.Right = build(inorder, mid+1, inR, postorder, postL+leftNum, postR-1)
	return &root
}
