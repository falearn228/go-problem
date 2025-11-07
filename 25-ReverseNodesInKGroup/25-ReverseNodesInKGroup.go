// Last updated: 11/7/2025, 2:49:25 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 1. вычислить длину до которой нужно заполнять
 // 2. рекурсивно менять местами элементы вплоть до k, сбрасывать счетчик
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k == 1 {
        return head
    }

    dummy := &ListNode{Next: head}
    prev, curr, next := dummy, head, head

    length := 0
    for curr != nil {
        length++
        curr = curr.Next
    }

    for i := 0; i+k <= length; i += k {
        curr = prev.Next
        next = curr.Next
        for j := 1; j < k; j++ {
            curr.Next = next.Next
            next.Next = prev.Next
            prev.Next = next
            next = curr.Next
        }
        prev = curr
    }

    return dummy.Next
}