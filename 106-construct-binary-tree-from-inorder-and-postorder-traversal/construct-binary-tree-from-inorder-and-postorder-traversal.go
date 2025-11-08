/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    idx := make(map[int]int, len(postorder))
    for i , v := range inorder {
        idx[v] = i
    }
    postIdx := len(postorder)-1
    var dfs func(inl, inr int) *TreeNode
    dfs = func(inl, inr int) *TreeNode {
        if inl > inr {
            return nil
        }
        rootVal := postorder[postIdx]
        root := &TreeNode{Val: rootVal}
        postIdx--
        mid := idx[rootVal]
        root.Right = dfs(mid+1, inr)
        root.Left = dfs(inl, mid-1)
        return root
    }

    return dfs(0, len(inorder)-1)
}