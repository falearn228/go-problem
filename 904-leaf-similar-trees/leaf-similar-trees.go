/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leafs1 := collectLeaves(root1)
    leafs2 := collectLeaves(root2)
    fmt.Println(leafs1, leafs2)

    return slices.Equal(leafs1, leafs2)
}

func collectLeaves(node *TreeNode) []int {
    res := []int{}
    var dfs func(n *TreeNode)
    dfs = func(n *TreeNode) {
        if n == nil {
            return
        }
        if n.Left == nil && n.Right == nil {
            res = append(res, n.Val)
            return
        }
        dfs(n.Left)
        dfs(n.Right)
    }
    dfs(node)
    return res
} 
