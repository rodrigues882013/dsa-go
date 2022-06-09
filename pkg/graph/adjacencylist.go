package graph

import (
	"container/heap"
	"data-structures/pkg/linears"
	"data-structures/pkg/utils"
	"log"
	"math"
)

// Edge is a struct representing a edge with its respective incoming node and its weight
type Edge struct {
	Id int
	Weight int
}

// LGraph is a struct to represent a graph using an adjacency list, having a map from int to Edge struct
type LGraph struct {
	Storage map[int][]*Edge
}

// BFS return the path from a start point to all reachable vertices from the graph using breadth-first approach
func (g *LGraph) BFS(start int) []int {
	visited := make([]bool, len(g.Storage))
	var queue []int
	var ans []int

	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		polled := queue[0]
		queue = queue[1:]
		ans = append(ans, polled)

		for _, edge := range g.Storage[polled] {
			if !visited[edge.Id] {
				visited[edge.Id] = true
				queue = append(queue, edge.Id)
			}
		}
	}

	return ans
}

// DFS return the path from a start point to all reachable vertices from the graph using depth-first approach
func (g *LGraph) DFS(start int) []int {
	visited := make([]bool, len(g.Storage))
	var helper func(first int)
	var ans []int

	helper = func(_start int) {
		visited[_start] = true
		ans = append(ans, _start)

		for _, edge := range g.Storage[_start] {
			if !visited[edge.Id] {
				helper(edge.Id)
			}
		}
	}

	helper(start)
	return ans

}

func (g *LGraph) Dijkstra(start int) []int {
	numberOfVertices := len(g.Storage)
	visited := make([]bool, numberOfVertices)
	dist := make([]int, numberOfVertices)
	utils.Fill(dist, numberOfVertices, math.MaxInt)
	dist[0] = 0
	pq := make(linears.PriorityQueue, 1)

	pq.Push(&linears.Item{
		Value:    start,
		Priority: 0,
	})
	heap.Init(&pq)

	for pq.Len() > 0 {
		popped := pq.Pop().(*linears.Item)
		visited[popped.Value] = true

		for _, edge := range g.Storage[popped.Value] {

			if visited[edge.Id] {
				continue
			}

			newDistance := dist[popped.Value] + edge.Weight

			if newDistance < dist[edge.Id] {
				dist[edge.Id] = newDistance
				pq.Push(&linears.Item{
					Value: edge.Id,
					Priority: newDistance,
				})
			}
		}
	}

	return dist
}

func (g *LGraph) TopSort() []int {
	log.Fatal("implement me")
	return nil
}

// AddVertices adding a set o vertices in a graph
func (g *LGraph) AddVertices(vertices []int) {
	for _, vertex := range vertices {
		g.AddVertex(vertex)
	}
}

// AddVertex adding a vertex in a new or already created graph
func (g *LGraph) AddVertex(vertex int) {
	if g.Storage == nil {
		g.Storage = make(map[int][]*Edge)
	}
	g.Storage[vertex] = []*Edge{}
}

// AddEdges connect to vertices in an already created graph
func (g *LGraph) AddEdges(edges [][]int) {

	if g.Storage == nil {
		log.Fatal("implement me")
		return
	}

	for _, edge := range edges {
		g.Storage[edge[0]] = append(g.Storage[edge[0]], &Edge{edge[1], edge[2]})
	}
}
