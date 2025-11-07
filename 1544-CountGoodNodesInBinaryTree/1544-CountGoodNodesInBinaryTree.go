// Last updated: 11/7/2025, 2:46:38 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    count := 0

    var dfs func(maxNode int, node *TreeNode)
    dfs = func(maxNode int, node *TreeNode) {
        if node == nil {
            return 
        }
        if node.Val >= maxNode {
            count++
            maxNode = node.Val
        }
        dfs(maxNode, node.Left)
        dfs(maxNode, node.Right)

    }
    dfs(root.Val, root) 

    return count
}

