/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxSum := math.MinInt
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        if r == nil {
            return 0
        }
        leftGain := dfs(r.Left)
        leftGain = max(0, leftGain)

        rightGain := dfs(r.Right)
        rightGain = max(0, rightGain)

        curr := leftGain + r.Val + rightGain
        if curr > maxSum {
            maxSum = curr
        }
        
        if leftGain > rightGain {
            return r.Val + leftGain
        }
        return r.Val + rightGain
       
    }

    dfs(root)
    return maxSum
}