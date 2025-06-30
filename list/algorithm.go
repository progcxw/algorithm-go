package list

// QuickSort 链表快速排序
//
// 核心思想:
// 采用经典的快速排序算法，但针对链表结构进行调整。
// 1. 选择一个基准节点 (pivot)，通常是链表的第一个节点。
// 2. 将链表分区 (partition)，分为两部分：一部分所有节点的值都小于基准节点，另一部分所有节点的值都大于或等于基准节点。
// 3. 递归地对这两个子链表进行快速排序。
// 4. 将排序后的 小于部分 -> 基准节点 -> 大于等于部分 连接起来。
//
// 时间复杂度: O(n log n) - 平均情况。O(n^2) - 最坏情况（当链表已有序或反向有序时）。
// 空间复杂度: O(log n) - 递归调用栈的深度。
//
// 参数:
//   - head: 链表的头节点指针。
//
// 返回值:
//   - *ListNode: 排序后链表的头节点指针。
func QuickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 创建两个虚拟头节点，用于构建小于和大于等于 pivot 的两个子链表
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less := lessHead
	greater := greaterHead

	// 选择第一个节点作为 pivot
	pivot := head
	curr := head.Next // 从 pivot 的下一个节点开始遍历

	// 分区：将链表节点分配到两个子链表中
	for curr != nil {
		if curr.Value < pivot.Value {
			less.Next = curr
			less = less.Next
		} else {
			greater.Next = curr
			greater = greater.Next
		}
		curr = curr.Next
	}

	// 断开链表，防止循环引用
	less.Next = nil
	greater.Next = nil

	// 递归排序子链表
	lessSorted := QuickSort(lessHead.Next)
	greaterSorted := QuickSort(greaterHead.Next)

	// 合并结果：将排序好的 less 部分、pivot 和 greater 部分连接起来
	if lessSorted != nil {
		// 找到 less 部分的尾部
		tail := lessSorted
		for tail.Next != nil {
			tail = tail.Next
		}
		// 连接 pivot
		tail.Next = pivot
		// 连接 greater 部分
		pivot.Next = greaterSorted
		return lessSorted
	}

	// 如果 less 部分为空，则 pivot 就是新的头
	pivot.Next = greaterSorted
	return pivot
}

// ReverseKGroup 每 k 个一组翻转链表
//
// 核心思想:
// 1. 使用一个虚拟头节点 `dummy` 来简化边界处理。
// 2. `pre` 指针指向每一组待翻转链表的前一个节点。
// 3. `tail` 指针用于确定每一组的末尾节点。每次循环，`tail` 从 `pre` 开始前进 k 步。
// 4. 如果 `tail` 成功找到了第 k 个节点（即剩余节点数 >= k），则进行翻转。
// 5. `reverse` 函数负责翻转从 `head` 到 `tail` 的这一小段链表。
// 6. 翻转后，需要重新连接链表：
//   - `pre.Next` 指向翻转后的新头节点（原来的 `tail`）。
//   - 原来的 `head` 变成了翻转后的尾节点，它的 `Next` 需要指向下一组的起始节点 `next`。
//
// 7. 更新 `pre` 和 `head`，为下一组翻转做准备。
//
// 时间复杂度: O(n) - 每个节点都被访问和翻转一次。
// 空间复杂度: O(1) - 只使用了常数个额外指针。
//
// 参数:
//   - head: 链表的头节点指针。
//   - k: 每组的节点数。
//
// 返回值:
//   - *ListNode: 翻转后链表的头节点指针。
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre
		// 检查剩余部分是否足够 k 个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 不足 k 个，无需翻转，直接返回
				return dummy.Next
			}
		}

		// 保存下一组的头节点
		nextGroupHead := tail.Next
		// 翻转当前组
		newHead, newTail := reverse(head, tail)

		// 将翻转后的组重新连接到主链表
		pre.Next = newHead
		newTail.Next = nextGroupHead

		// 更新 pre 和 head，为下一组做准备
		pre = newTail
		head = nextGroupHead
	}

	return dummy.Next
}

// reverse 翻转从 head 到 tail 的链表段
//
// 这是一个标准的链表翻转操作，但它在 `tail.Next` 处停止。
//
// 参数:
//   - head: 需要翻转部分的头节点。
//   - tail: 需要翻转部分的尾节点。
//
// 返回值:
//   - *ListNode: 翻转后的新头节点 (即原 tail)。
//   - *ListNode: 翻转后的新尾节点 (即原 head)。
func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode = nil
	curr := head
	stopNode := tail.Next // 翻转操作的终止条件

	for curr != stopNode {
		next := curr.Next // 保存下一个节点
		curr.Next = prev  // 反转指针
		prev = curr       // prev 前进
		curr = next       // curr 前进
	}

	// 翻转后，prev 指向新的头节点（原 tail），head 指向新的尾节点
	return tail, head
}
