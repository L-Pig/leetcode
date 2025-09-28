package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-28 17:30
 * @link: https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
 */
func TestLetterCombinations(t *testing.T) {

	testCases := []struct {
		input    string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := letterCombinations(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

// 思路跟题解一样，就直接抄了题解的答案了
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i])) // index +1 每次暂时舍弃最前面的字符
		}
	}
}
