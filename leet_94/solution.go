package leet_94

func init() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	} else {
		return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
	}
}

func TestInOrder(root *TreeNode) []int {
	return inorderTraversal(root)
}
