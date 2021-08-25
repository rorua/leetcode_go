package binary_tree_preorder_traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {

	var (
		res   = make([]int, 0)
		stack = make([]*TreeNode, 0)
	)

	if root == nil {
		return res
	}

	stack = append(stack, root)
	nodes := 1

	for nodes > 0 {
		node := stack[nodes-1]
		stack = stack[0 : nodes-1]
		nodes--
		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
			nodes++
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			nodes++
		}
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
