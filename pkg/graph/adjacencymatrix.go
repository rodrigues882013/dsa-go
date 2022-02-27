package graph

import (
	"data-structures/pkg/utils"
)

// MGraph is a struct to represent a graph using an adjacency matrix
type MGraph struct {
	Storage [][]int
}

func (M *MGraph) BFS(start int) []int {
	panic("implement me")
}

// DFS return the path from a start point to all reachable vertices from the graph using depth-first approach
func (M *MGraph) DFS(start int) []int {
	visited := make([]bool, len(M.Storage))
	var helper func(first int)
	var ans []int

	helper = func(_start int) {
		visited[_start] = true
		ans = append(ans, _start)

		for idx := 0; idx < len(M.Storage[_start]); idx++ {
			edge := M.Storage[_start][idx]

			if !visited[idx] && edge != 0 {
				helper(idx)
			}
		}
	}

	helper(start)
	return ans
}

func (M *MGraph) Dijkstra(start int) []int {
	panic("implement me")
}

func (M *MGraph) TopSort() []int {
	panic("implement me")
}

// AddVertices adding a set o vertices in a graph
func (M *MGraph) AddVertices(vertices []int) {
	max := utils.MaxFromIterable(vertices...)

	if M.Storage == nil {
		M.Storage = make([][]int, max)

		for i := range M.Storage {
			M.Storage[i] = make([]int, max)
		}
	}
}

// AddVertex adding a vertex in a new or already created graph
func (M *MGraph) AddVertex(vertex int) {
	size := vertex

	if M.Storage == nil || len(M.Storage) < size {
		M.Storage = make([][]int, size)

		for i := range M.Storage {
			M.Storage[i] = make([]int, size)
		}
	}
}

// AddEdges connect to vertices in an already created graph
func (M *MGraph) AddEdges(edges [][]int) {
	if M.Storage == nil {
		panic("There is no vertices")
	}

	for _, edge := range edges {
		M.Storage[edge[0]][edge[1]] = 1
	}
}
