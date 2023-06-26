/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/26 11:56
 */

package three_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	nums []int
}

func TestThreeSum(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect [][]int
	}{
		{input{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{input{[]int{0, 1, 1}}, [][]int{}},
		{input{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
	}
	for _, test := range tests {
		ast.Equal(ThreeSum(test.input.nums), test.expect)
	}

}
