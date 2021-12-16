package leet2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func calRes(c int, head1, head2 *ListNode) (res, nc int) {
	val1, val2 := 0, 0
	if head1 != nil {
		val1 = head1.Val
	}
	if head2 != nil {
		val2 = head2.Val
	}
	res = c + val1 + val2
	nc = res / 10
	res = res % 10
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head1, head2 := l1, l2
	headRes := &ListNode{}
	p_next := headRes
	c := 0
	for head1 != nil || head2 != nil || c != 0 {
		posRes, nc := calRes(c, head1, head2)
		c = nc
		p_next.Next = &ListNode{Val: posRes}
		p_next = p_next.Next
		if head1 != nil {
			head1 = head1.Next
		}
		if head2 != nil {
			head2 = head2.Next
		}
	}
	return headRes.Next
}
