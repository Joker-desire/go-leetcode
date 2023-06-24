/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/20 18:44
 */

package leetcode

import "sort"

// FindMedianSortedArrays 4. 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	num := len(nums1) / 2
	if len(nums1)%2 == 0 {
		return float64(nums1[num-1]+nums1[num]) / 2
	}
	return float64(nums1[num])
}
