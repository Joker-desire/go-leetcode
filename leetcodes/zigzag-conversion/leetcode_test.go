/**
 * @Author: yy
 * @Description:
 * @File:  leetcode_test
 * @Version: 1.0.0
 * @Date: 2023/06/24 00:06
 */

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type input struct {
	s       string
	numRows int
}

func TestConvert(t *testing.T) {
	ast := assert.New(t)
	var tests = []struct {
		input  input
		expect string
	}{
		{input{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{input{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{input{"A", 1}, "A"},
	}
	for _, test := range tests {
		ast.Equal(Convert(test.input.s, test.input.numRows), test.expect)
	}

}
