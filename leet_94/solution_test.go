package leet_94

import "testing"

func init() {

}

func TestSolution(t *testing.T) {
	res := inorderTraversal(&TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	})
	t.Log(res)
}
