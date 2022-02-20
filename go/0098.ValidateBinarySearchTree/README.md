## 递归

```go
func isValidBST(root *TreeNode) bool {
    return helper(root, nil, nil)
}

func helper(root *TreeNode, min *TreeNode, max *TreeNode) bool {
    if root == nil {
        return true
    }
    // 如果当前节点的值小于最小值，不是一个有效的二叉搜索树
    if min != nil && root.Val <= min.Val {
        return false
    }
    // 如果当前节点的值大于最大值，不是一个有效的二叉搜索树 
    if max != nil && root.Val >= max.Val {
        return false
    }

    // 返回当前节点的左右子树验证结果
    return helper(root.Left, min, root) && helper(root.Right, root, max)
}
```

## 复杂度
- 时间复杂度O(N), N为二叉树的节点
- 空间复杂度O(1)## 递归

```go
func isValidBST(root *TreeNode) bool {
    return helper(root, nil, nil)
}

func helper(root *TreeNode, min *TreeNode, max *TreeNode) bool {
    if root == nil {
        return true
    }
    // 如果当前节点的值小于最小值，不是一个有效的二叉搜索树
    if min != nil && root.Val <= min.Val {
        return false
    }
    // 如果当前节点的值大于最大值，不是一个有效的二叉搜索树 
    if max != nil && root.Val >= max.Val {
        return false
    }

    // 返回当前节点的左右子树验证结果
    return helper(root.Left, min, root) && helper(root.Right, root, max)
}
```

## 复杂度
- 时间复杂度O(N), N为二叉树的节点
- 空间复杂度O(1)