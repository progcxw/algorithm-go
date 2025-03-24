package list

// DoubleLinkedListNode 双向链表节点
type DoubleLinkedListNode struct {
	Key   int
	Value int
	Prev  *DoubleLinkedListNode
	Next  *DoubleLinkedListNode
}

// ListNode 链表节点结构
// 包含指向下一个节点的指针和节点值
type ListNode struct {
	Next  *ListNode // 指向下一个节点的指针
	Value int       // 节点存储的整数值
}
