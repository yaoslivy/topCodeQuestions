package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode //hash table
	head, tail *DLinkedNode         //头节点的next为第一个节点，尾节点的prev为最后一个节点
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok { // Get操作时，key不存在时返回-1
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node) //最近使用过，移动到头部
	return node.value
}

//在含头尾指针的双向链表中从当前节点移动到头节点，即删除当前节点，再添加当前节点到头节点尾部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node // 哈希表中每一个key都对应着一个双向节点
		this.addToHead(node)   //添加到双向链表头部，尾部为最近最少使用的
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail() //移除最后一个
			delete(this.cache, removed.key)
			this.size--
		}
	} else { //如果对应的key存在，则通过哈希表直接找到节点，重写节点的值，再移动到双向链表头部，尾部对应着最近最少使用的
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}
