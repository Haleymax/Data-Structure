package main

import "fmt"

// 采用邻接矩阵来实现图
type Graph struct {
	adjMatrix [][]int
	ver       int
	visited   []bool
}

// 初始化一个邻接矩阵
func (g *Graph) NewGraph(v int) {
	adj := make([][]int, v) //创建一个长为v的二维数组（长）
	for i := range adj {
		adj[i] = make([]int, v) //将每个切片的长度又设置为v（宽）
	}
	g.adjMatrix = adj
	g.ver = v
	g.visited = make([]bool, v)
}

// 添加边(无向图，from和to是代表在这两个节点之间添加边)
func (g *Graph) AddEdge(from, to int) {
	if from > g.ver || to > g.ver {
		panic("Surpass range of graph")
	}
	g.adjMatrix[from][to] = 1
	g.adjMatrix[to][from] = 1
}

// PrintMatrix 打印邻接矩阵
func (g *Graph) PrintMatrix() {
	for i := 0; i < g.ver; i++ {
		for j := 0; j < g.ver; j++ {
			fmt.Printf("%d ", g.adjMatrix[i][j])
		}
		fmt.Println()
	}
}

func main1() {
	// 创建一个具有5个顶点的图
	var graph Graph
	graph.NewGraph(5)

	// 添加一些边
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// 打印邻接矩阵
	fmt.Println("邻接矩阵:")
	graph.PrintMatrix()

	// 深度优先搜索从顶点0开始
	fmt.Println("\n深度优先搜索从顶点1开始:")
	graph.DFS(1)

	// 广度优先搜索从顶点0开始
	fmt.Println("\n广度优先搜索从顶点1开始:")
	graph.BFS(1)
}
