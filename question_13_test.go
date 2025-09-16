package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-16 10:50
 * @link: https://leetcode.cn/problems/roman-to-integer/
 */
func TestRomanToInt(t *testing.T) {

	testCases := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"MMMDCCXLIX", 3749},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := romanToInt(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

// 硬编码
var hash = map[int32]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (ans int) {

	for i, char := range s {
		cur := hash[char]
		if i < len(s)-1 && cur < hash[int32(s[i+1])] { // 正常罗马字符排序应该是前面的比后面的大，如果发现后面的比前面的大，说明时4 或者 9开头 ，按照罗马转阿拉伯数字的规则，需要做减法
			ans -= cur
		} else {
			ans += cur
		}
	}
	return ans
}
