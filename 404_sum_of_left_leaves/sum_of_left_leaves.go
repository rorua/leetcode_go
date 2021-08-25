package sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	isLeaf := func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Left == nil && node.Right == nil {
			return true
		}
		return false
	}(root.Left)

	if root.Left != nil && isLeaf {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
