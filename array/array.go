package array

import (
	"math"
	"sort"
)

// TwoSum 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
// 时间复杂度: O(n) - 我们只遍历数组一次。
// 空间复杂度: O(n) - 我们使用了一个 map 来存储数字及其索引，最坏情况下需要存储所有元素。
//
// 参数:
//   - nums: 整数数组。
//   - target: 目标和。
//
// 返回值:
//   - []int: 和为目标值的两个数的索引组成的切片，如果不存在则返回 nil。
func TwoSum(nums []int, target int) []int {
	// 创建一个 map，键为数组中的数字，值为其对应的下标。
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		// 检查 map 中是否存在 target-v 这个键
		if index, ok := numMap[target-v]; ok {
			// 如果存在，说明找到了两个和为 target 的数，返回它们的索引。
			return []int{index, i}
		}
		// 如果不存在，将当前数字和它的索引存入 map。
		numMap[v] = i
	}
	// 如果遍历完整个数组都没有找到，返回 nil。
	return nil
}

// FindMedianSortedArrays 寻找两个正序数组的中位数 (递归/分治法)
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 核心思想:
// 要找到中位数，我们实际上需要找到合并后数组中的第 k 小的元素。
// 如果总长度是奇数，中位数就是第 (m+n)/2 + 1 小的元素。
// 如果总长度是偶数，中位数就是第 (m+n)/2 和第 (m+n)/2 + 1 小的两个元素的平均值。
// `getKthElement` 函数通过每次比较两个数组的第 k/2 个元素，排除掉不可能包含第 k 小元素的前 k/2 个元素，从而将问题规模减半。
//
// 时间复杂度: O(log(m+n)) - 每次递归调用都将 k 的值减半。
// 空间复杂度: O(log(m+n)) - 递归调用栈的深度。
//
// 参数:
//   - nums1: 第一个正序数组。
//   - nums2: 第二个正序数组。
//
// 返回值:
//   - float64: 两个数组合并后的中位数。
func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0.0
	}
	if totalLength%2 == 1 {
		// 总长度为奇数，中位数是合并后数组的中间那个元素。
		midIndex := totalLength/2 + 1
		return float64(getKthElement(nums1, nums2, midIndex))
	}
	// 总长度为偶数，中位数是合并后数组中间两个元素的平均值。
	midIndex1, midIndex2 := totalLength/2, totalLength/2+1
	k1 := getKthElement(nums1, nums2, midIndex1)
	k2 := getKthElement(nums1, nums2, midIndex2)
	return float64(k1+k2) / 2.0
}

