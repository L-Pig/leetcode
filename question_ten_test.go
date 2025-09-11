package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-10 14:30
 * @link: https://leetcode.cn/problems/regular-expression-matching/description/
 * @description: 抱歉，题解都没看懂🤣🤣🤣
 */
func TestIsMatch(t *testing.T) {

	testCases := []struct {
		input    string
		pattern  string
		expected bool
	}{
		{
			input:    "aa",
			pattern:  "a",
			expected: false,
		},
		{
			input:    "aa",
			pattern:  "a*",
			expected: true,
		},
		{
			input:    "aa",
			pattern:  ".*",
			expected: true,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := isMatch(testCases[i].input, testCases[i].pattern)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
