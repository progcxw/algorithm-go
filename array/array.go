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

// min 返回两个整数中的较小值
// 参数:
//   - x: 第一个整数
//   - y: 第二个整数
//
// 返回值:
//   - int: 较小的整数
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
