package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-09 10:35
 * @link: https://leetcode.cn/problems/string-to-integer-atoi/?envType=problem-list-v2&envId=7raZsXAi
 */
func TestMyAtoi(t *testing.T) {

	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    "   -042",
			expected: -42,
		},
		{
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			input:    "0-1",
			expected: 0,
		},
		{
			input:    "4193 with words",
			expected: 4193,
		},
		{
			input:    "+-20",
			expected: 0,
		},
		{
			input:    "20+-",
			expected: 20,
		},
		{
			input:    "      ",
			expected: 0,
		},
		{
			input:    "+1",
			expected: 1,
		},
		{
			input:    "-2147483649",
			expected: -2147483648,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := myAtoi(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func myAtoi(s string) int {

	result, symbol := 0, 1

	idx := 0
	// 去空格
	for _, c := range s {
		if c == ' ' {
			idx++
		} else {
			break
		}
	}

	if idx >= len(s) {
		return 0
	}

	// 确定符号，没有符号则默认是正数
	if s[idx] == '-' {
		symbol = -1
		idx++
	} else if s[idx] == '+' {
		symbol = 1
		idx++
	}

	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		if symbol*(result*10+int(s[idx]-'0')) > (1<<31 - 1) { // c-'0' 吧 c 减去 0 的 ASCII码，将转换成0-9的数字，比如‘1’ 的ASCII码是49，减去0的ASCII码是48，结果为1, 先判断防止越界，
			return 1<<31 - 1
		}
		if symbol*(result*10+int(s[idx]-'0')) < (-1 << 31) {
			return -1 << 31
		}

		result = result*10 + int(s[idx]-'0')
		idx++
	}
	return result * symbol
}
