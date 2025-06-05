package list

// DLLNode (Double Linked List Node)双向链表节点
// 使用list保存所属链表指针，方便Prev()和Next()方法判断是否已到链表边界
type DLLNode struct {
	key   int
	value int
	list  *DoubleLinkedList
	prev  *DLLNode
	next  *DLLNode
}

// Prev 返回前一个节点，如果当前节点是未插入链表的节点或头节点，则返回nil
func (n *DLLNode) Prev() *DLLNode {
	p := n.prev
	if n.list != nil && p != &n.list.root {
		return p
	}

	return nil
}

// Next 返回后一个节点，如果当前节点是未插入链表的节点或尾节点，则返回nil
func (n *DLLNode) Next() *DLLNode {
	next := n.next
	if n.list != nil && next != &n.list.root {
		return next
	}

	return nil
}

// ListNode 链表节点结构
// 包含指向下一个节点的指针和节点值
type ListNode struct {
	Next  *ListNode // 指向下一个节点的指针
	Value int       // 节点存储的整数值
}
