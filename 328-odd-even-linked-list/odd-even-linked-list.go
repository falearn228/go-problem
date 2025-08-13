func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
	startOdd, startEven := head, head.Next
	endOdd, endEven := head, head.Next
	for endOdd != nil && endOdd.Next != nil && endEven != nil && endEven.Next != nil {
		endOdd.Next = endEven.Next
        endOdd = endOdd.Next
        endEven.Next = endOdd.Next
        endEven = endEven.Next
	}
	endOdd.Next = startEven
	return startOdd
}