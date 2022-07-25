package main

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{} //返回一个虚拟头节点, Val=0, Next=nil
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 { //不在0-index中
		return -1
	}
	// this->0->1->2
	cur := this.Next
	for cur != nil && index != 0 {
		cur = cur.Next
		index--
	}
	if cur == nil { //超出了已存在的元素个数
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// this->0->1->2
	node := MyLinkedList{Val: val}
	node.Next = this.Next //没有元素则指向nil
	this.Next = &node
}

func (this *MyLinkedList) AddAtTail(val int) {
	// this->0->1->2
	tail := MyLinkedList{Val: val}
	cur := this
	for cur.Next != nil { //找到最后一个元素
		cur = cur.Next
	}
	cur.Next = &tail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	//如果index小于0，则在头部插入节点。
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	length := 0
	cur := this.Next
	for cur != nil {
		length++
		cur = cur.Next
	}
	//如果 index 等于链表的长度，则该节点将附加到链表的末尾
	if index == length {
		this.AddAtTail(val)
		return
	}
	//如果 index 大于链表长度，则不会插入节点。
	if index > length {
		return
	}
	//在链表中的第 index 个节点之前添加值为 val  的节点
	cur = this
	// this->0->1->2
	for index != 0 {
		cur = cur.Next
		index--
	}
	node := MyLinkedList{Val: val}
	node.Next = cur.Next
	cur.Next = &node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	// this->0->1->2
	cur := this
	for cur.Next != nil && index != 0 {
		index--
		cur = cur.Next
	}
	if cur.Next == nil {
		return
	}
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
