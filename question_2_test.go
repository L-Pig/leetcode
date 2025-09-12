package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @author: L-Pig
 * @date: 2025-08-23 11:00
 * @link: https://leetcode.cn/problems/add-two-numbers/description/
 */
func TestAddTwoNumbers(t *testing.T) {

	testCases := []struct {
		input struct {
			l1 *ListNode
			l2 *ListNode
		}
		expected *ListNode
	}{
		{
			input: struct {
				l1 *ListNode
				l2 *ListNode
			}{
				l1: newListNode(2, 4, 3),
				l2: newListNode(5, 6, 4),
			},
			expected: newListNode(7, 0, 8),
		},
		{
			input: struct {
				l1 *ListNode
				l2 *ListNode
			}{
				l1: newListNode(0),
				l2: newListNode(0),
			},
			expected: newListNode(0),
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			actual := addTwoNumbers(testCases[i].input.l1, testCases[i].input.l2)
			assert.Equal(t, testCases[i].expected, actual)
		})
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {

	var carry = 0
	var cur *ListNode

	for l1 != nil || l2 != nil { // 两个节点中有一个节点还有数据都需要继续计算
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum} // 此处 cur 不会为 nil
			cur = cur.Next
		}
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将 2 ，4，3 构建成单向链表
func newListNode(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	ret := &ListNode{}
	ret.Val = vals[0]
	ret.Next = newListNode(vals[1:]...)
	return ret
}