// getKthElement 获取两个有序数组中第 k 小的元素。
//
// 参数:
//   - nums1: 第一个有序数组。
//   - nums2: 第二个有序数组。
//   - k: 要查找的位置（从 1 开始计数）。
//
// 返回值:
//   - int: 第 k 小的元素值。
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			// 如果 nums1 已经全部被排除，直接从 nums2 中返回。
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			// 如果 nums2 已经全部被排除，直接从 nums1 中返回。
			return nums1[index1+k-1]
		}
		if k == 1 {
			// 如果 k=1，说明要找的就是当前两个数组头部的较小值。
			return min(nums1[index1], nums2[index2])
		}

		// 每次尝试排除 k/2 个元素。
		half := k / 2
		// 计算需要比较的两个元素在各自数组中的索引。
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 <= pivot2 {
			// 如果 nums1 的 pivot 小于等于 nums2 的 pivot，
			// 那么 nums1 中从 index1 到 newIndex1 的所有元素都不可能是第 k 小的元素，可以排除。
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			// 反之，排除 nums2 中的相应元素。
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

// FindMedianSortedArraysByBinary 寻找两个正序数组的中位数 (二分查找法)
//
// 核心思想:
// 该方法通过在较短的数组上进行二分查找来找到一个完美的分割点。
// 这个分割点 i 将短数组分为 left_part 和 right_part。
// 同时，根据中位数的定义，我们可以在长数组中计算出对应的分割点 j。
// 完美的分割点需要满足: `nums1[i-1] <= nums2[j]` 并且 `nums2[j-1] <= nums1[i]`。
// 这意味着第一个数组左侧的所有元素都小于等于第二个数组右侧的所有元素，反之亦然。
// 找到这个点后，中位数就可以由两个部分四个边界元素轻松计算得出。
//
// 时间复杂度: O(log(min(m,n))) - 对较短的数组进行二分查找。
// 空间复杂度: O(1) - 只使用了常数个变量。
//
// 参数:
//   - nums1: 第一个正序数组。
//   - nums2: 第二个正序数组。
//
// 返回值:
//   - float64: 两个数组合并后的中位数。
func FindMedianSortedArraysByBinary(nums1, nums2 []int) float64 {
	// 为保证复杂度，对较短的数组进行二分查找。
	if len(nums1) > len(nums2) {
		return FindMedianSortedArraysByBinary(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0.0
	}
	left, right := 0, m
	// median1 是中位数左边的元素，median2 是中位数右边的元素。
	// totalLeft 是合并后数组左半部分的元素个数。
	median1, median2, totalLeft := 0, 0, (m+n+1)/2

	for left <= right {
		// i 是 nums1 的分割线索引，j 是 nums2 的分割线索引。
		i := left + (right-left)/2
		j := totalLeft - i

		// 获取 nums1 和 nums2 分割线两侧的四个值。
		nums1LeftMax := math.MinInt32
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}
		nums1RightMin := math.MaxInt32
		if i < m {
			nums1RightMin = nums1[i]
		}
		nums2LeftMax := math.MinInt32
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}
		nums2RightMin := math.MaxInt32
		if j < n {
			nums2RightMin = nums2[j]
		}

		if nums1LeftMax <= nums2RightMin {
			// 找到了合适的分割，计算中位数。
			median1 = max(nums1LeftMax, nums2LeftMax)
			median2 = min(nums1RightMin, nums2RightMin)
			left = i + 1 // 继续搜索是为了找到最右边的满足条件的i
		} else {
			// i 太大了，需要向左移动。
			right = i - 1
		}
	}

	// 如果总长度是奇数，中位数就是左半部分的最大值。
	if (m+n)%2 == 1 {
		return float64(median1)
	}
	// 如果总长度是偶数，中位数是左半部分最大值和右半部分最小值的平均。
	return float64(median1+median2) / 2.0
}

// ThreeSum 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 请你找出所有和为 0 且不重复的三元组。
//
// 核心思想:
// 1. 对数组进行排序，这是为了方便地跳过重复元素和使用双指针法。
// 2. 遍历数组，将当前元素 `nums[i]` 作为第一个数。
// 3. 使用左右双指针 `left` 和 `right` 在 `i` 之后的部分寻找另外两个数。
// 4. `left` 从 `i+1` 开始，`right` 从数组末尾开始。
// 5. 计算三数之和 `sum`：
//   - 如果 `sum == 0`，则找到了一个解。将 `left` 和 `right` 向中间移动，并跳过所有重复的元素。
//   - 如果 `sum < 0`，说明和太小，需要增大，将 `left` 右移。
//   - 如果 `sum > 0`，说明和太大，需要减小，将 `right` 左移。
//
// 6. 为了避免第一个数的重复，如果 `nums[i]` 与前一个元素相同，则跳过。
//
// 时间复杂度: O(n^2) - 排序需要 O(n log n)，双指针遍历需要 O(n^2)。
// 空间复杂度: O(log n) or O(n) - 取决于排序算法使用的空间。
//
// 参数:
//   - nums: 整数数组。
//
// 返回值:
//   - [][]int: 所有不重复的三元组。
func ThreeSum(nums []int) [][]int {
	// 首先对数组进行排序
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// 避免重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				// 找到了一个解
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 避免重复的第二个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 避免重复的第三个数
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 移动指针寻找新的可能
				left++
				right--
			} else if sum < target {
				// 和太小，左指针右移
				left++
			} else {
				// 和太大，右指针左移
				right--
			}
		}
	}
	return result
}

