/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 14:57
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	nums1 []int
	nums2 []int
}

func TestFindMedianSortedArrays(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect float64
	}{
		{input{[]int{1, 3}, []int{2}}, 2},
		{input{[]int{3, 2, 4}, []int{1, 2, 3}}, 2.5},
		{input{[]int{3, 3}, []int{2, 3, 4, 4, 5}}, 3},
	}
	for _, test := range tests {
		ast.Equal(FindMedianSortedArrays(test.input.nums1, test.input.nums2), test.expect)
	}

}
