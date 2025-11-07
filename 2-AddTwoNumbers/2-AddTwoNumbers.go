// Last updated: 11/7/2025, 2:49:47 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 18 % 10 = 8
 // 18 // 10 = 1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    curr1, curr2 := l1, l2
    dummy := &ListNode{}
    result := dummy
    var carry int
    for curr1 != nil || curr2 != nil || carry != 0 {
        curr1Val, curr2Val := 0, 0
        if curr1 != nil {
            curr1Val = curr1.Val
            curr1 = curr1.Next
        }
        if curr2 != nil {
            curr2Val = curr2.Val
            curr2 = curr2.Next
        }
        
        sum := curr1Val + curr2Val + carry
        carry = sum / 10
        digit := sum % 10

        result.Next = &ListNode{Val: digit}

        result = result.Next
    }

    return dummy.Next
}