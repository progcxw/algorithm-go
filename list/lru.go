package list

// LRUCache LRU缓存, Last Recently Used, 最近使用的排到前面，最久未使用的末位淘汰
type LRUCache struct {
	Capacity int
	CacheMap map[int]*DoubleLinkedListNode
	Head     *DoubleLinkedListNode
	Tail     *DoubleLinkedListNode
}

// Constructor 构造LRU缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		CacheMap: make(map[int]*DoubleLinkedListNode, capacity),
		Head:     nil,
		Tail:     nil,
	}
}

// moveToHead 将节点移动到头部
func (lru *LRUCache) moveToHead(node *DoubleLinkedListNode) {
	// 如果节点已经在头部，直接返回
	if node == lru.Head {
		return
	}

	// 如果是第一个节点
	if lru.Head == nil {
		lru.Head = node
		lru.Tail = node
		return
	}

	// 如果节点是尾部，更新尾部
	if node == lru.Tail {
		lru.Tail = node.Prev
		if lru.Tail != nil {
			lru.Tail.Next = nil
		}
	}

	// 从当前位置移除节点
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// 将节点放到头部
	node.Prev = nil
	node.Next = lru.Head
	lru.Head.Prev = node
	lru.Head = node
}

// removeTail 删除尾部节点
func (lru *LRUCache) removeTail() {
	if lru.Tail == nil {
		return
	}

	delete(lru.CacheMap, lru.Tail.Key)
	lru.Tail = lru.Tail.Prev
	if lru.Tail != nil {
		lru.Tail.Next = nil
	}
}

// Get 获取缓存
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.CacheMap[key]; ok {
		lru.moveToHead(node)
		return node.Value
	}

	// 按题目要求，找不到返回-1
	return -1
}

// Put 添加缓存
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.CacheMap[key]; ok {
		node.Value = value
		lru.moveToHead(node)
		return
	}

	// 创建新节点, 添加到缓存
	newNode := &DoubleLinkedListNode{
		Key:   key,
		Value: value,
	}
	lru.CacheMap[key] = newNode

	// 如果是第一个节点, 更新头部和尾部; 否则, 添加到头部
	if lru.Head == nil {
		lru.Head = newNode
		lru.Tail = newNode
	} else {
		newNode.Next = lru.Head
		lru.Head.Prev = newNode
		lru.Head = newNode
	}

	// 如果缓存超过容量，删除尾部节点
	if len(lru.CacheMap) > lru.Capacity {
		lru.removeTail()
	}
}
