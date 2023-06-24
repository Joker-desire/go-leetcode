/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:16
 */

package leetcode

// 2. 两数相加
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/add-two-numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个结果node
	var resNode *ListNode
	//	resNode = &ListNode{0, nil}
	// 定义一个标识是否需要+1
	isNum := 0
	// 获取Next
	l1NextNode := l1
	l2NextNode := l2
	// 获取链表中的数值
	for {
		sum := isNum
		// 判断是否将链表数据取完
		if l1NextNode == nil && l2NextNode == nil {
			// 表示l1和l2链表数据已取完,需要判断是否需要+1
			if isNum == 0 {
				break
			}
		} else if l1NextNode == nil {
			sum += l2NextNode.Val
			l2NextNode = l2NextNode.Next
		} else if l2NextNode == nil {
			sum += l1NextNode.Val
			l1NextNode = l1NextNode.Next
		} else {
			sum += l1NextNode.Val + l2NextNode.Val
			// 将下一节点进行赋值
			l1NextNode = l1NextNode.Next
			l2NextNode = l2NextNode.Next
		}
		if sum >= 10 {
			sum -= 10
			isNum = 1
		} else {
			isNum = 0
		}
		// 新节点
		newNode := &ListNode{sum, nil}
		// 将新的节点进行添加
		next := resNode
		if next != nil {
			for next.Next != nil {
				next = next.Next
			}
			next.Next = newNode
		} else {
			resNode = newNode
		}
	}
	return resNode
}
