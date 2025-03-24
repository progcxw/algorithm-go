package string

// LengthOfLongestSubstring 寻找字符串中不含重复字符的最长子串
// 参数:
//   - s: 输入的字符串
//
// 返回值:
//   - int: 最长子串的长度
func LengthOfLongestSubstring(s string) int {
	var head, tail int            // 双指针
	var maxLen int                // 最长子串的长度
	byteMap := make(map[byte]int) // 记录字符出现的位置

	// 遍历字符串，使用双指针维护一个滑动窗口
	for tail < len(s) {
		// 如果字符已经出现过，则移动头指针
		if _, ok := byteMap[s[tail]]; ok {
			head = max(head, byteMap[s[tail]]+1)
		}

		// 记录字符出现的位置
		byteMap[s[tail]] = tail
		// 更新最长子串的长度
		maxLen = max(maxLen, tail-head+1)
		// 移动尾指针
		tail++
	}

	return maxLen
}
