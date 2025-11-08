/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    idx := make(map[int]int, len(preorder))
    for i , v := range inorder {
        idx[v] = i
    }
    preIdx := 0
    var dfs func(inl, inr int) *TreeNode
    dfs = func(inl, inr int) *TreeNode {
        if inl > inr {
            return nil
        }
        rootVal := preorder[preIdx]
        root := &TreeNode{Val: rootVal}
        preIdx++
        mid := idx[rootVal]
        root.Left = dfs(inl, mid-1)
        root.Right = dfs(mid+1, inr)
        return root
    }

    return dfs(0, len(inorder)-1)
}