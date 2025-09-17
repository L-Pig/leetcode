package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-17 10:30
 * @link: https://leetcode.cn/problems/longest-common-prefix/description/
 */
func TestLongestCommonPrefix(t *testing.T) {

	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
		{[]string{"a"}, "a"},
		{[]string{"a", "a", "b"}, ""},
		{[]string{"aaa", "aa", "aaa"}, "aa"},
		{[]string{"cir", "car"}, "c"},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := longestCommonPrefix(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

// 一开始考虑先找一个最小的元素，然后从数组中循环所有元素，从尾部开始比较，依次减少尾部下表，知道找到一个匹配的元素，发现过不了用例3后，改成从第一个字符开始比较，一直到第一个不匹配的字符传。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix = strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = getMaxCommonPrefix(prefix, strs[i]) // 每次找到一个公共前缀都更新一开始的前缀字符串
		if prefix == "" {                            // 如果为空，说明没有公共前缀，则没有继续比对的必要
			return prefix
		}
	}

	return prefix
}

func getMaxCommonPrefix(str1, str2 string) string {
	index := 0
	l := min(len(str1), len(str2))

	for i := 0; i < l; i++ {
		if index < l && str1[index] == str2[index] {
			index++
		}
	}

	return str1[:index]
}
