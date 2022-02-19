## 递归实现

```go
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    flatten(root.Left)
    flatten(root.Right)

    // 将左子树连接到右子树
    left := root.Left
    right := root.Right

    root.Left = nil
    root.Right = left

    p := root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = right
    
    
}
```

## 复杂度
- 时间复杂度O(2N)， N为二叉树节点数。
- 时间复杂度O(1)。
