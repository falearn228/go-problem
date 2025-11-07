// Last updated: 11/7/2025, 2:48:20 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var prev *ListNode
    curr := head
    for curr != nil {
        nextPtr := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextPtr
    }
    return prev
}