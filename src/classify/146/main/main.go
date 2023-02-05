package main

// 使用双向链表+哈希表解决
// 双向链表即按照被使用顺序存储k-v，靠近头部的k-v是最近使用的
// 哈希表存储k-node映射关系，使得能够快速定位k所在的节点位置

type LRUCache struct {
	head, tail *DLinkNode
	nodeMap    map[int]*DLinkNode // key-node
	capacity   int
	size       int
}

type DLinkNode struct {
	key, value int
	pre, next  *DLinkNode
}

// 使用一个双向链表保存kv，最近使用过的放置在链表头节点
// 使用一个map记录现有的k作为value

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     &DLinkNode{},
		tail:     &DLinkNode{},
		nodeMap:  map[int]*DLinkNode{},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	// <-head<->tail->
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.deleteNode(node)
		this.addToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// <-head<->[1, 2]<->tail->
	// 节点key存在，更新value，并将节点移动到链表首部
	if node, ok := this.nodeMap[key]; ok {
		node.value = value
		this.deleteNode(node)
		this.addToHead(node)
	} else {
		// 新节点
		node := &DLinkNode{key: key, value: value}
		this.addToHead(node)
		this.nodeMap[key] = node
		this.size++
		// 容量不够则删除最后一个节点
		if this.size > this.capacity {
			tailNode := this.tail.pre
			this.deleteNode(tailNode)
			delete(this.nodeMap, tailNode.key)
			this.size--
		}
	}
}

// 双向链表中删除节点
func (this *LRUCache) deleteNode(node *DLinkNode) {
	// <-head<-[1, 2]<->tail->
	// <-head<->tail
	node.pre.next = node.next
	node.next.pre = node.pre
}

// 在双向链表首部添加节点
func (this *LRUCache) addToHead(node *DLinkNode) {
	// <-head<->[1, 2]<->tail
	node.next = this.head.next
	this.head.next = node
	node.pre = this.head
	node.next.pre = node
}
