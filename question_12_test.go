package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-15 15:30
 * @link: https://leetcode.cn/problems/integer-to-roman/description/
 */
func TestIntToRoman(t *testing.T) {

	testCases := []struct {
		input    int
		expected string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := intToRoman(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

// 硬编码
var table = [4][10]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位
	{"", "M", "MM", "MMM"}, // 千位
}

func intToRoman(num int) string {

	return table[3][num/1000] + table[2][num/100%10] + table[1][num/10%10] + table[0][num%10]
}
