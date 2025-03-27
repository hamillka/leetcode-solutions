func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p != nil && q != nil {
        return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else {
        return p == q
    }
}
