/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/20 15:09
 */

package leetcode

// LengthOfLongestSubstring 3. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
func LengthOfLongestSubstring(s string) int {
	// 如果字符串为空则直接返回结果0
	if s == "" {
		return 0
	}
	// 初始化下标
	movIndex := 1
	// 定义结果变量，将第一个变量写入Map
	res := map[uint8]int{
		s[0]: 0,
	}
	// 数量汇总
	count := len(res)
	for {
		// 如果移动下标大于等于字符串成长度，则退出循环
		if movIndex >= len(s) {
			break
		}
		// 获取指定的数值
		val := s[movIndex]
		// 判断是否已存在结果中
		if index, ok := res[val]; ok {
			// 如果已存在
			// 1. 获取结果长度，如果长度大于count则进行更新
			if len(res) >= count {
				count = len(res)
			}
			// 2. 将结果中小于此下标的清除掉
			for k, v := range res {
				if v < index {
					delete(res, k)
				}
			}
			// 将movIndex下标更新到res[val]中
			res[val] = movIndex
		} else {
			// 如果不存在，则写入结果中
			res[val] = movIndex
			// 并且将数量汇总+1
			if len(res) >= count {
				count = len(res)
			}
		}
		// 向下移动
		movIndex++
	}
	return count
}
