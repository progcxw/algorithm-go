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