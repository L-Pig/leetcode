package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-08-25 18:15
 * @link: https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 * @description: 此题一开始用暴力解法写了一个解，后来实在没想到好办法了就看了题解，然后一步步调试理解题解
 */
func TestLongestPalindrome(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := longestPalindrome(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

// 从题解中拷贝的代码，我一开始想的思路也是从中间字符往外扩展，但是不知道如何实现，题解中这个解对我来说最容易理解，
func longestPalindrome(s string) string {

	return dp(s)
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1) // 可以解决 cbbd这种 无法用某一个中间字符向外扩展的情况，比如 如果以 c 为中间字符 则cb不符合回文串，如果以 第一个b为中间字符，则cbb不符合回文串，如果以第二个b为中间字符 则 bbd不符合回文串，至此，就可以预见会判断这个字符串不是回文串，实际上 bb 符合回文串的特性。
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func dp(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	ans := ""
	for l := 0; l < n; l++ { // l为本次循环遍历的子串长度
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
			}
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
