func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0
    tmp := head
    for tmp != nil {
        tmp = tmp.Next
        l++
    }

    tmp = head
    var prev *ListNode
    for i := 0; i < l - n; i++ {
        prev = tmp
        tmp = tmp.Next
    }

    if prev == nil {
        return tmp.Next
    }

    prev.Next = tmp.Next
    return head
}
