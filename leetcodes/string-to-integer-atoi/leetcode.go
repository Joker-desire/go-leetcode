/**
 * @Author: yy
 * @Description:
 * @File:  leetcode
 * @Version: 1.0.0
 * @Date: 2023/06/24 15:27
 */

package leetcode

import (
	"strconv"
)

// MyAtoi 8. 字符串转换整数 (atoi)
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
// 函数 myAtoi(string s) 的算法如下：
// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −2^31 ，大于 2^31 − 1 的整数应该被固定为 2^31 − 1 。
// 返回整数作为最终结果。
// 注意：
// 本题中的空白字符只包括空格字符 ' ' 。
// 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
func MyAtoi(s string) int {
	if s == "" {
		return 0
	}
	// 空格：32
	// 数字：48~57 0~9
	// -：45
	// +：43
	// 是否有符号，默认为0为正数，1为正数，-1为负数
	isSymbol := 0
	// 是否出现过非空格字符
	isNonSpace := false
	// 是否先出现的0
	isZero := false
	// 第一个非0 数字
	isNotZero := false
	// 是否第一个非空字符是符号
	isFirstSymbol := false
	// 非空格字符之前的字符个数
	nonSpaceCount := 0
	numStr := ""
	for i, val := range s {
		// 如果是空格则忽略掉
		if val == 32 && !isNonSpace {
			if isZero {
				return 0
			}
			nonSpaceCount++
			continue
		}
		// 如果是0并且没有出现过非空格字符，则忽略掉，或者首个非空格字符是符号也忽略掉
		if val == 48 {
			if isFirstSymbol && !isNotZero {
				continue
			}
			if !isNonSpace {
				isZero = true
				continue
			}
		}
		// 如果是-号并且没有出现过符号，则设置为负数
		if val == 45 {
			// 如果第一个非空格字符是符号
			if i == nonSpaceCount {
				isFirstSymbol = true
			}
			if isNonSpace {
				break
			}
			if isSymbol == 0 && !isZero {
				isSymbol = -1
				isNonSpace = true
				continue
			}
			// 如果第二次出现则返回0
			return 0
		}
		// 如果是+号并且没有出现过符号，则设置为正数
		if val == 43 {
			// 如果第一个非空格字符是符号
			if i == nonSpaceCount {
				isFirstSymbol = true
			}
			if isNonSpace {
				break
			}
			if isSymbol == 0 && !isZero {
				isSymbol = 1
				isNonSpace = true
				continue
			}
			// 如果第二次出现则返回0
			return 0
		}
		isNonSpace = true
		// 如果是数字，则拼接
		if val >= 48 && val <= 57 {
			isNotZero = true
			isNonSpace = true
			numStr += string(val)
			continue
		}
		// 如果是其他字符，则退出循环
		break
	}
	// 如果符号是默认值0，则赋值为1表示﹢
	if isSymbol == 0 {
		isSymbol = 1
	}
	// 转换数字
	resNum, _ := strconv.Atoi(numStr)
	// 增加符号
	resNum *= isSymbol
	// 对结果范围进行处理
	scope := 1 << 31
	if resNum < -1*scope {
		return -1 * scope
	} else if resNum > scope-1 {
		return scope - 1
	}
	return resNum
}
