package invert_binary_tree

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right, root.Left = root.Left, root.Right

	invertTree(root.Right)
	invertTree(root.Left)

	return root
}
