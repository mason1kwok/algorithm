## DFS
```go
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
    res = [][]int{}
    if root == nil {
        return res
    }
    dfs(root, targetSum, []int{})
    return res
}

func dfs(root *TreeNode, sum int, path []int) {
    if root == nil {
        return
    }
    remain := sum - root.Val
    if root.Left == nil && root.Right == nil {
        if remain == 0 {
            path = append(path, root.Val)
            res = append(res, append([]int{}, path...))
            path = path[:len(path)-1]
        }
        return
    }

    path = append(path, root.Val)
    dfs(root.Left, remain, path)
    path = path[:len(path)-1]

    path = append(path, root.Val)
    dfs(root.Right, remain, path)
    path = path[:len(path)-1]
}
```
## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(n)