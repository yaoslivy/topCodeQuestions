package main

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	// this->1-> 2-> 3
	//获取链表中第 index 个节点的值
	if index < 0 {
		return -1
	}
	cur := this.Next
	for cur != nil && index != 0 {
		cur = cur.Next
		index--
	}
	if cur == nil {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 在链表的第一个元素之前添加一个值为 val 的节点
	// this->1-> 2-> 3
	node := &MyLinkedList{Val: val}
	node.Next = this.Next
	this.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	//将值为 val 的节点追加到链表的最后一个元素
	// this->1-> 2-> 3
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// this->1-> 2-> 3
	listLen := 0
	cur := this.Next
	for cur != nil {
		listLen++
		cur = cur.Next
	}
	// 如果 index 等于链表的长度，则该节点将附加到链表的末尾
	if index == listLen {
		this.AddAtTail(val)
		return
	}
	//如果 index 大于链表长度，则不会插入节点
	if index > listLen {
		return
	}
	//如果index小于0，则在头部插入节点。
	if index < 0 {
		this.AddAtHead(val)
	}
	//在链表中的第 index 个节点之前添加值为 val  的节点
	cur = this
	for cur.Next != nil && index != 0 {
		index--
		cur = cur.Next
	}
	node := &MyLinkedList{Val: val}
	node.Next = cur.Next
	cur.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 如果索引 index 有效，则删除链表中的第 index 个节点。
	// this->1-> 2-> 3
	if index < 0 {
		return
	}
	cur := this
	for cur.Next != nil && index != 0 {
		index--
		cur = cur.Next
	}
	if cur.Next == nil { //index超过了链表长度
		return
	}
	//删除最后一个节点 cur->delNode
	cur.Next = cur.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
