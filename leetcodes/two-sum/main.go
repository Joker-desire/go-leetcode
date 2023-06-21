package main

import "fmt"

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/two-sum

func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for y, val2 := range nums {
			// 判断两数相加并且下标不能相等
			if val+val2 == target && i != y {
				return []int{i, y}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
