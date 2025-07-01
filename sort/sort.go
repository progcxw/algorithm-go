package sort

// BubbleSort 冒泡排序
//
// 核心思想:
// 通过重复遍历待排序的列表，比较相邻的两个元素，如果它们的顺序错误就把它们交换过来。
// 遍历列表的工作是重复地进行直到没有再需要交换，也就是说该列表已经排序完成。
// 每一轮遍历都会将当前未排序部分的最大（或最小）元素“冒泡”到其最终位置。
//
// 时间复杂度: O(n²) - 在所有情况下（最好、平均、最坏）都需要进行 n² 次比较。
// 空间复杂度: O(1) - 只使用了常数个额外空间。
//
// 参数:
//   - nums: 需要排序的整数切片。
func BubbleSort(nums []int) {
	n := len(nums)
	// 外层循环控制排序的轮数，每轮确定一个元素的位置。
	for i := 0; i < n-1; i++ {
		// 内层循环进行相邻元素的比较和交换。
		// 每次内层循环结束后，最大的元素会被“冒泡”到数组的末尾。
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// SelectionSort 选择排序
//
// 核心思想:
// 1. 在未排序序列中找到最小（或最大）元素，存放到排序序列的起始位置。
// 2. 再从剩余未排序元素中继续寻找最小（或最大）元素，然后放到已排序序列的末尾。
// 3. 重复第二步，直到所有元素均排序完毕。
//
// 时间复杂度: O(n²) - 无论输入数据如何，都需要进行 n² 次比较。
// 空间复杂度: O(1) - 只使用了常数个额外空间。
//
// 参数:
//   - nums: 需要排序的整数切片。
func SelectionSort(nums []int) {
	n := len(nums)
	// 外层循环控制已排序部分的边界。
	for i := 0; i < n-1; i++ {
		minIndex := i // 记录当前未排序部分的最小元素的索引。
		// 内层循环在未排序部分查找最小元素。
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 将找到的最小元素与当前未排序部分的第一个元素交换。
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// InsertionSort 插入排序
//
// 核心思想:
// 1. 将第一个元素视为已排序部分，其余元素视为未排序部分。
// 2. 遍历未排序部分，取出每个元素，在已排序部分中从后向前扫描。
// 3. 找到该元素在已排序部分中的正确位置，并将其插入。
//   - 如果当前元素小于已排序部分的元素，则将已排序部分的元素向后移动一位。
//   - 直到找到一个小于或等于当前元素的元素，或者到达已排序部分的开头。
//
// 时间复杂度: O(n²) - 最坏情况（逆序）。O(n) - 最好情况（已排序）。
// 空间复杂度: O(1) - 只使用了常数个额外空间。
//
// 参数:
//   - nums: 需要排序的整数切片。
func InsertionSort(nums []int) {
	n := len(nums)
	// 从第二个元素开始遍历，将其插入到已排序部分。
	for i := 1; i < n; i++ {
		current := nums[i] // 待插入的元素。
		j := i - 1         // 已排序部分的最后一个元素的索引。
		// 在已排序部分中从后向前查找插入位置。
		for j >= 0 && nums[j] > current {
			nums[j+1] = nums[j] // 元素后移。
			j--
		}
		nums[j+1] = current // 插入元素。
	}
}

// ShellSort 希尔排序
//
// 核心思想:
// 希尔排序是插入排序的一种更高效的改进版本。它通过比较相距一定间隔的元素来工作。
// 1. 选择一个增量序列 (gap sequence)，例如 Knuth 序列 (1, 4, 13, 40, ...)。
// 2. 对每个增量 `gap`，将数组分成 `gap` 个子序列，每个子序列由相距 `gap` 的元素组成。
// 3. 对每个子序列进行插入排序。
// 4. 逐渐减小 `gap`，重复步骤 2 和 3，直到 `gap` 为 1。当 `gap` 为 1 时，整个数组被视为一个子序列，进行最后一次插入排序，此时数组基本有序，插入排序效率很高。
//
// 时间复杂度: 平均 O(n^1.3) - 具体取决于增量序列。最坏 O(n²)。
// 空间复杂度: O(1) - 只使用了常数个额外空间。
//
// 参数:
//   - nums: 需要排序的整数切片。
func ShellSort(nums []int) {
	n := len(nums)
	// 选择增量序列，这里使用 n/2, n/4, ..., 1
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个子序列进行插入排序
		for i := gap; i < n; i++ {
			current := nums[i]
			j := i - gap
			// 在当前子序列中进行插入排序
			for j >= 0 && nums[j] > current {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = current
		}
	}
}

// MergeSort 归并排序
//
// 核心思想:
// 归并排序是分治法的一个典型应用。
// 1. 分解 (Divide): 将待排序的序列分解成两个子序列。
// 2. 解决 (Conquer): 递归地对这两个子序列进行排序。
// 3. 合并 (Combine): 将两个已排序的子序列合并成一个完整的有序序列。
// 合并操作是归并排序的关键，它通过比较两个子序列的元素，逐步构建出最终的有序序列。
//
// 时间复杂度: O(n log n) - 无论输入数据如何，性能都稳定。
// 空间复杂度: O(n) - 需要额外的空间来存储合并过程中的临时数组。
//
// 参数:
//   - nums: 需要排序的整数切片。
//
// 返回值:
//   - []int: 排序后的新切片。
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])  // 递归排序左半部分
	right := MergeSort(nums[mid:]) // 递归排序右半部分

	// 合并两个有序切片
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// 将剩余元素添加到结果切片
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// QuickSort 快速排序
//
// 核心思想:
// 快速排序是一种高效的、基于比较的排序算法，采用分治法。
// 1. 选择一个“基准”元素 (pivot)。
// 2. 分区 (Partition): 重新排列数组，使得所有小于基准值的元素都移到基准的左边，所有大于基准值的元素都移到基准的右边。在这个分区结束之后，该基准就处于其最终的正确位置。
// 3. 递归 (Recursion): 递归地对基准左边和右边的两个子序列重复上述步骤，直到所有元素都排好序。
//
// 时间复杂度: 平均 O(n log n) - 每次分区都能将问题规模大致减半。最坏 O(n²) - 当选择的基准导致不平衡分区时（例如，每次都选择最大或最小元素）。
// 空间复杂度: O(log n) - 递归调用栈的深度，平均情况。最坏 O(n)。
//
// 参数:
//   - nums: 需要排序的整数切片。
func QuickSort(nums []int) {
	quickSortRecursive(nums, 0, len(nums)-1)
}

// quickSortRecursive 快速排序的递归辅助函数。
//
// 参数:
//   - nums: 整数切片。
//   - low: 当前子切片的起始索引。
//   - high: 当前子切片的结束索引。
func quickSortRecursive(nums []int, low, high int) {
	if low < high {
		// 进行分区操作，获取基准元素的最终位置。
		pivotIndex := partition(nums, low, high)

		// 递归排序基准左右两边的子切片。
		quickSortRecursive(nums, low, pivotIndex-1)
		quickSortRecursive(nums, pivotIndex+1, high)
	}
}

// partition 分区操作 (Lomuto 分区方案)。
//
// 核心思想:
// 1. 选择最后一个元素作为基准 (pivot)。
// 2. 遍历数组，将所有小于基准的元素移到数组的左侧。
// 3. 将基准元素放到正确的位置，即所有小于它的元素之后，所有大于它的元素之前。
//
// 参数:
//   - nums: 整数切片。
//   - low: 当前子切片的起始索引。
//   - high: 当前子切片的结束索引。
//
// 返回值:
//   - int: 基准元素最终的索引位置。
func partition(nums []int, low, high int) int {
	pivot := nums[high] // 选择最后一个元素作为基准。
	i := low - 1        // i 指向小于基准的区域的右边界。

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i] // 交换，将小于基准的元素移到左侧。
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1] // 将基准元素放到正确位置。
	return i + 1
}

