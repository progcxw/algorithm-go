package sort

// QuickSort 快速排序
func QuickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	pos := 0
	end := length - 1
	for i := 0; i < length; i++ {
		// 以数组的尾项为基准，将比其小的挪至数组前方并记录位置
		if nums[i] < nums[end] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}
	// 最后挪至的位置后一位即是分界点（前面比基准小，后面比基准大），把基准数swap到分界点
	nums[end], nums[pos] = nums[pos], nums[end]

	// 分而治之
	QuickSort(nums[:pos])
	QuickSort(nums[pos+1:])
}

// HeapSort 堆排
func HeapSort(nums []int) {
	// 从最后一个非叶子节点开始建立最大堆
	for i := len(nums)/2 - 1; i > 0; i-- {
		sink(nums, i)
	}

	// 最大堆形成，根节点为最大值，直接将其挪到最后
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		// 除了最后的节点外，再建一次最大堆，此时只用从首项开始sink
		sink(nums[:i], 0)
	}
}

func sink(nums []int, i int) {
	for {
		// 使用biggest记录当前子树中最大节点的index
		biggest := i
		length := len(nums)
		// 注意需要判断子节点index是否越界
		lChild := 2*i + 1
		rChild := 2*i + 2

		// 先找出i节点与左子节点中最大值的index
		if lChild < length && nums[i] < nums[lChild] {
			biggest = lChild
		}

		// 此时的biggest储存了i节点与左子节点中的最大值index， 再将该index节点与右子节点比较，得到三者中最大节点
		if rChild < length && nums[biggest] < nums[rChild] {
			biggest = rChild
		}

		if biggest == i {
			// 当前子树是最大堆，直接break
			break
		}

		// swap 最大节点与i节点
		nums[i], nums[biggest] = nums[biggest], nums[i]

		// 此时被换下来的节点的子节点可能比它大，往下sink
		i = biggest
	}
}
