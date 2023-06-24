/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/22 08:52
 */

package leetcode

// LongestPalindrome 5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。
//如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-palindromic-substring
func LongestPalindrome(s string) string {
	// 是否是回文字符串判断函数
	var isPalindrome = func(s string) bool {
		// 字符串长度
		length := len(s)
		// 长度/2
		point := length / 2
		for i := 0; i < point; i++ {
			if s[i] == s[length-i-1] {
				continue
			} else {
				return false
			}
		}
		return true
	}

	// 先判断s是否回文字符串，如果是则直接返回
	if isPalindrome(s) {
		return s
	}
	// 字符串长度
	length := len(s)
	// 存储组合字符
	var vals []string

	// 循环进行组合字符
	for i := 0; i < length-1; i++ {
		// 需要循环多少次
		loop := length - i
		for index := 0; index < loop; index++ {
			vals = append(vals, s[index:index+i+1])
		}
	}
	// 通过组合进行判断是否为回文字符串
	maxLen := 0  // 定义回文最大长度
	maxStr := "" // 定义回文最大长度的值
	for _, val := range vals {
		if isPalindrome(val) {
			// 如果回文字符串大于最大长度，则进行覆写
			if len(val) > maxLen {
				maxLen = len(val)
				maxStr = val
			}
		}
	}
	// 返回最大长度的值
	return maxStr
}
