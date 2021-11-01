package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = helper(nums, l, mid-1)
	node.Right = helper(nums, mid+1, r)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return helper(nums, 0, len(nums)-1)
}
