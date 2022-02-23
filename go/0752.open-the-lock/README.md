## BFS

```go
func openLock(deadends []string, target string) int {
    deads := map[string]bool{}
    for _, v := range deadends {
        deads[v] = true
    }
    visited := map[string]bool{}
    // 初始点入列
    queue := []string{"0000"}
    step := 0
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            // 出列
            node := queue[0]
            queue = queue[1:]

            if node == target {
                return step
            }
            // 跳过走过的点
            if _, ok := visited[node]; ok {
                continue
            }
            // 跳过死亡点
            if _, ok := deads[node]; ok {
                continue
            }
            // 记录走过的点
            visited[node] = true
            // 密码锁上下两个数字入列
            for j := 0; j < len(node); j++ {
                num := int(node[j] - '0')
                up := (num + 1) % 10
                down := (num + 9) % 10
                queue = append(queue, node[:j] + strconv.Itoa(up) + node[j+1:])
                queue = append(queue, node[:j] + strconv.Itoa(down) + node[j+1:])
            }
        }
        step++
    }

    return -1
}
```