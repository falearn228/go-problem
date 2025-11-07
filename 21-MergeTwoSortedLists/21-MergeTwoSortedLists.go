// Last updated: 11/7/2025, 2:49:26 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    curr1, curr2 := list1, list2
    dummy := &ListNode{}
    result := dummy
    for curr1 != nil && curr2 != nil {
        if curr1 != nil && curr1.Val < curr2.Val {
            result.Next = &ListNode{Val: curr1.Val}
            curr1 = curr1.Next
        } else if curr2 != nil {
            result.Next = &ListNode{Val: curr2.Val}
            curr2 = curr2.Next
        }
        result = result.Next
    }

    for curr1 != nil {
        result.Next = &ListNode{Val: curr1.Val}
        result = result.Next
        curr1 = curr1.Next
    }
    for curr2 != nil {
        result.Next = &ListNode{Val: curr2.Val}
        result = result.Next
        curr2 = curr2.Next
    }

    return dummy.Next
}