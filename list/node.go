package list

// DLLNode (Double Linked List Node) 双向链表节点结构。
// 每个节点包含一个键 (key) 和一个值 (value)，以及指向前一个和后一个节点的指针。
// `list` 字段用于指向该节点所属的双向链表，方便在节点层面进行操作时判断其归属。
//
// 字段:
//   - key: 节点的键，通常用于在哈希表中快速查找节点。
//   - value: 节点存储的实际值。
//   - list: 指向该节点所属的 `DoubleLinkedList` 实例。如果节点未插入任何链表，则为 nil。
//   - prev: 指向链表中前一个节点的指针。如果当前节点是头节点，则指向链表的哨兵节点。
//   - next: 指向链表中后一个节点的指针。如果当前节点是尾节点，则指向链表的哨兵节点。
type DLLNode struct {
	key   int
	value int
	list  *DoubleLinkedList
	prev  *DLLNode
	next  *DLLNode
}

// Prev 返回当前节点在链表中的前一个实际节点。
//
// 核心思想:
// 1. 检查节点是否已插入链表 (`n.list != nil`)。
// 2. 检查 `n.prev` 是否是链表的哨兵节点 (`&n.list.root`)。
//   - 如果 `n.prev` 是哨兵节点，说明当前节点是链表的第一个实际节点，其前一个节点不是实际数据节点。
//   - 否则，`n.prev` 就是前一个实际节点。
//
// 时间复杂度: O(1)
//
// 返回值:
//   - *DLLNode: 前一个实际节点，如果当前节点是链表头或未插入链表，则返回 nil。
func (n *DLLNode) Prev() *DLLNode {
	p := n.prev
	// 只有当节点属于某个链表，并且其前一个节点不是该链表的哨兵节点时，才返回前一个节点。
	if n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

// Next 返回当前节点在链表中的后一个实际节点。
//
// 核心思想:
// 1. 检查节点是否已插入链表 (`n.list != nil`)。
// 2. 检查 `n.next` 是否是链表的哨兵节点 (`&n.list.root`)。
//   - 如果 `n.next` 是哨兵节点，说明当前节点是链表的最后一个实际节点，其后一个节点不是实际数据节点。
//   - 否则，`n.next` 就是后一个实际节点。
//
// 时间复杂度: O(1)
//
// 返回值:
//   - *DLLNode: 后一个实际节点，如果当前节点是链表尾或未插入链表，则返回 nil。
func (n *DLLNode) Next() *DLLNode {
	next := n.next
	// 只有当节点属于某个链表，并且其后一个节点不是该链表的哨兵节点时，才返回后一个节点。
	if n.list != nil && next != &n.list.root {
		return next
	}
	return nil
}

// ListNode 单向链表节点结构。
// 包含指向下一个节点的指针和节点值。
//
// 字段:
//   - Next: 指向链表中下一个节点的指针。如果当前节点是尾节点，则为 nil。
//   - Value: 节点存储的整数值。
type ListNode struct {
	Next  *ListNode // 指向下一个节点的指针
	Value int       // 节点存储的整数值
}
