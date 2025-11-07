// Last updated: 11/7/2025, 2:47:50 PM
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    n1 := length(l1)
    n2 := length(l2)

    if n1 < n2 {
        l1 = padList(l1, n2 - n1)
    } else if n2 < n1 {
        l2 = padList(l2, n1 - n2)
    }

    sumHead, carry := addList(l1, l2) 
    if carry > 0 {
        node := &ListNode{Val: carry, Next: sumHead}
        return node
    }
    return sumHead
}

func length(node *ListNode) int {
    n := node
    len := 0
    for n != nil {
        len++
        n = n.Next
    }
    return len
}
func padList(node *ListNode, count int) *ListNode {
    for range count {
        node = &ListNode{Val: 0, Next: node}
    }
    return node
}

func addList(l1 *ListNode, l2 *ListNode) (*ListNode, int) {
    if l1 == nil && l2 == nil {
        return nil, 0
    }
    nextNode, carry := addList(l1.Next, l2.Next)
    sum := l1.Val + l2.Val + carry
    node := &ListNode{Val: sum % 10, Next: nextNode}
    return node, sum / 10
}