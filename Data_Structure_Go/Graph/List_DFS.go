package main

import "fmt"

// 基于邻接表的深度优先搜索
func (g *LinkGraph) DFS(start int) {
	g.read[start] = true
	fmt.Println(start)
	for i := g.adjList[start].Front(); i != nil; i = i.Next() {
		neigbor := i.Value.(int)
		if !g.read[neigbor] {
			g.DFS(neigbor)
		}
	}
}
