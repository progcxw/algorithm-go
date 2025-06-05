package array

// TwoSum 两数之和
// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数的索引
// 时间复杂度: O(n)，空间复杂度: O(n)
// 参数:
//   - nums: 整数数组
//   - target: 目标和
//
// 返回值:
//   - []int: 和为目标值的两个数的索引数组，如果不存在则返回nil
func TwoSum(nums []int, target int) []int {
	// 创建数字map，key为数组成员内容，value为数组下标
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		index, ok := numMap[target-v]
		if !ok {
			numMap[v] = i
			continue
		}

		return []int{index, i}
	}

	return nil
}

// FindMedianSortedArrays 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组，找出这两个数组的中位数
// 时间复杂度: O(log(m+n))，空间复杂度: O(1)
// 参数:
//   - nums1: 第一个正序数组
//   - nums2: 第二个正序数组
//
// 返回值:
//   - float64: 两个数组合并后的中位数
func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		// 总长度为奇数，中位数是合并数组的中间元素
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}

	// 总长度为偶数，中位数是合并数组中间两个元素的平均值
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

// getKthElement 获取两个有序数组合并后的第k小的元素
// 使用二分查找的思想，每次排除部分元素
// 参数:
//   - nums1: 第一个有序数组
//   - nums2: 第二个有序数组
//   - k: 要查找的位置（从1开始）
//
// 返回值:
//   - int: 第k小的元素值
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			// nums1已经全部排除，直接返回nums2中的第k小元素
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			// nums2已经全部排除，直接返回nums1中的第k小元素
			return nums1[index1+k-1]
		}
		if k == 1 {
			// 只需要找第1小的元素，返回两个数组首元素的较小值
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			// pivot1 <= pivot2，排除nums1中的前half个元素
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			// pivot1 > pivot2，排除nums2中的前half个元素
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

// FindMedianSortedArraysByBinary 寻找两个正序数组的中位数（二分查找实现）
// 注意：此函数未完成实现
// 参数:
//   - nums1: 第一个正序数组
//   - nums2: 第二个正序数组
//
// 返回值:
//   - float64: 两个数组合并后的中位数
func FindMedianSortedArraysByBinary(nums1, nums2 []int) float64 {
	// 确保nums1的长度较小
	if len(nums1) > len(nums2) {
		return FindMedianSortedArraysByBinary(nums2, nums1)
	}

	// 从长度较小的nums1开始二分搜索（确保nums2的游标不会溢出），nums2的游标移动反向相同距离（使index1+index2 = len/2）
	index1, index2 := (len(nums1)-1)/2, (len(nums2)-1)/2
	left, right := 0, len(nums1)-1
	for left <= right {

	}

	// 以下代码未完成实现
	// index1移到尽头，则在nums2中找
	if index1 == 0 && nums1[index1] > nums2[index2] {
		// 未完成实现
	}
	if index1 == len(nums1)-1 && nums1[index1] < nums2[index2] {
		// 未完成实现
	}

	return 0 // 临时返回值，实际应返回计算结果
}

// FindKthLargest 查找数组中第k大的元素
// 使用快速选择算法，时间复杂度为O(n)
// 参数:
//   - nums: 输入的整数数组
//   - k: 要查找的第k大的元素（k从1开始）
//
// 返回值:
//   - int: 第k大的元素
func FindKthLargest(nums []int, k int) int {
	// 将第k大转换为第(n-k+1)小，这样可以复用快速选择的partition逻辑
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k+1)
}

// quickSelect 快速选择算法的核心实现
// 参数:
//   - nums: 输入数组
//   - left: 左边界
//   - right: 右边界
//   - k: 要查找的第k小的元素（k从1开始）
//
// 返回值:
//   - int: 第k小的元素
func quickSelect(nums []int, left, right, k int) int {
	// 如果区间只有一个元素，直接返回
	if left == right {
		return nums[left]
	}

	// 进行partition操作，返回pivot的位置
	pivotIndex := partition(nums, left, right)

	// 计算pivot是第几小的元素（从1开始计数）
	count := pivotIndex - left + 1

	if count == k {
		// 如果pivot正好是第k小的元素，返回它
		return nums[pivotIndex]
	} else if count > k {
		// 如果pivot位置大于k，说明第k小的元素在左半部分
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		// 如果pivot位置小于k，说明第k小的元素在右半部分
		// 注意要减去左半部分的元素个数
		return quickSelect(nums, pivotIndex+1, right, k-count)
	}
}

