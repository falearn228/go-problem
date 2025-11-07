// Last updated: 11/7/2025, 2:47:52 PM
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
    freq := make(map[int]int)
    freq[0] = 1
    prefSum := make([]int, 1)
    var dfs func(node *TreeNode, prefSum []int)
    dfs = func(node *TreeNode, prefSum []int) {
        if node == nil {
            return
        }

        prefSum = append(prefSum, prefSum[len(prefSum)-1] + node.Val)
        prefA := prefSum[len(prefSum)-1] - targetSum
        prefB := prefSum[len(prefSum)-1]
        if freq[prefA] > 0 {
            count += freq[prefA]
        } 
        freq[prefB]++
        dfs(node.Left, prefSum)
        dfs(node.Right, prefSum)
        freq[prefB]--
    }
    dfs(root, prefSum)
    return count
}