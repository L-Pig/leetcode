package leetcode

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-08 15:00
 * @link: https://leetcode.cn/problems/reverse-integer/?envType=problem-list-v2&envId=7raZsXAi
 */
func TestReverse(t *testing.T) {

	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    123,
			expected: 321,
		},
		{
			input:    -123,
			expected: -321,
		},
		{
			input:    120,
			expected: 21,
		},
		{
			input:    0,
			expected: 0,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := reverse(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func reverse(x int) (result int) {

	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 { // 判断溢出
			return 0
		}
		result = (x % 10) + result*10

		x = x / 10
	}
	return result
}
