## 快慢指针+DFS
```go
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
   return build(head, nil)
}

// 构建树
func build(begin, end *ListNode) *TreeNode {
    if begin == end {
        return nil
    }
    mid := getMid(begin, end)
    root := &TreeNode{Val: mid.Val}
    root.Left = build(begin, mid)
    root.Right = build(mid.Next, end)
    return root
}

// 快慢指针获取中间值
func getMid(begin, end *ListNode) *ListNode {
    fast, slow := begin, begin
    for fast != end && fast.Next != end {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
```

## 复杂度
- 时间复杂度O(nlogn)
- 空间复杂度O(1)