// HeapSort 堆排序
//
// 核心思想:
// 堆排序是一种利用堆数据结构（通常是最大堆）进行排序的算法。
// 1. 构建最大堆 (Build Max Heap): 将待排序的数组构建成一个最大堆。构建完成后，堆的根节点（数组的第一个元素）是最大的元素。
// 2. 堆排序 (Heapify): 将堆顶元素（最大值）与堆的最后一个元素交换。然后将堆的大小减 1，并对新的堆顶元素进行“下沉”操作 (sink)，以恢复堆的性质。重复此过程，直到堆的大小为 1。
//
// 时间复杂度: O(n log n) - 构建堆需要 O(n)，每次调整堆需要 O(log n)，共 n-1 次调整。
// 空间复杂度: O(1) - 原地排序，不需要额外的存储空间。
//
// 参数:
//   - nums: 需要排序的整数切片。
func HeapSort(nums []int) {
	n := len(nums)
	// 1. 构建最大堆：从最后一个非叶子节点开始，向上遍历所有非叶子节点，并对每个节点执行下沉操作。
	// 最后一个非叶子节点的索引是 n/2 - 1。
	for i := n/2 - 1; i >= 0; i-- {
		sink(nums, i, n)
	}

	// 2. 堆排序：将堆顶元素（最大值）与当前堆的最后一个元素交换，然后对剩余元素重新调整堆。
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 交换堆顶元素和当前堆的最后一个元素。
		sink(nums, 0, i)                    // 对剩余的堆（大小为 i）进行下沉操作，恢复堆性质。
	}
}

// sink 下沉操作，用于维护最大堆的性质。
//
// 核心思想:
// 将指定索引 `i` 处的元素与其子节点进行比较，如果子节点更大，则交换，并继续向下比较，直到该元素找到其在堆中的正确位置。
//
// 参数:
//   - nums: 堆数组。
//   - i: 需要下沉的节点索引。
//   - heapSize: 当前堆的有效大小。
func sink(nums []int, i, heapSize int) {
	for {
		// 假设左子节点是最大的节点
		maxIdx := 2*i + 1
		// 判断左子节点的index有无越界或溢出
		if maxIdx >= heapSize || maxIdx < 0 {
			break
		}
		rChild := 2*i + 2 // 右子节点索引。

		// 比较当前节点与左子节点，找出三者中最大的。
		if rChild < heapSize && nums[rChild] > nums[maxIdx] {
			maxIdx = rChild
		}
		// 如果当前节点已经是最大的，则下沉结束。
		if nums[i] >= nums[maxIdx] {
			break
		}

		// 交换当前节点与最大的子节点。
		nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
		// 继续向下沉。
		i = maxIdx
	}
}
