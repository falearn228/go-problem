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
    stack := []*TreeNode{node}
    var dfs func(n *TreeNode)
    dfs = func(n *TreeNode) {
        for len(stack) > 0 {
            n := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if n == nil {
                continue
            }
            if n.Left == nil && n.Right == nil {
                res = append(res, n.Val)
                continue
            }
            stack = append(stack, n.Left)
            stack = append(stack, n.Right)
        }
    }
    dfs(node)
    return res
} 
