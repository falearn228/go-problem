func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    maxSum := math.MinInt
    bestLevel := 1

    q := make([]*TreeNode, 0, 64)
    q = append(q, root)

    level := 0
    for head := 0; head < len(q); {
        level++
        // узлы текущего уровня в диапазоне [head, tail)
        tail := len(q)
        currSum := 0

        for i := head; i < tail; i++ {
            n := q[i]
            currSum += n.Val
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        head = tail

        if currSum > maxSum {
            maxSum = currSum
            bestLevel = level
        }
    }
    return bestLevel
}
