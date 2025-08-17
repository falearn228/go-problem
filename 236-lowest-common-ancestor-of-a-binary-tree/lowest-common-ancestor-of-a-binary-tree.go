/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(node *TreeNode) *TreeNode
    dfs = func(node *TreeNode) *TreeNode {
        if node == nil {
            return nil
        }

        if node == p || node == q {
            return node
        }
        lNode := dfs(node.Left)
        rNode := dfs(node.Right)

        if lNode != nil && rNode != nil {
            return node
        }
        if rNode == nil  {
            return lNode
        }
        return rNode

    }
    answer := dfs(root)
    return answer
}