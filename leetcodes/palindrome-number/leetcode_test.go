/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 14:52
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	x int
}

func TestIsPalindrome(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect bool
	}{
		{input{121}, true},
		{input{-121}, false},
		{input{10}, false},
	}
	for _, test := range tests {
		ast.Equal(IsPalindrome(test.input.x), test.expect)
	}

}
