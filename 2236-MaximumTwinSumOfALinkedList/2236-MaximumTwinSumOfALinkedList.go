// Last updated: 11/7/2025, 2:46:12 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    if head.Next.Next == nil {
        return head.Val + head.Next.Val
    }
    maxSum := 0

    slow, fast := head, head.Next.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // reverse 2'nd half
    var prev *ListNode
    curr := slow
    for curr != nil {
        nextPtr := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextPtr 
    }

    
    revNode, normNode := prev, head
    for revNode != normNode && revNode != nil && normNode != nil {
        nodeSum := revNode.Val + normNode.Val
        if maxSum < nodeSum {
            maxSum = nodeSum
        }
        revNode = revNode.Next
        normNode = normNode.Next
    }

    return maxSum
}