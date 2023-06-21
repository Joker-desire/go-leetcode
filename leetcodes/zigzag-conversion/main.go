package main

import (
	"fmt"
	"strings"
)

// 时间：2023年06月21日22:33:23
// 6. N字形变换
//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//请你实现这个将字符串进行指定行数变换的函数：
//string convert(string s, int numRows);

func convert(s string, numRows int) string {
	var res [][]string
	// 通过传入的行数构建多维数组
	for i := 0; i < numRows; i++ {
		res = append(res, []string{})
	}
	// 有几列有单个元素
	singlelement := numRows - 2
	// 取得字符串下标
	index := 0
	// 字符串长度
	length := len(s)
	// 是否退出循环
	isBreak := false
	for {
		// 如果循环次数比剩余取字符串小标大
		if numRows > (length - index) {
			numRows = length - index
			isBreak = true // 标识下面循环结束后退出循环
		}
		// 如果是第一列，则所有数组都需要写入数据
		for i := 0; i < numRows; i++ {
			res[i] = append(res[i], string(s[index]))
			index++ // 取字符串下标+1
		}
		// 判断是否退出循环
		if isBreak {
			break
		}

		// 如果循环次数比剩余取字符串小标大
		if singlelement > (length - index) {
			singlelement = length - index
			isBreak = true // 标识下面循环结束后退出循环
		}
		// 如果是第二列
		for i := 1; i <= singlelement; i++ {
			res[numRows-i-1] = append(res[numRows-i-1], string(s[index]))
			index++ // 取字符串下标+1
		}
		// 判断是否退出循环
		if isBreak {
			break
		}
	}

	resStr := ""
	// 整合结果
	for _, val := range res {
		resStr += strings.Join(val, "")
	}
	return resStr
}

func main() {
	cases := [][]any{{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"}, {"PAYPALISHIRING", 4, "PINALSIGYAHRPI"}, {"A", 1, "A"}}
	for _, val := range cases {
		s2 := convert(val[0].(string), val[1].(int))
		fmt.Println(s2 == val[2].(string))
	}

}
