/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// if sumNodes > targetSum -> switch root (after left, right)
func pathSum(root *TreeNode, targetSum int) int {
    count := 0
    var dfs func(node *TreeNode, need int) 
    dfs = func(node *TreeNode, need int) {
        if node == nil {
            return
        }
        newNeed := need + node.Val
        if newNeed == targetSum {
            count++
        }
        
        dfs(node.Left, newNeed)
        dfs(node.Right, newNeed)
    }
   
    var walkStarts func(node *TreeNode)
    walkStarts = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node, 0)
        walkStarts(node.Left)
        walkStarts(node.Right)
        return
    }
    walkStarts(root)
    return count
}