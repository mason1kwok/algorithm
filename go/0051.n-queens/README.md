## 回溯算法

```go
var res [][]string

func solveNQueens(n int) [][]string {
    res = [][]string{}
    board := make([]string, n)
    // board = ["...."] * n
    for i := 0; i < n; i++ {
        board[i] = strings.Repeat(".", n)
    }
    backtrack(board, 0)
    return res
}

func backtrack(board []string, row int) {
    // 符合要求，加入结果
    if row == len(board) {
        res = append(res, append([]string{}, board...))
        return   
    }

    n := len(board[row])

    for col := 0; col < n; col++ {
        if (!vaild(board, row, col)) {
            continue
        }
        s := board[row]
        // 做选择
        board[row] = s[:col] + "Q" + s[col+1:]
        // 递归
        backtrack(board, row + 1)
        // 撤销选择
        board[row] = s[:col] + "." + s[col+1:]
    }
}

// 判断皇后是否冲突
func vaild(board []string, row, col int) bool {
    n := len(board)
    // 判断列是否有皇后冲突
    for i := 0; i < n; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    // 判断右上方是否有皇后冲突
    for i, j := row - 1, col + 1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    // 判断左上方是否有皇后冲突
    for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    return true
}
```

## 复杂度
- 时间复杂度O(n!)，n是皇后的数量。
- 空间复杂度O(n)