package sort

// BubbleSort 冒泡排序
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
			// 当前子树是最大堆，无需swap，直接break
			break
		}

		// swap 最大节点与i节点
		nums[i], nums[biggest] = nums[biggest], nums[i]

		// 此时被换下来的节点的子节点可能比它大，往下sink
		i = biggest
	}
}
