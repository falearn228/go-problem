// Last updated: 11/12/2025, 5:35:23 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
        return 0
    }

    ans := 1_000_000
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil || ans != 1_000_000 {
            return
        }

        dfs(r.Left)
        if ans != 1_000_000 {
            return
        }
        k--
        if k == 0 {
            ans = r.Val
        }
        dfs(r.Right)
    }

    dfs(root)
    return ans
}