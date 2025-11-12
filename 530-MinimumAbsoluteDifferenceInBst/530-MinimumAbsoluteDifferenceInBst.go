// Last updated: 11/12/2025, 5:35:04 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Node struct {
    r *TreeNode 
    Val int
}

func getMinimumDifference(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var min = math.MaxInt
    var prev = -1
    var dfs func(root *TreeNode) 
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }
        
        dfs(r.Left)
        if prev != -1 {
            diff := r.Val - prev
            if min > diff {
                min = diff
            }
        }
        prev = r.Val
        dfs(r.Right)
    }

    dfs(root)
    return min
}

func abs(a, b int) int {
    if a < b {
        return b-a
    }
    return a-b
}