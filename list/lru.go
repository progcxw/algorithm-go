package list

// LRUCache 最近最少使用 (Least Recently Used) 缓存实现。
//
// 核心思想:
// LRU 缓存是一种常用的缓存淘汰策略，其核心思想是“如果一个数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小”。
// 为了实现这一策略，LRUCache 结合了哈希表 (map) 和双向链表 (DoubleLinkedList)。
// - 哈希表用于存储键值对，并提供 O(1) 的查找速度，通过键快速定位到链表中的节点。
// - 双向链表用于维护数据的访问顺序。链表头部是最近访问的数据，链表尾部是最久未访问的数据。
//
// 操作:
// - Get: 当访问一个数据时，如果数据存在，将其从链表中移除并重新插入到链表头部，表示其最近被访问过。如果不存在，返回 -1。
// - Put: 当插入一个新数据时，如果数据已存在，更新其值并将其移到链表头部。如果数据不存在：
//   - 如果缓存未满，直接插入到链表头部。
//   - 如果缓存已满，淘汰链表尾部（最久未使用）的数据，然后插入新数据到链表头部。
//
// 优点:
// - 查找、插入、删除操作的平均时间复杂度为 O(1)。
// - 能够有效地管理缓存空间，淘汰不常用的数据。
//
// 缺点:
// - 需要额外的空间来维护双向链表和哈希表。
// - 无法处理缓存击穿、缓存雪崩等问题（需要额外的机制）。
type LRUCache struct {
	cap   int               // 缓存的最大容量
	size  int               // 当前缓存中存储的元素数量
	cache map[int]*DLLNode  // 哈希表，键为数据的 key，值为链表中的节点指针
	list  *DoubleLinkedList // 双向链表，用于维护数据的访问顺序
}

// New 创建一个 LRUCache 实例。
//
// 时间复杂度: O(1)
// 空间复杂度: O(1)
//
// 参数:
//   - capacity: 缓存的最大容量。
//
// 返回值:
//   - *LRUCache: 新创建的 LRUCache 实例。
func New(capacity int) *LRUCache {
	return &LRUCache{
		cap:   capacity,
		size:  0, // 初始时缓存大小为 0
		cache: make(map[int]*DLLNode, capacity),
		list:  NewDoubleLinkedList(),
	}
}

// Get 从缓存中获取数据。
//
// 核心思想:
// 1. 通过 `key` 在哈希表中查找对应的节点。
// 2. 如果节点不存在，表示缓存中没有该数据，返回 -1。
// 3. 如果节点存在，表示数据被访问，需要更新其在链表中的位置，将其移动到链表头部（表示最近使用）。
// 4. 返回节点存储的值。
//
// 时间复杂度: O(1) - 哈希表查找和双向链表操作都是 O(1)。
//
// 参数:
//   - key: 要获取数据的键。
//
// 返回值:
//   - int: 对应键的值，如果键不存在则返回 -1。
func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1 // 缓存中不存在该 key
	}

	// 将节点从当前位置移除，并重新插入到链表头部
	c.list.Remove(node)
	c.list.InsertHead(node)

	return node.value
}

// Put 将数据放入缓存。
//
// 核心思想:
//  1. 检查 `key` 是否已存在于缓存中。
//  2. 如果 `key` 已存在：
//     a. 更新节点的值。
//     b. 将节点从链表中移除，并重新插入到链表头部（表示最近使用）。
//  3. 如果 `key` 不存在：
//     a. 创建一个新节点。
//     b. 检查缓存是否已满 (`c.size == c.cap`)。
//     c. 如果已满，淘汰链表尾部（最久未使用）的节点：从链表中移除，并从哈希表中删除。
//     d. 将新节点插入到链表头部。
//     e. 将新节点添加到哈希表中。
//     f. 增加缓存大小 `c.size`。
//
// 时间复杂度: O(1) - 哈希表查找、插入、删除和双向链表操作都是 O(1)。
//
// 参数:
//   - key: 要放入数据的键。
//   - val: 要放入数据的值。
func (c *LRUCache) Put(key int, val int) {
	// 检查 key 是否已存在
	node, ok := c.cache[key]
	if ok {
		// 如果存在，更新值并将其移到链表头部
		node.value = val
		c.list.Remove(node)
		c.list.InsertHead(node)
		return
	}

	// 如果 key 不存在，需要插入新节点
	// 检查缓存是否已满
	if c.size == c.cap {
		// 缓存已满，淘汰最久未使用的节点（链表尾部）
		tailNode := c.list.Tail()
		c.list.Remove(tailNode)
		delete(c.cache, tailNode.key)
		c.size-- // 淘汰一个，大小减一
	}

	// 创建新节点并插入到链表头部
	newNode := &DLLNode{
		key:   key,
		value: val,
	}
	c.list.InsertHead(newNode)
	c.cache[key] = newNode
	c.size++ // 插入一个，大小加一
}
