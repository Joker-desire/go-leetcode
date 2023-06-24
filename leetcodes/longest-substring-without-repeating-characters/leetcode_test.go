/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:02
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	s string
}

func TestLengthOfLongestSubstring(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect int
	}{
		{input{"abcabcbb"}, 3},
		{input{"bbbbb"}, 1},
		{input{"pwwkew"}, 3},
		{input{" "}, 1},
		{input{"au"}, 2},
		{input{""}, 0},
		{input{"aabaab!bb"}, 3},
	}
	for _, test := range tests {
		ast.Equal(LengthOfLongestSubstring(test.input.s), test.expect)
	}

}
