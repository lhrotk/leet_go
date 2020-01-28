package q001

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	assert.Equal(t, twoSum(nums, 9), []int{1, 0})
}
