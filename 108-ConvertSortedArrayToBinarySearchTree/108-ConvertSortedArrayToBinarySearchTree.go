// Last updated: 11/7/2025, 2:48:56 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	l, r := 0, len(nums)-1

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := (l + r) / 2

		node := &TreeNode{}
		node.Val = nums[mid]
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)

		return node
	}

	node := build(l, r)
	return node
}

