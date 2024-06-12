package main

import "fmt"

func (g *Graph) BFS(start int) {
	queue := make([]int, 0)         // 采用队列来实现广度优先搜索
	g.visited = make([]bool, g.ver) // 确保visited数组与图的顶点数相匹配
	for i := range g.visited {
		g.visited[i] = false // 初始化visited数组
	}
	g.visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:] // 出队
		fmt.Println(v)

		for i := 0; i < g.ver; i++ {
			if g.adjMatrix[v][i] == 1 && !g.visited[i] {
				g.visited[i] = true
				queue = append(queue, i) // 如果存在邻接节点且未被访问就将其入队
			}
		}
	}
}
