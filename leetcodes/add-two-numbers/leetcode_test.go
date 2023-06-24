/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:17
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	l1 *ListNode
	l2 *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	ast := assert.New(t)

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

	res := &ListNode{7, nil}
	res1 := &ListNode{0, nil}
	res2 := &ListNode{8, nil}
	res.Next = res1
	res1.Next = res2

	var tests = []struct {
		input  input
		expect *ListNode
	}{
		{input{l1, l2}, res},
	}
	for _, test := range tests {
		ast.Equal(AddTwoNumbers(test.input.l1, test.input.l2), test.expect)
	}

}
