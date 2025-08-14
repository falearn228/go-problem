/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    depth := dfs(root, 0)
    return depth
}

func dfs(root *TreeNode, depth int) int {
    if root == nil {
        return depth
    }
    depth += 1
    depL := dfs(root.Left, depth)
    depR := dfs(root.Right, depth)
    if depL > depR {
        return depL
    }
    return depR
}