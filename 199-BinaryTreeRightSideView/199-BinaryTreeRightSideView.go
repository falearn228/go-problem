// Last updated: 11/7/2025, 2:48:23 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    var bfs func(n *TreeNode)
    bfs = func(n *TreeNode) {
        queue := make([]*TreeNode, 1)
        queue[0] = n

        for len(queue) > 0 {
            size := len(queue)
            var last *TreeNode

            for i := 0; i < size; i++ {
                node := queue[0]
                queue[0] = nil
                queue = queue[1:]

                if node == nil {
                    continue
                }

                last = node
                if node.Left != nil {
                    queue = append(queue, node.Left)
                }
                if node.Right != nil {
                    queue = append(queue, node.Right)
                }
            }

            if last != nil {
                res = append(res, last.Val)
            }
        }
    }
    bfs(root)
    return res
}