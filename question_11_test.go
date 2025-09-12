package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-11 10:30
 * @link: https://leetcode.cn/problems/container-with-most-water/
 */
func TestMaxArea(t *testing.T) {

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := maxArea(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func maxArea(height []int) (ans int) {

	left, right := 0, len(height)-1

	// 当left指针遇到了right指针，则表示都计算过
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)

		// 因为每次滑动指针取的是left和right中较小的那个值的位置进行滑动，所以不会漏掉对于左边位置不变，而右边出现一个比上一个右指针更大的值
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
