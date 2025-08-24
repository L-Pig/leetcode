package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-08-24 11:06
 * @link: https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 */
func TestLengthOfLongestSubstring(t *testing.T) {

	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := lengthOfLongestSubstring(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	last := map[rune]int{} // 记录字符上一次出现的位置
	left := 0
	ans := 0
	for right, letter := range s {
		if index, ok := last[letter]; ok && index >= left { // 字符在上一次出现的位置之后出现，更新左指针，抛弃记录的上一次出现的位置
			left = index + 1
		}
		last[letter] = right
		ans = max(ans, right+1-left) // 计算最常字串的长度是否需要更新
	}

	return ans
}
