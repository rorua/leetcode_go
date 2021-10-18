package max_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := maxDepth(root.Left)
	rmax := maxDepth(root.Right)
	if lmax > rmax {
		return 1 + lmax
	}
	return 1 + rmax
}
