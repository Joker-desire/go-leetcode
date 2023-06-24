/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/22 10:02
 */

package leetcode

import "strconv"

// IsPalindrome 9. 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/palindrome-number
func IsPalindrome(x int) bool {
	// 将int类型转成字符串
	s := strconv.Itoa(x)
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
