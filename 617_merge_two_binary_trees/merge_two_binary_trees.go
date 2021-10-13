package merge_two_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}

	var res = new(TreeNode)
	res.Val = root1.Val + root2.Val
	res.Left = mergeTrees(root1.Left, root2.Left)
	res.Right = mergeTrees(root1.Right, root2.Right)

	return res
}
