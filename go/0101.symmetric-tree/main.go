package main

func isSymmetric(root *TreeNode) bool {
    // 基本条件
    if root == nil {
        return true
    }

    var check func(left, right *TreeNode) bool
    check = func(left, right *TreeNode) bool {
        // 直接结束
        if left == nil || right == nil {
            return left == right
        }
        // 判断值
        if left.Val != right.Val {
            return false
        }

        // 继续向下
        return check(left.Left, right.Right) && check(left.Right, right.Left)
    }

    return check(root.Left, root.Right)
}
