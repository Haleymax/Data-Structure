package main

import (
	"container/list"
	"fmt"
)

// 基于邻接表实现的图的广度优先搜索
func (g *LinkGraph) BFS(start int) {
	queue := list.New()
	queue.PushBack(start)
	g.read[start] = true
	fmt.Println(start)

	for queue.Len() > 0 {
		vertex := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for e := g.adjList[vertex].Front(); e != nil; e = e.Next() {
			neighbor := e.Value.(int)
			if !g.read[neighbor] {
				g.read[neighbor] = true
				fmt.Println(neighbor)
				queue.PushBack(neighbor)
			}
		}
	}
}
