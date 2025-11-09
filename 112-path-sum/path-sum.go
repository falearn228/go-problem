/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    answer := false 
    var dfs func(r *TreeNode, targetSum int)
    dfs = func(r *TreeNode, trgSum int) {
        if r == nil || answer {
            return
        }
        
        trgSum -= r.Val
        if r.Left == nil && r.Right == nil && trgSum == 0 {
            answer = true
        }
        dfs(r.Left, trgSum)
        dfs(r.Right, trgSum)

    }
    dfs(root, targetSum)
    return answer
}