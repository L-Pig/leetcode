package leetcode

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-08-25 18:15
 * @link: https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 * @description: 此题一开始用暴力解法写了一个解，后来实在没想到好办法了就看了题解，然后一步步调试理解题解
 */
func TestFindMedianSortedArrays(t *testing.T) {

	testCases := []struct {
		input    [2][]int
		expected float64
	}{
		{
			input:    [2][]int{{1, 3}, {2}},
			expected: 2.0000,
		},
		{
			input:    [2][]int{{1, 2}, {3, 4}},
			expected: 2.5000,
		},
		{
			input:    [2][]int{{1, 2, 3, 4}, {1, 2, 3, 4, 5, 6, 7, 8, 9}},
			expected: 4.0000,
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := findMedianSortedArrays(testCases[i].input[0], testCases[i].input[1])
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(getKElement(nums1, nums2, (len(nums1)+len(nums2))/2)) + float64(getKElement(nums1, nums2, (len(nums1)+len(nums2))/2-1))/2.0000
	} else {
		return float64(getKElement(nums1, nums2, (len(nums1)+len(nums2))/2+1))
	}
}

func getKElement(num1, num2 []int, k int) int {
	index1, index2 := 0, 0

	for {

		if index1 == len(num1) {
			return num2[index2+k-1]
		}

		if index2 == len(num2) {
			return num1[index1+k-1]
		}

		if k == 1 {
			return min(num1[index1], num2[index2]) // 因为一直都是通过index1,index2来伪造一种数组元素被抛弃的假象，实际上数组中的元素还是存在，所以用index1,index2来获取元素而不是用k-1来获取元素
		}

		half := k / 2 // /2 是因为k/2个不可能是第k小的元素，此处可以直接舍弃掉一半的元素。

		// 先比较两个数组中第k/2个元素的大小，小的一部分舍弃掉，同时更新index的值和k的值
		newIndex1 := min(index1+half, len(num1)) - 1 // -1 是因为数组下标从0开始，而数组长度从1开始
		newIndex2 := min(index2+half, len(num2)) - 1

		if num1[newIndex1] <= num2[newIndex2] {
			k -= newIndex1 - index1 + 1 // 第一轮因为前面3个元素已经被丢弃掉了，所以k重新计算，减去前面舍弃掉的元素个数
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

// 暴力解法
func f1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)

	start, end := 0, len(nums1)-1

	if len(nums1)%2 == 1 {
		return float64(nums1[start+end/2])
	} else {
		return (float64(nums1[start+end/2]) + float64(nums1[start+end/2+1])) / 2.0000
	}
}
