/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(root *TreeNode, root2 *TreeNode) bool
    dfs = func(root *TreeNode, root2 *TreeNode) bool {
        if root == nil && root2 == nil{
            return true
        }
        if root == nil || root2 == nil || root.Val != root2.Val {
            return false
        }
        dfs(root.Left, root2.Left)
        dfs(root.Right, root2.Right)

        return  dfs(root.Left, root2.Left) && dfs(root.Right, root2.Right)
    }

    answer := dfs(p, q)
    return answer
}