// partition 快速选择的分区操作
// 选择最右边的元素作为pivot，将数组分为小于pivot和大于pivot的两部分
// 参数:
//   - nums: 输入数组
//   - left: 左边界
//   - right: 右边界
//
// 返回值:
//   - int: pivot的最终位置
func partition(nums []int, left, right int) int {
	// 选择最右边的元素作为pivot
	pivot := nums[right]
	// i表示小于pivot的元素应该放置的位置
	i := left

	// 遍历left到right-1的元素
	for j := left; j < right; j++ {
		// 如果当前元素小于pivot，就放到前面
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// 将pivot放到正确的位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// FindKthLargestByHeap 使用最大堆查找数组中第k大的元素
// 时间复杂度：建堆O(n) + k次删除操作O(klogn) = O(n + klogn)
// 参数:
//   - nums: 输入的整数数组
//   - k: 要查找的第k大的元素（k从1开始）
//
// 返回值:
//   - int: 第k大的元素
func FindKthLargestByHeap(nums []int, k int) int {
	length := len(nums)
	// 原地建立最大堆
	for i := length/2 - 1; i >= 0; i-- {
		sinkForKth(nums, i, length)
	}

	// 进行k-1次删除堆顶操作
	for i := 0; i < k-1; i++ {
		// 将堆顶元素（最大值）与末尾元素交换
		nums[0], nums[length-1-i] = nums[length-1-i], nums[0]
		// 对新的堆顶元素进行下沉操作，注意堆的大小在减小
		sinkForKth(nums, 0, length-1-i)
	}

	// 返回当前堆顶元素，即第k大的元素
	return nums[0]
}

// sinkForKth 用于FindKthLargestByHeap的下沉操作
// 与原sink函数类似，但增加了堆大小参数
// 参数:
//   - nums: 堆数组
//   - i: 需要下沉的节点索引
//   - heapSize: 当前堆的大小
func sinkForKth(nums []int, i, heapSize int) {
	for {
		biggest := i
		// 计算左右子节点的索引
		lChild := 2*i + 1
		rChild := 2*i + 2

		// 比较当前节点与左子节点
		if lChild < heapSize && nums[lChild] > nums[biggest] {
			biggest = lChild
		}

		// 比较当前最大值与右子节点
		if rChild < heapSize && nums[rChild] > nums[biggest] {
			biggest = rChild
		}

		// 如果当前节点已经是最大的，结束下沉
		if biggest == i {
			break
		}

		// 交换当前节点与最大子节点
		nums[i], nums[biggest] = nums[biggest], nums[i]
		// 继续下沉
		i = biggest
	}
}

// FindKthLargestOptimized 优化版本的第k大元素查找
// 使用三路快速选择，特别适合处理有大量重复元素的情况
// 时间复杂度: O(n)，但对重复元素更友好
// 参数:
//   - nums: 输入的整数数组
//   - k: 要查找的第k大的元素（k从1开始）
//
// 返回值:
//   - int: 第k大的元素
func FindKthLargestOptimized(nums []int, k int) int {
	// 将第k大转换为第(n-k+1)小
	return quickSelectThreeWay(nums, 0, len(nums)-1, len(nums)-k+1)
}

// quickSelectThreeWay 三路快速选择的实现
// 将数组分为 <pivot, =pivot, >pivot 三部分
// 参数:
//   - nums: 输入数组
//   - left: 左边界
//   - right: 右边界
//   - k: 要查找的第k小的元素（k从1开始）
//
// 返回值:
//   - int: 第k小的元素
func quickSelectThreeWay(nums []int, left, right, k int) int {
	if left >= right {
		return nums[left]
	}

	// 随机选择pivot，避免最坏情况
	pivotIndex := left + (right-left)/2
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	pivot := nums[right]

	// lt表示“小于区”的右边界，gt表示“大于区”的左边界
	lt, i, gt := left-1, left, right

	// 三路划分
	for i < gt {
		if nums[i] < pivot {
			// 当前元素小于pivot，放入左边
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else if nums[i] > pivot {
			// 当前元素大于pivot，放入右边
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		} else {
			// 当前元素等于pivot，跳过
			i++
		}
	}

	// 将pivot放到正确的位置
	nums[gt], nums[right] = nums[right], nums[gt]
	gt++

	// 计算小于区和等于区的大小
	leftSize := lt - left + 1
	equalSize := gt - lt - 1

	if k <= leftSize {
		// 目标在小于区
		return quickSelectThreeWay(nums, left, lt, k)
	} else if k <= leftSize+equalSize {
		// 目标在等于区
		return pivot
	} else {
		// 目标在大于区
		return quickSelectThreeWay(nums, gt, right, k-leftSize-equalSize)
	}
}

// FindKthLargestByMinHeap 使用最小堆查找第k大元素
// 只维护k个元素的最小堆，适合处理大规模数据
// 时间复杂度: O(nlogk)
// 空间复杂度: O(k)
// 参数:
//   - nums: 输入的整数数组
//   - k: 要查找的第k大的元素（k从1开始）
//
// 返回值:
//   - int: 第k大的元素
func FindKthLargestByMinHeap(nums []int, k int) int {
	// 创建一个大小为k的最小堆
	heap := make([]int, k)

	// 初始化堆
	for i := 0; i < k; i++ {
		heap[i] = nums[i]
	}

	// 建立最小堆
	for i := k/2 - 1; i >= 0; i-- {
		sinkMin(heap, i, k)
	}

	// 处理剩余元素
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			sinkMin(heap, 0, k)
		}
	}

	return heap[0]
}

// sinkMin 最小堆的下沉操作
func sinkMin(nums []int, i, heapSize int) {
	for {
		smallest := i
		lChild := 2*i + 1
		rChild := 2*i + 2

		if lChild < heapSize && nums[lChild] < nums[smallest] {
			smallest = lChild
		}
		if rChild < heapSize && nums[rChild] < nums[smallest] {
			smallest = rChild
		}

		if smallest == i {
			break
		}

		nums[i], nums[smallest] = nums[smallest], nums[i]
		i = smallest
	}
}

// 三数之和
// 给定一个整数数组，找出所有满足三个数相加为0的组合
// 参数:
//   - nums: 整数数组
func ThreeSum(nums []int) [][]int {

}
