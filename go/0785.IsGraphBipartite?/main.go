package main

var (
	ok      bool
	visited []bool
	color   []bool
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	ok = true
	visited = make([]bool, n)
	color = make([]bool, n)
	// 遍历每一个节点，验证所有子图都不是二分图
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(graph, i)
		}
	}

	return ok
}

func dfs(graph [][]int, v int) {
	// 如果不是二分图，就不再遍历了
	if !ok {
		return
	}
	visited[v] = true
	for _, w := range graph[v] {
		// 相邻节点w未被访问过，给w涂v不同的色
		// 相邻节点w被访问过，若w与v同色，则不是二分图
		if !visited[w] {
			color[w] = !color[v]
			// 继续遍历
			dfs(graph, w)
		} else if color[w] == color[v] {
			ok = false
		}
	}
}
