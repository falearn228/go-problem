// Last updated: 11/7/2025, 2:49:27 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return head
    }
    dummy := &ListNode{Next: head}
    slow, fast := dummy, dummy
    for range n+1 {
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next
    return dummy.Next
}