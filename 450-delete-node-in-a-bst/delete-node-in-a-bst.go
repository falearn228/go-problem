/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }

    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
        return root
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
        return root
    }

    // key == root.Val
    if root.Left == nil {
        return root.Right
    }
    if root.Right == nil {
        return root.Left
    }

    maxNode := root.Left
    for maxNode.Right != nil {
        maxNode = maxNode.Right
    }
    root.Val = maxNode.Val
    root.Left = deleteNode(root.Left, maxNode.Val)
    return root
}