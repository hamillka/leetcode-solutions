func pushFront(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{Val: val, Next: nil}
	}

	newHead := &ListNode{Val: val, Next: head}

	return newHead
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode = nil
	for head != nil {
		res = pushFront(res, head.Val)
		head = head.Next
	}

	return res
}