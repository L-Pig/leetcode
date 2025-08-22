package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2021-08-17 17:53
 * @link: https://leetcode.cn/problems/two-sum/description/
 */
func TestTwoSum(t *testing.T) {

	testCases := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 4}, 6, []int{}},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := twoSum(testCases[i].input, testCases[i].target)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}
