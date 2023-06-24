/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:27
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	s string
}

func TestMyAtoi(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect int
	}{
		{input{"42"}, 42},
		{input{"   -42"}, -42},
		{input{"4193 with words"}, 4193},
		{input{"  123 words and 987"}, 123},
		{input{"00000-42a1234"}, 0},
		{input{"00000-42"}, 0},
		{input{"  1 2 34 "}, 1},
		{input{" 0-123"}, 0},
		{input{"+1"}, 1},
		{input{"+-12"}, 0},
		{input{"123+-12"}, 123},
		{input{"  0000000000012345678"}, 12345678},
		{input{"-000000000000001"}, -1},
		{input{"     -000000000000001"}, -1},
		{input{"2147483648"}, 2147483647},
		{input{"0  123"}, 0},
		{input{"1  123"}, 1},
	}
	for _, test := range tests {
		ast.Equal(MyAtoi(test.input.s), test.expect)
	}

}

func TestName(t *testing.T) {
	MyAtoi("+123 words and 987")
}
