/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/26 11:56
 */

package three_sum

import "sort"

// ThreeSum 15. 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针, j从i+1开始向右移，k从n-1开始向左移
		j, k := i+1, n-1
		for j < k {
			// 三数之和
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// 如果和为0，将三个数加入到结果中
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++ // j右移
				k-- // k左移
				// 如果j和j-1相等，j再次右移
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// 如果k和k+1相等，k再次左移
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				// 如果和大于0，k需要左移
				k--
			} else {
				// 如果和小于0，j需要右移
				j++
			}
		}
	}
	return ans

}
