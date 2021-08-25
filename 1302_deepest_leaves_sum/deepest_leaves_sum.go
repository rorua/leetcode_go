package deepest_leaves_sum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	left, right := 0, 0
	if root.Left != nil {
		left += deepestLeavesSum(root.Left)
	}
	if root.Right != nil {
		right += deepestLeavesSum(root.Right)
	}

	if left > right {
		return left + root.Val
	}
	return right + root.Val
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
