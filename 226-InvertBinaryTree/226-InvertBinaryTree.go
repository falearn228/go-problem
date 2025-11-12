// Last updated: 11/12/2025, 5:35:24 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }

        dfs(r.Left)
        dfs(r.Right)
        if r.Left != nil || r.Right != nil {
            r.Left, r.Right = r.Right ,r.Left 
        }
    }

    dfs(root)
    return root
}