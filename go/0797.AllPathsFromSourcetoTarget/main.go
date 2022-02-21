package main

var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = [][]int{}
	path := []int{}
	traverse(graph, 0, path)
	return res
}

// 遍历图框架
func traverse(graph [][]int, s int, path []int) {
	// 做选择
	path = append(path, s)

	n := len(graph)
	// 到达终点
	if s == n-1 {
		res = append(res, append([]int(nil), path...))
		path = path[:len(path)-1]
		return
	}

	// 递归相邻节点
	for _, v := range graph[s] {
		traverse(graph, v, path)
	}
	// 撤销选择
	path = path[:len(path)-1]
}
