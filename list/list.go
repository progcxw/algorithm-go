package list

// DoubleLinkedList 双向链表
// 使用 `root` 节点作为哨兵（sentinel），简化边界情况处理。
// 哨兵节点不存储实际数据，它的 `next` 指向链表的第一个实际节点，`prev` 指向链表的最后一个实际节点。
// 当链表为空时，`root.next` 和 `root.prev` 都指向 `root` 自身。
type DoubleLinkedList struct {
	root DLLNode // 哨兵节点
	size int     // 链表中实际节点的数量
}

// NewDoubleLinkedList 创建一个空的双向链表。
//
// 时间复杂度: O(1)
// 空间复杂度: O(1)
//
// 返回值:
//   - *DoubleLinkedList: 新创建的空双向链表实例。
func NewDoubleLinkedList() *DoubleLinkedList {
	l := new(DoubleLinkedList)
	// 初始化哨兵节点，使其指向自身，表示链表为空。
	l.root.prev = &l.root
	l.root.next = &l.root
	return l
}

// Len 返回链表中实际节点的数量。
//
// 时间复杂度: O(1)
//
// 返回值:
//   - int: 链表的长度。
func (l *DoubleLinkedList) Len() int {
	return l.size
}

// Head 返回链表的第一个实际节点。
//
// 时间复杂度: O(1)
//
// 返回值:
//   - *DLLNode: 链表的头节点，如果链表为空则返回 nil。
func (l *DoubleLinkedList) Head() *DLLNode {
	// 如果链表不为空且 root.next 不是哨兵节点本身，则返回 root.next。
	if l.size != 0 && l.root.next != &l.root {
		return l.root.next
	}
	return nil
}

// Tail 返回链表的最后一个实际节点。
//
// 时间复杂度: O(1)
//
// 返回值:
//   - *DLLNode: 链表的尾节点，如果链表为空则返回 nil。
func (l *DoubleLinkedList) Tail() *DLLNode {
	// 如果链表不为空且 root.prev 不是哨兵节点本身，则返回 root.prev。
	if l.size != 0 && l.root.prev != &l.root {
		return l.root.prev
	}
	return nil
}

// InsertHead 在链表头部插入一个节点。
//
// 核心思想:
// 1. 将新节点的 `next` 指向当前链表的第一个实际节点（即 `l.root.next`）。
// 2. 将新节点的 `prev` 指向哨兵节点 `l.root`。
// 3. 更新原第一个实际节点的 `prev` 指向新节点。
// 4. 更新哨兵节点 `l.root` 的 `next` 指向新节点。
// 5. 设置新节点的 `list` 属性为当前链表。
// 6. 增加链表大小 `size`。
//
// 时间复杂度: O(1)
//
// 参数:
//   - node: 要插入的节点指针。
func (l *DoubleLinkedList) InsertHead(node *DLLNode) {
	// 将新节点插入到 root 和 root.next 之间
	next := l.root.next

	node.prev = &l.root
	node.next = next
	node.list = l

	next.prev = node
	l.root.next = node

	l.size++
}

// Remove 从链表中移除一个节点。
//
// 核心思想:
// 1. 检查要移除的节点是否属于当前链表，如果不是则直接返回。
// 2. 将要移除节点的前一个节点的 `next` 指向要移除节点的后一个节点。
// 3. 将要移除节点的后一个节点的 `prev` 指向要移除节点的前一个节点。
// 4. 清空被移除节点的 `prev`, `next` 和 `list` 属性，使其不再引用链表。
// 5. 减小链表大小 `size`。
//
// 时间复杂度: O(1)
//
// 参数:
//   - node: 要移除的节点指针。
func (l *DoubleLinkedList) Remove(node *DLLNode) {
	// 只有当节点属于当前链表时才执行移除操作
	if node.list != l {
		return
	}

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	// 清除被移除节点的引用，防止悬挂指针
	node.prev = nil
	node.next = nil
	node.list = nil

	l.size--
}
