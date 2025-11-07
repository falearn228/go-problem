// Last updated: 11/7/2025, 2:46:13 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }
    slow, fast := head, head.Next
    for fast.Next != nil && fast.Next.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }
    // Теперь slow указывает на середину списка
    // 1 1 -> 2 3 -> 3 ...
    slow.Next = slow.Next.Next
    return head
}