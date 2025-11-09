/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func exists(idx int, h int, node *TreeNode) bool {
    left, right := 0, (1<<h)-1
    for i := 0; i < h && node != nil; i++ {
        mid := left + (right-left)/2
        if idx <= mid {
            node = node.Left
            right = mid
        } else {
            node = node.Right
            left = mid + 1
        }
    }
    return node != nil
}

func depth(root *TreeNode) int {
    h := 0
    for root != nil && root.Left != nil {
        root = root.Left
        h++
    }
    return h
}

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    h := depth(root)
    if h == 0 {
        return 1
    }
    
    low, high := 0, (1<<h)-1
    for low <= high {
        mid := low + (high-low) / 2
        if exists(mid, h, root) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return (1<<h - 1) + low
}