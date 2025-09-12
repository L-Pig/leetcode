package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-08-25 18:15
 * @link: https://leetcode.cn/problems/zigzag-conversion/description/
 * @description：抄题解然后单步调试理解的，自己没有想明白
 */
func TestConvert(t *testing.T) {

	testCases := []struct {
		input    string
		input2   int
		expected string
	}{
		{
			input:    "PAYPALISHIRING",
			input2:   3,
			expected: "PAHNAPLSIIGYIR",
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := convert(testCases[i].input, testCases[i].input2)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}
