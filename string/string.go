package string

// LengthOfLongestSubstring 寻找字符串中不含重复字符的最长子串。
//
// 核心思想:
// 使用“滑动窗口”的方法来解决此问题。滑动窗口由两个指针 `head` (窗口起始) 和 `tail` (窗口结束) 维护。
//  1. `byteMap` (哈希表/字典) 用于存储窗口内每个字符的最新索引。
//  2. `tail` 指针向前移动，扩展窗口。
//  3. 当 `tail` 指针遇到一个在 `byteMap` 中已经存在的字符时，说明窗口内出现了重复字符。
//  4. 此时，需要收缩窗口，将 `head` 指针移动到重复字符的下一个位置 (即 `byteMap[s[tail]] + 1`)。
//     注意：`head` 只能向前移动，不能向后移动，所以需要取 `head` 和 `byteMap[s[tail]] + 1` 的最大值。
//  5. 每次 `tail` 移动后，更新 `byteMap` 中当前字符的索引，并计算当前窗口的长度 (`tail - head + 1`)，更新 `maxLen`。
//
// 时间复杂度: O(n) - `tail` 指针从头到尾遍历字符串一次，`head` 指针也最多从头到尾遍历一次。
// 空间复杂度: O(m) - `m` 是字符集的大小（例如，ASCII 字符集为 128 或 256），`byteMap` 最多存储 `m` 个字符。
//
// 参数:
//   - s: 输入的字符串。
//
// 返回值:
//   - int: 不含重复字符的最长子串的长度。
func LengthOfLongestSubstring(s string) int {
	var head, tail int            // 双指针，分别表示滑动窗口的起始和结束位置。
	var maxLen int                // 记录找到的最长无重复字符子串的长度。
	byteMap := make(map[byte]int) // 存储字符及其在字符串中最新出现的索引。

	// 遍历字符串，使用双指针维护一个滑动窗口。
	for tail < len(s) {
		// 如果当前字符 `s[tail]` 已经在 `byteMap` 中存在，
		// 并且其上次出现的索引 `idx` 在当前窗口 `[head, tail)` 内部，
		// 则说明出现了重复字符，需要移动 `head` 指针来收缩窗口。
		if idx, ok := byteMap[s[tail]]; ok && idx >= head {
			head = idx + 1 // 将 `head` 移动到重复字符的下一个位置。
		}

		// 更新当前字符 `s[tail]` 在 `byteMap` 中的最新索引。
		byteMap[s[tail]] = tail

		// 计算当前窗口的长度 `tail - head + 1`，并更新 `maxLen`。
		maxLen = max(maxLen, tail-head+1)

		// 移动 `tail` 指针，扩展窗口。
		tail++
	}

	return maxLen
}

// max 辅助函数，返回两个整数中的较大值。
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
