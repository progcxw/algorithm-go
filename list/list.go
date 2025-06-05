package list

// DoubleLinkedList 双向链表
// 使用root节点作为哨兵，简化边界情况处理
type DoubleLinkedList struct {
	root DLLNode
	size int
}

// NewDoubleLinkedList 创建一个空的双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	l := new(DoubleLinkedList)
	l.root.prev = &l.root
	l.root.next = &l.root
	return l
}

func (l *DoubleLinkedList) Len() int {
	return l.size
}

// Head 返回链表头节点，如果链表为空则返回nil
func (l *DoubleLinkedList) Head() *DLLNode {
	if l.size != 0 && l.root.next != &l.root {
		return l.root.next
	}
	return nil
}

// Tail 返回链表尾节点，如果链表为空则返回nil
func (l *DoubleLinkedList) Tail() *DLLNode {
	if l.size != 0 && l.root.prev != &l.root {
		return l.root.prev
	}
	return nil
}

func (l *DoubleLinkedList) InsertHead(node *DLLNode) {
	head := l.Head()
	head.prev = node
	node.next = head
	node.prev = &l.root
	node.list = l
	l.root.next = node
	l.size++
}

func (l *DoubleLinkedList) Remove(node *DLLNode) {
	if node.list != l {
		return
	}

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
	node.prev, node.next = nil, nil
	node.list = nil
	l.size--
}
