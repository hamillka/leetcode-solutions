func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
    	if fast.Next.Next == slow.Next {
    		return true
    	}
    	fast = fast.Next.Next
    	slow = slow.Next
    }

    return false
}