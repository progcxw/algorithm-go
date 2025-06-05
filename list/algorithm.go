package list

// QuickSort 链表快速排序
// 对链表进行原地快速排序，不需要额外的存储空间
// 参数:
//   - head: 链表头节点指针
func QuickSort(head *ListNode) {
	recursiveQuickSort(head, nil)
}

// recursiveQuickSort 链表快速排序的递归实现
// 参数:
//   - head: 当前需要排序的链表头节点
//   - tail: 排序范围的尾部边界（不包含在排序范围内）
func recursiveQuickSort(head *ListNode, tail *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 快排实现
	dummy := &ListNode{
		Next:  head,
		Value: 0, // 创建一个虚拟头节点，简化边界情况处理
	}
	// 选取首项作为基准点
	pivot := head
	pre := dummy
	for head != tail {
		if head.Value < pivot.Value {
			// 取出当前节点
			node := head
			pre.Next = head.Next

			// 插入到链表首部（小于基准值的节点都放到链表前面）
			node.Next = dummy.Next
			dummy.Next = node

			// 移动到下一个节点
			head = pre.Next
		} else {
			// 当前节点值大于等于基准值，保持位置不变
			pre = head
			head = head.Next
		}
	}

	// 分治递归处理
	recursiveQuickSort(dummy.Next, pivot) // 处理小于基准值的部分
	recursiveQuickSort(pivot.Next, nil)   // 处理大于等于基准值的部分
}

// ReverseKGroup 每k个一组翻转链表
// 将链表每k个节点为一组进行翻转，如果最后一组不足k个节点则保持原有顺序
// 参数:
//   - head: 链表头节点指针
//   - k: 每组的节点数
//
// 返回值:
//   - *Node: 翻转后的链表头节点指针
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next:  head,
		Value: 0, // 创建虚拟头节点，简化边界情况处理
	}
	pre := dummy // pre指向当前处理组的前一个节点

	for head != nil {
		tail := pre

		// 把tail移动到第k个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 如果剩余节点不足k个，直接返回
				return dummy.Next
			}
		}

		next := tail.Next                // 保存下一组的起始节点
		head, tail = reverse(head, tail) // 翻转当前组
		pre.Next = head                  // 将翻转后的组连接到前面的链表
		tail.Next = next                 // 将翻转后的组连接到后面的链表
		pre = tail                       // 更新pre为当前组的最后一个节点
		head = tail.Next                 // 更新head为下一组的起始节点
	}
	return dummy.Next // 返回翻转后的链表头节点
}

// reverse 翻转链表的指定部分
// 将链表从head到tail的部分进行翻转
// 参数:
//   - head: 需要翻转部分的起始节点
//   - tail: 需要翻转部分的结束节点
//
// 返回值:
//   - *Node: 翻转后的起始节点（原tail）
//   - *Node: 翻转后的结束节点（原head）
func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next // 保存tail的下一个节点作为新的尾部
	p := head         // 当前处理的节点
	for prev != tail {
		next := p.Next // 保存下一个要处理的节点
		p.Next = prev  // 当前节点指向前一个节点（反转指针）
		prev = p       // 更新prev为当前节点
		p = next       // 移动到下一个节点
	}
	return tail, head // 返回新的头节点（原tail）和新的尾节点（原head）
}
