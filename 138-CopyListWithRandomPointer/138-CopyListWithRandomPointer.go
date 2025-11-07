// Last updated: 11/7/2025, 2:48:48 PM
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    nodeMap := make(map[*Node]*Node)
    curr := head

    for ; curr != nil; curr = curr.Next {
        nodeMap[curr] = &Node{Val: curr.Val}
    }

    curr = head
    for ;curr != nil; curr = curr.Next {
        nodeMap[curr].Next = nodeMap[curr.Next]
        nodeMap[curr].Random = nodeMap[curr.Random]
    }
    newDeepNode := nodeMap[head]
    return newDeepNode
}