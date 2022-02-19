## 递归

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // 交换左右子树
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)

    return root
}
```

## 复杂度
- 时间复杂度O(N), N是二叉树的节点数
- 空间复杂度O(1)
