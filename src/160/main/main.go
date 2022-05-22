package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	//pA先遍历A序列，再遍历B序列，pB先遍历B序列，再遍历A序列，若有相同部分，则一定会在相同部分开头相遇，没有相同部分，则会一同遍历到nil
	//证明：A: a+c   B: b+c    pA: a+c+b+c   pB: b+c+a+c   => a+c+b == b+c+a
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	mm := make(map[*ListNode]bool, 0)
	//使用哈希表记录所有在A中出现节点，然后遍历B的同时判断在哈希表中第一个出现的节点
	for pA := headA; pA != nil; pA = pA.Next {
		mm[pA] = true
	}
	for pB := headB; pB != nil; pB = pB.Next {
		if _, ok := mm[pB]; ok {
			return pB
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	S1 := make([]*ListNode, 0)
	S2 := make([]*ListNode, 0)
	pA := headA
	pB := headB
	for pA != nil {
		S1 = append(S1, pA)
		pA = pA.Next
	}
	for pB != nil {
		S2 = append(S2, pB)
		pB = pB.Next
	}
	var common *ListNode
	for len(S1) != 0 && len(S2) != 0 {
		if S1[len(S1)-1] == S2[len(S2)-1] {
			common = S1[len(S1)-1]
			S1 = S1[:len(S1)-1]
			S2 = S2[:len(S2)-1]
		} else {
			break
		}
	}
	if common != nil {
		return common
	}
	return nil
}
