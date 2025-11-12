// Last updated: 11/12/2025, 5:35:25 PM
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    height := calcLeftHeight(root)
    return helper(root, height)
}

func helper(r *TreeNode, h int) int {
    if r == nil {
        return 0
    }
    leftSubRightLeg := calcRightHeight(r.Left)
    leftSubLeftLeg := h-1

    leftSubSize, rightSubSize := 0, 0
    if leftSubRightLeg != leftSubLeftLeg {
        leftSubSize = helper(r.Left, h-1)
        rightSubSize = (1<<(h-2))-1
    } else {
        leftSubSize = (1<<(h-1))-1
        rightSubSize = helper(r.Right, h-1)
    }
    return 1 + leftSubSize + rightSubSize
}

func calcLeftHeight(root *TreeNode) int {
    leng := 0
    for root != nil {
        root = root.Left
        leng++
    }
    return leng
}

func calcRightHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return  1 + calcRightHeight(root.Right)
}


