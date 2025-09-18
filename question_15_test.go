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
 * @link: https://leetcode.cn/problems/3sum/description/
 */
func TestThreeSum(t *testing.T) {

	testCases := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := threeSum(testCases[i].input)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func threeSum(nums []int) (ans [][]int) {
	ans = make([][]int, 0) // 初始化，防止：[][]int(nil)
	sort.Ints(nums)
	for firstIndex := 0; firstIndex < len(nums); firstIndex++ {
		if firstIndex > 0 && nums[firstIndex] == nums[firstIndex-1] { // 去重
			continue
		}
		// 此时已经确定了一个值，可以将此题降维成两数之和，
		target := -nums[firstIndex] //，因为已经确定了target的值，所以需要后续的两个数字之和是 0-target，只需要将target * -1 即可，如果target是整数，*-1之后表示 后面的两个数之和需要为负数，如果target 是负数，*-1之后将target转成正数

		// 双指针，左指针为secondIndex
		right := len(nums) - 1 // 右指针

		for secondIndex := firstIndex + 1; secondIndex < len(nums); secondIndex++ {
			// 内层也需要去重
			if secondIndex > firstIndex+1 && nums[secondIndex] == nums[secondIndex-1] {
				continue
			}

			for right > secondIndex && nums[secondIndex]+nums[right] > target {
				right--
			}

			// 两个指针碰到说明都后续不需要进行计算
			if secondIndex == right {
				break
			}
			// 如果有 nums[secondIndex]+nums[right] = target ,则表示nums[secondIndex]+nums[right] + nums[firstIndex] = 0
			if nums[secondIndex]+nums[right] == target {
				ans = append(ans, []int{nums[firstIndex], nums[secondIndex], nums[right]})
			}
		}
	}

	return ans
}
