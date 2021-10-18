package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	helper(root, &d)
	return d
}

func helper(node *TreeNode, d *int) int {
	if node == nil {
		return -1
	}

	l := helper(node.Left, d)
	r := helper(node.Right, d)

	if *d < (l + r + 2) {
		*d = l + r + 2
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
