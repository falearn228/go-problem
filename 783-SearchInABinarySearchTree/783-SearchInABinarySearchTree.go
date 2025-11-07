// Last updated: 11/7/2025, 2:47:19 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func init() {
	debug.SetMemoryLimit(1)
}

func searchBST(root *TreeNode, val int) *TreeNode {
    var dfs func(*TreeNode) *TreeNode
    dfs = func(n *TreeNode) *TreeNode {
        if n == nil {
            return nil
        }

        if n.Val == val {
            return n
        }

        left := dfs(n.Left)
        right := dfs(n.Right)

        if left != nil {
            return left
        }
        return right
    }
    return dfs(root)
}