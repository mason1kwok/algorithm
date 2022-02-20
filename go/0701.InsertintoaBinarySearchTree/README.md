## 递归

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    // 找到新节点的位置，插入
    if root == nil {
        return &TreeNode{Val: val}
    }

    if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    } else if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
```

## 复杂度
- 时间复杂度O(logN), N为二叉搜索树的节点数
- 空间复杂度O(1)