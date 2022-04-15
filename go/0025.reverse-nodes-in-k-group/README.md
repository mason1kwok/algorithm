## 模仿
```go
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 增加dummy节点
    dummy := &ListNode{ Next: head}
    pre := dummy
    
    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if tail == nil { return dummy.Next }
        }

        next := tail.Next
        head, tail = reverse(head, tail)
        pre.Next = head
        tail.Next = next
        pre = tail
        head = tail.Next
    }

    return dummy.Next
}

// 反转链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode){
    prev := tail.Next
    p := head
    for prev != tail {
        next := p.Next
        p.Next = prev
        prev = p
        p = next
    }
    return tail, head
}
```

## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(1)## 模仿
```go
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 增加dummy节点
    dummy := &ListNode{ Next: head}
    pre := dummy
    
    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if tail == nil { return dummy.Next }
        }

        next := tail.Next
        head, tail = reverse(head, tail)
        pre.Next = head
        tail.Next = next
        pre = tail
        head = tail.Next
    }

    return dummy.Next
}

// 反转链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode){
    prev := tail.Next
    p := head
    for prev != tail {
        next := p.Next
        p.Next = prev
        prev = p
        p = next
    }
    return tail, head
}
```

## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(1)
