package sort

// BubbleSort 冒泡排序
// 通过相邻元素比较和交换，将最大元素逐步"冒泡"到数组末尾
// 时间复杂度: O(n²)，空间复杂度: O(1)
// 参数:
//   - nums: 需要排序的整数数组
func BubbleSort(nums []int) {
	for j := len(nums) - 1; j > 0; j-- {
		// 每次i循环转完，j位置上的数位置定好，所以j从数组尾项遍历到首项时排序完毕
		for i := 0; i < j; i++ {
			// 比j大的位置都已排好序，只用遍历0到j
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

// SelectionSort 选择排序
// 每次从未排序部分找出最小元素，放到已排序部分的末尾
// 时间复杂度: O(n²)，空间复杂度: O(1)
// 参数:
//   - nums: 需要排序的整数数组
func SelectionSort(nums []int) {
	for j := 0; j < len(nums)-1; j++ {
		// 找到j到尾项中最小的数字，将其与nums[j]交换位置
		smallest := j
		for i := j; i < len(nums); i++ {
			if nums[i] < nums[smallest] {
				smallest = i
			}
		}
		nums[j], nums[smallest] = nums[smallest], nums[j]
	}
}

// InsertionSort 插入排序
// 构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
// 时间复杂度: O(n²)，空间复杂度: O(1)
// 参数:
//   - nums: 需要排序的整数数组
func InsertionSort(nums []int) {
	for j := 1; j < len(nums); j++ {
		// 把首项视为一个有序数组，从第二项开始将值插入到该有序数组中
		for i := j - 1; i >= 0; i-- {
			// 0~j-1的有序数组，从后往前swap
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

// ShellSort 希尔排序
// 希尔排序是把数组按下标的一定增量分组
// 对每组使用直接插入排序算法排序
// 随着增量逐渐减少，每组成员越来越多
// 当增量减至1时，整个数组恰被分成一组，算法便终止。
// 时间复杂度: 平均O(n^1.3)，空间复杂度: O(1)
// 参数:
//   - nums: 需要排序的整数数组
func ShellSort(nums []int) {
	// 以i为间隔分组，如[0, i, 2i]、[1, 1+i, 1+2i]，分组间隔从len(nums)/2逐渐除以2，直到间隔为1
	for i := len(nums) / 2; i > 0; i /= 2 {
		// 组内进行插入排序
		// 跳过第一个i间隔后，j进行自增1的循环，因为每组相同位置的差值为1
		// 例：groupA=[0, i, 2i]，groupB=[1, 1+i, 1+2i]，groupB[j] = groupA[j]+1
		// 所以使用j++的for循环并在循环内swap符合条件的nums[j]与nums[j-i],能对所有组进行插入排序
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[j-i] {
				nums[j], nums[j-i] = nums[j-i], nums[j]
			}
		}
	}
}

// MergeSort 归并排序
// 采用分治法，将已有序的子序列合并，得到完全有序的序列
// 时间复杂度: O(nlogn)，空间复杂度: O(n)
// 参数:
//   - nums: 需要排序的整数数组
//
// 返回值:
//   - []int: 排序后的数组
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])  // 递归排序左半部分
	right := MergeSort(nums[mid:]) // 递归排序右半部分

	// 合并两个有序数组
	l, r, i := 0, 0, 0
	ret := make([]int, len(nums))
	for l < mid && r < len(right) {
		if left[l] < right[r] {
			ret[i] = left[l]
			l++
		} else {
			ret[i] = right[r]
			r++
		}

		i++
	}

	// 左或右数组遍历完后，另一边数组还有剩余成员的情况
	for l < mid || r < len(right) {
		if l < mid {
			ret[i] = left[l]
			l++
		}

		if r < len(right) {
			ret[i] = right[r]
			r++
		}

		i++
	}

	return ret
}

// QuickSort 快速排序
// 通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小
// 时间复杂度: 平均O(nlogn)，最坏O(n²)，空间复杂度: O(logn)
// 参数:
//   - nums: 需要排序的整数数组
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
	QuickSort(nums[:pos])   // 递归排序小于基准的部分
	QuickSort(nums[pos+1:]) // 递归排序大于基准的部分
}

// HeapSort 堆排序
// 利用堆这种数据结构所设计的一种排序算法
// 时间复杂度: O(nlogn)，空间复杂度: O(1)
// 参数:
//   - nums: 需要排序的整数数组
func HeapSort(nums []int) {
	// 从最后一个非叶子节点开始建立最大堆
	// 最后一个非叶子节点的索引是 len(nums)/2 - 1
	// 例如：数组长度为6时，最后一个非叶子节点索引为2
	// 数组：[0,1,2,3,4,5]
	// 树结构：
	//      0
	//     / \
	//    1   2    <- 最后一个非叶子节点
	//   / \ /
	//  3  4 5
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i)
	}

	// 最大堆形成后，根节点(索引0)为最大值
	// 将最大值与末尾元素交换，然后对剩余元素重新建立最大堆
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		// 对剩余元素重新建立最大堆
		// 注意：这里只对nums[:i]进行sink操作，因为nums[i:]已经排好序
		sink(nums[:i], 0)
	}
}

// sink 下沉操作，用于堆排序
// 将指定节点下沉到合适位置，保持最大堆性质
// 时间复杂度: O(logn)
// 参数:
//   - nums: 堆数组
//   - i: 需要下沉的节点索引
func sink(nums []int, i int) {
	for {
		// 使用li(largest index)记录当前子树中最大节点的index，初始化为左子节点
		li := 2*i + 1
		length := len(nums)
		// 检查li是否越界或溢出
		if li > length || li < 0 {
			break
		}

		// rc(right child)为右子节点
		rc := li + 1
		if rc < length && nums[rc] > nums[li] {
			li = rc
		}

		// 如果当前节点比最大子节点还大，说明已经下沉到正确位置
		if nums[i] > nums[li] {
			break
		}

		// 交换当前节点与最大子节点
		nums[i], nums[li] = nums[li], nums[i]
		// 继续下沉被交换下来的节点
		i = li
	}
}
