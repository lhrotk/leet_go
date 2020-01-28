package q002

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) getNext() *ListNode {
	if l == nil {
		return nil
	} else {
		return l.Next
	}
}

func (l *ListNode) getVal() int {
	if l == nil {
		return 0
	} else {
		return l.Val
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	c := 0
	head := &ListNode{}
	tail := head
	for !(c == 0 && l1 == nil && l2 == nil) {
		n := c + l1.getVal() + l2.getVal()
		l1 = l1.getNext()
		l2 = l2.getNext()
		c = n / 10
		tail.Next = &ListNode{Val: n % 10}
		tail = tail.Next
	}
	return head.Next
}
