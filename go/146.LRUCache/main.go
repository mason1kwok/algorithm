package main

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type DoubleList struct {
	Head *Node
	Tail *Node
	Size int
}

// 新增节点到双向链表尾部
func (d *DoubleList) addLast(n *Node) {
	n.Prev = d.Tail.Prev
	n.Next = d.Tail
	d.Tail.Prev.Next = n
	d.Tail.Prev = n
	d.Size++
}

// 删除给定节点
func (d *DoubleList) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	d.Size--
}

// 删除第一节点并返回
func (d *DoubleList) removeFirst() *Node {
	if d.Tail.Next == d.Tail {
		return nil
	}
	first := d.Head.Next
	d.remove(first)
	return first
}

// 返回链表长度
func (d *DoubleList) getSize() int {
	return d.Size
}

type LRUCache struct {
	mapping  map[int]*Node
	cache    *DoubleList
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &Node{Key: 0, Val: 0}
	tail := &Node{Key: 0, Val: 0}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		cache: &DoubleList{
			Head: head,
			Tail: tail,
			Size: 0,
		},
		mapping:  map[int]*Node{},
		capacity: capacity,
	}
}

// 将某个key作为最近使用的元素
func (this *LRUCache) makeRecently(key int) {
	node, ok := this.mapping[key]
	if ok {
		this.cache.remove(node)
		this.cache.addLast(node)
	}
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key, val int) {
	node := &Node{Key: key, Val: val}
	// 添加到链表最后
	this.cache.addLast(node)
	// 增加key
	this.mapping[key] = node
}

// 删除key
func (this *LRUCache) deleteKey(key int) {
	node, ok := this.mapping[key]
	if ok {
		// 从链表中删除
		this.cache.remove(node)
		// 从hashMap删除
		delete(this.mapping, key)
	}
}

// 删除最久未使用的元素
func (this *LRUCache) removeLeastRecently() {
	node := this.cache.removeFirst()
	if node != nil {
		delete(this.mapping, node.Key)
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mapping[key]; ok {
		this.makeRecently(key)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.mapping[key]; ok {
		this.deleteKey(key)
	}

	// 容量满了以后，删除最久未使用的元素
	if this.cache.Size == this.capacity {
		this.removeLeastRecently()
	}

	this.addRecently(key, value)
}
