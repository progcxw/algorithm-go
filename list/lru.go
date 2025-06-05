package list

// LRUCache 最近最少使用老化机制的缓存，使用双向链表记录使用顺序，map实现缓存
type LRUCache struct {
	cap   int
	size  int
	cache map[int]*DLLNode
	list  *DoubleLinkedList
}

// New 创建一个LRUCache实例
func New(capacity int) *LRUCache {
	return &LRUCache{
		cap:   capacity,
		cache: make(map[int]*DLLNode, capacity),
		list:  NewDoubleLinkedList(),
	}
}

// Get 从缓存中获取数据，如果不存在则返回-1
func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.list.Remove(node)
	c.list.InsertHead(node)
	return node.value
}

// Put 将数据放入缓存，如果缓存已满则淘汰最久未使用的数据
func (c *LRUCache) Put(key int, val int) {
	node, ok := c.cache[key]
	if ok {
		c.list.Remove(node)
		node.value = val
		c.list.InsertHead(node)
		return
	}

	if c.size == c.cap {
		tail := c.list.Tail()
		c.list.Remove(tail)
		delete(c.cache, tail.key)
	}

	node = &DLLNode{
		key:   key,
		value: val,
	}
	c.list.InsertHead(node)
	c.cache[key] = node
}
