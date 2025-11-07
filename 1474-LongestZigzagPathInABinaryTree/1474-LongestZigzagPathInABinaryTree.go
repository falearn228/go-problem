// Last updated: 11/7/2025, 2:46:42 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    maxPath := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0
        }

        _, lRight := dfs(node.Left)
        rLeft, _ := dfs(node.Right)

        currLeft := 0
        if node.Left != nil {
            currLeft = 1 + lRight
        } 
        currRight := 0
        if node.Right != nil {
            currRight = 1 + rLeft 
        }

        maxPath = max(currRight, currLeft, maxPath)
        return currLeft, currRight
    }

    dfs(root)
    
    return maxPath
}

func max(a, b, c int) int {
    if a >= b && a >= c {
        return a
    }
    if b >= a && b >= c {
        return b
    }
    return c
}