// FindKthLargest 查找数组中第 k 大的元素 (快速选择)
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 核心思想:
// 快速选择算法是快速排序的变体。它的目标是找到第 k 小（或大）的元素，而不是完全排序整个数组。
// 1. 选择一个枢轴（pivot）元素，并将数组分区，使得所有小于枢轴的元素都在它左边，所有大于的都在右边。
// 2. 分区后，枢轴位于其最终排序位置 `p`。
// 3. 比较 `p` 和 `k`：
//   - 如果 `p == k`，则枢轴就是我们要找的元素。
//   - 如果 `p > k`，则第 k 大的元素在左半部分，递归地在左侧查找。
//   - 如果 `p < k`，则第 k 大的元素在右半部分，递归地在右侧查找。
//
// 平均情况下，每次都能将问题规模减半。
//
// 时间复杂度: O(n) - 平均情况。O(n^2) - 最坏情况（但通过随机化 pivot 可以有效避免）。
// 空间复杂度: O(log n) - 递归调用栈的深度。
//
// 参数:
//   - nums: 整数数组。
//   - k: 要查找的第 k 大的位置（从 1 开始）。
//
// 返回值:
//   - int: 第 k 大的元素。
func FindKthLargest(nums []int, k int) int {
	// 第 k 大的元素，等价于第 (n-k) 小的元素（如果从0开始索引）。
	// 或者说，等价于第 (n-k+1) 小的元素（如果从1开始索引）。
	// quickSelect 函数查找的是第 k 小的元素。
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k+1)
}

// quickSelect 快速选择算法的递归实现。
//
// 参数:
//   - nums: 数组。
//   - left: 当前查找区间的左边界。
//   - right: 当前查找区间的右边界。
//   - k: 要查找的第 k 小的元素（k 基于 1 开始计数，相对于整个数组）。
//
// 返回值:
//   - int: 第 k 小的元素。
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// 进行分区操作，返回 pivot 的最终位置。
	pivotIndex := partition(nums, left, right)
	// pivotIndex 是从0开始的，所以它前面有 pivotIndex - left 个元素。
	// pivot 是当前分区中的第 (pivotIndex - left + 1) 小的元素。
	count := pivotIndex - left + 1

	if count == k {
		return nums[pivotIndex]
	} else if count > k {
		// 第 k 小的元素在左半部分。
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		// 第 k 小的元素在右半部分。
		// 注意 k 需要减去左半部分的元素数量。
		return quickSelect(nums, pivotIndex+1, right, k-count)
	}
}

