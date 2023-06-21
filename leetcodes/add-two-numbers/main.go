package main

import "fmt"

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/add-two-numbers
//
//type ListNode struct {
//	Val   int
//	Next  *ListNode
//	count int
//}
//
//func NewListNode() *ListNode {
//	return &ListNode{0, nil, 0}
//}
//
//// Length 获取链表长度
//func (n *ListNode) Length() int {
//	return n.count
//}
//
//// Add 添加元素
//func (n *ListNode) Add(val int) {
//	next := n.Next
//	if next != nil {
//		for next.Next != nil {
//			next = next.Next
//		}
//		next.Next = &ListNode{Val: val, Next: nil}
//	} else {
//		n.Next = &ListNode{Val: val, Next: nil}
//	}
//	n.count++
//}
//
//// Get 获取链表某一下标的数据
//func (n *ListNode) Get(index int) int {
//	next := n.Next
//	for i := 0; i < index; i++ {
//		next = next.Next
//	}
//	return next.Val
//}
//
//// Traverse 链表遍历
//func (n *ListNode) Traverse() []int {
//	var res []int
//	node := n.Next
//	for {
//		if node == nil {
//			break
//		}
//		res = append(res, node.Val)
//		node = node.Next
//	}
//	return res
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	l1Len := l1.Length()
//	l2len := l2.Length()
//	loop := l1Len
//	maxLenNode := l1
//	if l1Len > l2len {
//		loop = l2len
//	} else if l1Len == l2len {
//		maxLenNode = nil
//	} else {
//		maxLenNode = l2
//	}
//	newListNode := NewListNode()
//	isNum := 0
//	for i := 0; i < loop; i++ {
//		sum := l1.Get(i) + l2.Get(i) + isNum
//		if sum >= 10 {
//			sum -= 10
//			isNum = 1
//		} else {
//			isNum = 0
//		}
//		newListNode.Add(sum)
//	}
//	if maxLenNode != nil {
//		for i := loop; i < maxLenNode.Length(); i++ {
//			sum := maxLenNode.Get(i) + isNum
//			if sum >= 10 {
//				sum -= 10
//				isNum = 1
//			} else {
//				isNum = 0
//			}
//			newListNode.Add(sum)
//		}
//	}
//	if isNum == 1 {
//		newListNode.Add(1)
//	}
//	return newListNode
//}
//
//func TestOne() {
//	listNode1 := NewListNode()
//	listNode1.Add(0)
//	listNode2 := NewListNode()
//	listNode2.Add(0)
//	node := addTwoNumbers(listNode1, listNode2)
//	traverse := node.Traverse()
//	fmt.Println("node ", traverse)
//}
//func TestTwo() {
//	listNode1 := NewListNode()
//	listNode1.Add(2)
//	listNode1.Add(4)
//	listNode1.Add(3)
//
//	listNode2 := NewListNode()
//	listNode2.Add(5)
//	listNode2.Add(6)
//	listNode2.Add(4)
//	node := addTwoNumbers(listNode1, listNode2)
//	traverse := node.Traverse()
//	fmt.Println("node ", traverse)
//}
//
//func TestThree() {
//	listNode1 := NewListNode()
//	listNode1.Add(9)
//	listNode1.Add(9)
//	listNode1.Add(9)
//	listNode1.Add(9)
//	listNode1.Add(9)
//	listNode1.Add(9)
//	listNode1.Add(9)
//
//	listNode2 := NewListNode()
//	listNode2.Add(9)
//	listNode2.Add(9)
//	listNode2.Add(9)
//	listNode2.Add(9)
//	node := addTwoNumbers(listNode1, listNode2)
//	traverse := node.Traverse()
//	fmt.Println("node ", traverse)
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
func main() {
	l1 := &ListNode{2, nil}
	node1 := &ListNode{4, nil}
	node2 := &ListNode{3, nil}
	l1.Next = node1
	node1.Next = node2

	l2 := &ListNode{5, nil}
	node21 := &ListNode{6, nil}
	node22 := &ListNode{4, nil}
	l2.Next = node21
	node21.Next = node22

	listNode := addTwoNumbers(l1, l2)
	next := listNode
	for {
		if next == nil {
			break
		}
		fmt.Println("val: ", next.Val)
		next = next.Next
	}

}
