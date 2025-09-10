package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-08 14:40
 * @link: https://leetcode.cn/problems/palindrome-number/
 */
func TestIsPalindrome(t *testing.T) {

	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    121,
			expected: true,
		},
		{
			input:    -121,
			expected: false,
		},
		{
			input:    10,
			expected: false,
		},
		{
			input:    1234567899,
			expected: false,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := isPalindrome(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 { // 负数，末位为0 都不是回文数，因为整数不是0开头
		return false
	}
	rev := 0

	// 当rev < x   的时候，说明x已经计算到小于一半位数的位了， 此时就不需要再进行计算了，
	// 例如  121 x，和 rev的对应变化为
	//       x   ----   rev
	// 1.   121  ----    0
	// 2.    12  ----    1
	// 3.     1  ----    12
	// 计算到第三步的时候x已经小于rev了，此时值需要判断 x 和 rev 相不相等，或者原本的x为奇数位，则判断  rev/10 和当前的x是否相等，例子中 第三步  rev 最终结果为12 x 为1 ，则 rev/10 和 x 时相等的。

	// 同理，如果数字是 1221 ，则对应变化为
	//       x   ----   rev
	// 1.   1221  ----    0
	// 2.    122  ----    1
	// 3.     12  ----    12
	// 计算到第三步时 rev=x,所以跳出循环，此时 已经可以的到 x 是回文数
	for rev < x {

		rev = rev*10 + (x % 10)

		x /= 10
	}

	return rev == x || rev/10 == x // rev / 10 == x 用于处理奇数位的情况
}
