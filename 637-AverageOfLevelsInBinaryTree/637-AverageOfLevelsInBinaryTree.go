// Last updated: 11/12/2025, 5:34:51 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    res := []float64{}
    if root == nil {
        return res
    }

    queue := []*TreeNode{root} 
    start := 0
    for len(queue) - start > 0 {
        levelSize := len(queue) - start
        var levelSum float64

        for i := 0; i < levelSize; i++ {
            curNode := queue[start]
            start++

            levelSum += float64(curNode.Val)

            if curNode.Left != nil {
                queue = append(queue, curNode.Left)
            }
            if curNode.Right != nil {
                queue = append(queue, curNode.Right)
            }
        }
        
        // Считаем среднее и добавляем в результат
        res = append(res, levelSum/float64(levelSize))
    }

    return res
}
