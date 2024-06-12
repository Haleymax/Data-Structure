package main

import "fmt"

func (g *Graph) DFS(start int) {
	g.visited[start] = true
	fmt.Println(start)
	for i := 0; i < g.ver; i++ {
		if g.adjMatrix[start][i] == 1 && !g.visited[i] {
			g.DFS(i)
		}
	}
}
