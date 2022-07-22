package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//  return lowestCommonAncestorRecursion(root, p, q)

	return lowestCommonAncestorTraverse(root, p, q)
}

//迭代方式
func lowestCommonAncestorTraverse(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	S := make([]*TreeNode, 0)
	var Sp []*TreeNode
	var Sq []*TreeNode
	cur := root
	var pre *TreeNode //记录上一个访问节点
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
				continue
			}
			if cur == p { //找到节点p,则当前栈中所有节点都是祖先节点
				Sp = make([]*TreeNode, len(S))
				copy(Sp, S)
			}
			if cur == q {
				Sq = make([]*TreeNode, len(S))
				copy(Sq, S)
			}
			if len(Sp) != 0 && len(Sq) != 0 {
				break
			}

			S = S[:len(S)-1]
			pre = cur
			cur = nil
		}
	}
	//在两个祖先节点栈中找到最近的公共祖先节点
	for len(Sp) != len(Sq) {
		if len(Sp) > len(Sq) {
			Sp = Sp[:len(Sp)-1]
		} else {
			Sq = Sq[:len(Sq)-1]
		}
	}
	//长度相同情况下，栈顶元素相同则是最近公共祖先节点
	for Sp[len(Sp)-1] != Sq[len(Sq)-1] {
		Sp = Sp[:len(Sp)-1]
		Sq = Sq[:len(Sq)-1]
	}
	return Sp[len(Sp)-1]
}

//递归方式
func lowestCommonAncestorRecursion(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorRecursion(root.Left, p, q)
	right := lowestCommonAncestorRecursion(root.Right, p, q)

	if left != nil && right != nil { //说明p,q分布在当前节点两端
		return root //当前节点就是最近公共祖先节点
	}
	if left != nil { //说明p,q两者之一存在于当前节点的左子树中
		return left //该左子节点一定是p,q中的一个祖先节点
	}
	return right
}
