package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MaxInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	if node.Val <= lower || node.Val >= upper {
		return false
	}

	if !helper(node.Left, lower, node.Val) {
		return false
	}

	if !helper(node.Right, node.Val, upper) {
		return false
	}

	return true
}
