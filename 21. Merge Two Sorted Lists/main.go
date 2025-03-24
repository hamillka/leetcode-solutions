func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			res.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			res.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		res = res.Next
	}

	if list1 != nil {
		res.Next = list1
	}
	if list2 != nil {
		res.Next = list2
	}

	return tmp.Next
}
