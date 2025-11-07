// Last updated: 11/7/2025, 2:48:45 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    firstHalf := head
    secondHalf := slow.Next
    slow.Next = nil

    return merge(sortList(firstHalf), sortList(secondHalf))
}

func merge(l1, l2 *ListNode) (*ListNode) {
    dummy := &ListNode{}
    tail := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }

    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }


    for tail.Next != nil {
        tail = tail.Next
    }

    return dummy.Next
}