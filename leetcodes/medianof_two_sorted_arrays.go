package main

import (
	"fmt"
	"sort"
)

// 日期：2023年06月20日18:44:33
// 4. 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	num := len(nums1) / 2
	if len(nums1)%2 == 0 {
		return float64(nums1[num-1]+nums1[num]) / 2
	}
	return float64(nums1[num])
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println("median: ", median)
}
