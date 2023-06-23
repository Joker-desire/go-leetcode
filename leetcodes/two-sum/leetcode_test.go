/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/23 23:32
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect []int
	}{
		{input{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{input{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{input{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, test := range tests {
		ast.Equal(TwoSum(test.input.nums, test.input.target), test.expect)
	}

}
