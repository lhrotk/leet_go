package leet2

import (
	"fmt"
	"testing"
)

func ListNodeFromSlice(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, num := range nums {
		p.Next = &ListNode{num, nil}
		p = p.Next
	}
	return head.Next
}

func TestAddTwoNumber(t *testing.T) {
	res := addTwoNumbers(ListNodeFromSlice([]int{1, 2, 3}), ListNodeFromSlice([]int{7, 8, 9}))
	nums := []int{}
	for res != nil {
		nums = append(nums, res.Val)
		res = res.Next
	}
	fmt.Println(nums)
}
