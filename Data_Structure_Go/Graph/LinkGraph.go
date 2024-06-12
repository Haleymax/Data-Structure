package main

import (
	"container/list"
	"fmt"
)

type LinkGraph struct {
	ver     int //顶点
	adjList map[int]*list.List
	read    map[int]bool
}

// 初始化图
func (g *LinkGraph) Init(v int) {
	g.ver = v
	g.adjList = make(map[int]*list.List)
	g.read = make(map[int]bool)
	for i := 0; i < v; i++ {
		g.adjList[i] = list.New() //为每个节点创建一个边列表
	}
}

// 添加边（无向图）
func (g *LinkGraph) AddEdge(from, to int) {
	if from >= g.ver || to >= g.ver {
		panic("Vertex out of range")
	}
	g.adjList[from].PushBack(to)
	g.adjList[to].PushBack(from)
}

func main2() {
	var g LinkGraph
	g.Init(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3) // 自环

	fmt.Println("DFS 从节点 2 开始:")
	g.DFS(2)

	for k := range g.read {
		g.read[k] = false
	}

	fmt.Println("BFS 从节点 2 开始:")
	g.BFS(2)
}
