package q002

import (
	"github.com/magiconair/properties/assert"
	"math"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a1 := &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}
	a2 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	r := addTwoNumbers(a1, a2)
	result, pos := 0, 0
	for r != nil {
		result += int(math.Pow(10, float64(pos))) * r.Val
		r = r.Next
		pos++
	}
	assert.Equal(t, result, 1404)
}