// partition Lomuto 分区方案的实现。
//
// 参数:
//   - nums: 数组。
//   - left: 左边界。
//   - right: 右边界。
//
// 返回值:
//   - int: pivot 的最终索引位置。
func partition(nums []int, left, right int) int {
	// 选择最右边的元素作为 pivot。
	pivot := nums[right]
	// i 是小于 pivot 区域的右边界。
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将 pivot 放到正确的位置。
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// FindKthLargestByHeap 使用最大堆查找数组中第 k 大的元素。
//
// 核心思想:
// 1. 将整个数组原地构建成一个最大堆。根节点是最大的元素。
// 2. 执行 k-1 次“删除”操作：将堆顶（最大元素）与堆的最后一个元素交换，然后缩小堆的大小并对新的堆顶执行下沉（sift-down）操作以维持堆属性。
// 3. 经过 k-1 次操作后，堆顶的元素就是第 k 大的元素。
//
// 时间复杂度: O(n + k log n) - 建堆 O(n)，k-1 次删除 O(k log n)。
// 空间复杂度: O(1) - 原地建堆。
//
// 参数:
//   - nums: 整数数组。
//   - k: 要查找的第 k 大的位置（从 1 开始）。
//
// 返回值:
//   - int: 第 k 大的元素。
func FindKthLargestByHeap(nums []int, k int) int {
	length := len(nums)
	// 从最后一个非叶子节点开始，自下而上构建最大堆。
	for i := length/2 - 1; i >= 0; i-- {
		sinkForKth(nums, i, length)
	}

	// 执行 k-1 次删除操作。
	for i := 0; i < k-1; i++ {
		// 将堆顶（当前最大值）与堆末尾元素交换。
		nums[0], nums[length-1-i] = nums[length-1-i], nums[0]
		// 缩小堆的大小，并对新的堆顶进行下沉。
		sinkForKth(nums, 0, length-1-i)
	}

	// 此时的堆顶即为第 k 大的元素。
	return nums[0]
}

// sinkForKth 最大堆的下沉操作。
//
// 参数:
//   - nums: 堆数组。
//   - i: 需要下沉的节点索引。
//   - heapSize: 当前堆的有效大小。
func sinkForKth(nums []int, i, heapSize int) {
	for {
		biggest := i
		lChild := 2*i + 1
		rChild := 2*i + 2

		if lChild < heapSize && nums[lChild] > nums[biggest] {
			biggest = lChild
		}
		if rChild < heapSize && nums[rChild] > nums[biggest] {
			biggest = rChild
		}
		if biggest == i {
			break
		}
		nums[i], nums[biggest] = nums[biggest], nums[i]
		i = biggest
	}
}

// FindKthLargestByMinHeap 使用最小堆查找第 k 大元素。
//
// 核心思想:
// 维护一个大小为 k 的最小堆。
// 1. 先用数组的前 k 个元素构建一个最小堆。堆顶是这 k 个元素中的最小值。
// 2. 遍历数组中余下的元素：
//   - 如果当前元素小于或等于堆顶元素，则忽略它（因为它不可能成为第 k 大的元素）。
//   - 如果当前元素大于堆顶元素，则用它替换堆顶，并执行下沉操作以维持最小堆属性。
//
// 3. 遍历结束后，堆顶的元素就是整个数组中第 k 大的元素。
//
// 这种方法特别适合处理海量数据，因为我们只需要在内存中保留 k 个元素。
//
// 时间复杂度: O(n log k) - 遍历 n 个元素，每次堆操作 log k。
// 空间复杂度: O(k) - 用于存储堆。
//
// 参数:
//   - nums: 整数数组。
//   - k: 要查找的第 k 大的位置（从 1 开始）。
//
// 返回值:
//   - int: 第 k 大的元素。
func FindKthLargestByMinHeap(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return -1 // 或者其他错误处理
	}
	// 创建一个大小为 k 的最小堆。
	heap := nums[:k]

	// 建立最小堆。
	for i := k/2 - 1; i >= 0; i-- {
		sinkMin(heap, i, k)
	}

	// 遍历剩余的元素。
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			sinkMin(heap, 0, k)
		}
	}

	return heap[0]
}

// sinkMin 最小堆的下沉操作。
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

// FindMinDifference 获取数组中的最小差值 (排序法)
//
// 核心思想:
// 1. 对数组进行排序。
// 2. 排序后，任意两个元素之间的最小差值一定出现在相邻的两个元素之间。
// 3. 遍历排序后的数组，计算并比较所有相邻元素的差值，找出最小值。
//
// 这种方法比桶排序更通用，不受数据范围的影响。
//
// 时间复杂度: O(n log n) - 主要由排序决定。
// 空间复杂度: O(log n) or O(n) - 取决于排序算法。
//
// 参数:
//   - arr: 整数数组。
//
// 返回值:
//   - int: 数组中任意两个元素的最小差值。
func FindMinDifference(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}

	// 对数组进行排序
	sort.Ints(arr)

	minDiff := math.MaxInt32
	for i := 1; i < n; i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
