// Last updated: 11/7/2025, 2:48:47 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if fast == slow {
            return true
        }
    }

    return false
}