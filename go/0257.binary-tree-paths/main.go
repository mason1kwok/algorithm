pacakge main

var res []string

func binaryTreePaths(root *TreeNode) []string {
    res = []string{}
    dfs(root, "")
    return res
}

func dfs(root *TreeNode, path string) {
    if root == nil {
        return
    }
    if path == "" {
        path = fmt.Sprintf("%d", root.Val)
    } else {
        path = fmt.Sprintf("%s->%d", path, root.Val)
    }
    if root.Left == nil && root.Right == nil {
        res = append(res, path)
        return
    }
    dfs(root.Left, path)
    dfs(root.Right, path)
}