package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-09-18 16:20
 * @link: https://leetcode.cn/problems/3sum-closest/description/
 */
func TestThreeSumClosest(t *testing.T) {

	testCases := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2},
		{[]int{2, 3, 8, 9, 10}, 16, 15},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, 1, 60},
		{[]int{-100, -98, -2, -1}, -101, -101},
		{[]int{-4, 2, 2, 3, 3, 3}, 0, 0},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := threeSumClosest(testCases[i].input, testCases[i].target)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

// 参考了题解的，个人的思路和题解差不多，但是有些细节想不到，总是不能过所有的用例。比如每次都保证right 或者 left指针至少移动一次。
func threeSumClosest(nums []int, target int) (ans int) {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)

	ans = 1<<31 - 1 // 因为要求一个最接近的值 所以初始化一个 最大值，
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp == target {
				return tmp // 相等直接返回
			}
			// 如果 ans 距离 target 比 tmp 更远，更新ans
			if abs(ans-target) > abs(tmp-target) {
				ans = tmp
			}
			// 判断移动指针，如果tmp大于 target，说明当前的right值过大，则吧right移动到一个和当前指针的值不同的位置
			if tmp > target {
				rt := right - 1
				for left < rt && nums[right] == nums[rt] {
					rt--
				}
				right = rt // 保证right至少向左移动一次
			} else {
				// 否则 说明left值小了，则吧left移动到一个和当前指针的值不同的位置
				lt := left + 1
				for right > lt && nums[left] == nums[lt] {
					lt++
				}
				left = lt // 同理，保证left至少向右移动一次
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}
