/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:13
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	s string
}

func TestLongestPalindrome(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect string
	}{
		{input{"babad"}, "bab"},
		{input{"cbbd"}, "bb"},
		{input{"a"}, "a"},
	}
	for _, test := range tests {
		ast.Equal(LongestPalindrome(test.input.s), test.expect)
	}

}
