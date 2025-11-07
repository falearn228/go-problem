// Last updated: 11/7/2025, 2:48:58 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
     if head == nil || left == right {
        return head
    }

    dummy := &ListNode{Next: head}
    pre := dummy

    for range left-1 {
        pre = pre.Next
    }
    
    curr := pre.Next
    for range right - left {
        nxt := curr.Next
        curr.Next = nxt.Next
        nxt.Next = pre.Next
        pre.Next = nxt
    }

    return dummy.Next
}